package geonetApi

import (
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

const TimeFormatLayout = "2006-01-02T15:04:05.000Z"

type QFeatureCollection struct {
	Type     string     `json:"type"`
	Features []QFeature `json:"features"`
}

type QFeature struct {
	Type       string      `json:"type"`
	Geometry   QGeometry   `json:"geometry"`
	Properties QProperties `json:"properties"`
}

type QGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type QProperties struct {
	Depth     float64 `json:"depth"`
	Locality  string  `json:"locality"`
	Magnitude float64 `json:"magnitude"`
	Mmi       int     `json:"mmi"`
	PublicId  string  `json:"publicID"`
	Quality   string  `json:"quality"`
	Time      string  `json:"time"`
}

var mmiScale = []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8}

var qualityCriteria = []string{"", "automatic", "best", "deleted", "preliminary"}

var intensityGroup = []string{"", "minor", "moderate", "severe", "extreme",
	"strong", "servere", "extreme"}

var intensityCAP = map[int]string{
	1:  "minor",
	2:  "minor",
	3:  "minor",
	4:  "minor",
	5:  "minor",
	6:  "moderate",
	7:  "severe",
	8:  "extreme",
	9:  "extreme",
	10: "extreme",
	11: "extreme",
	12: "extreme",
}

func MapToString(m map[int]string) string {
	var keys []int
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	var result string
	for _, key := range keys {
		value := m[key]
		result += fmt.Sprintf("%d: %s\n", key, value)
	}
	return result
}

func GetKeysForValue(m map[int]string, value string) []int {
	var keys []int
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
		}
	}
	sort.Ints(keys)

	return keys
}

func IsTimeStringValid(s string) bool {
	pattern := `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\.\d{3}Z$`
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		fmt.Println("Regex error:", err)
		return false
	}
	return match
}

func ParseTimeString(timeConvert string, timeReference string) time.Time {
	parsedTime, err := time.Parse(timeReference, timeConvert)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Time{}
	}
	return parsedTime
}

func ParseFloatString(numberString string) float64 {
	floatValue, err := strconv.ParseFloat(numberString, 64)

	if err != nil {
		fmt.Println("Error parsing float:", err)
		return 0.0
	}
	return floatValue
}

func GetGeonetQuakes(
	publicId string,
	timeAfter string,
	depthAbove float64,
	intensity string,
	magnitudeAbove float64,
	mmi int,
	localityIn string,
	quality string,
) []QFeature {

	var quakeData HTTPGetResult
	var featureCollection QFeatureCollection
	var quakeFeatures []QFeature

	if !slices.Contains(mmiScale, mmi) {
		fmt.Println("Invalid level:", mmi)
		fmt.Println("Allowed level:",
			strings.Trim(strings.Replace(
				fmt.Sprint(mmiScale), " ", ", ", -1), "[]"))
		return quakeFeatures
	}

	if !slices.Contains(qualityCriteria, quality) {
		fmt.Println("Invalid quality:", quality)
		fmt.Println("Allowed quality:", strings.Join(qualityCriteria, ", "))
		return quakeFeatures
	}

	if !slices.Contains(intensityGroup, intensity) {
		fmt.Println("Invalid intensity:", intensity)
		fmt.Println("Allowed intensity:", strings.Join(intensityGroup, ", "))
		return quakeFeatures
	}

	if intensity != "" && intensityCAP[mmi] != intensity {
		fmt.Println("Invalid mmi not maching intensity cap range:", intensity)
		fmt.Println("Intensity cap table:\n", MapToString(intensityCAP))
		return quakeFeatures
	}

	if timeAfter != "" && !IsTimeStringValid(timeAfter) {
		fmt.Println("Invalid time:", timeAfter)
		fmt.Printf("Allowed time format sample: '%s'", TimeFormatLayout)
		return quakeFeatures
	}

	var intensityValues = []int{mmi}

	if intensity != "" {
		intensityValues = GetKeysForValue(intensityCAP, intensity)
	}

	for _, item := range intensityValues {
		newGeonetQuakeApiUrl := GeonetQuakeApiUrl + "?MMI=" +
			fmt.Sprintf("%d", item)
		quakeData = HTTPGet(newGeonetQuakeApiUrl, GeonetQuakeAcceptHeader)
		err := json.Unmarshal([]byte(quakeData.Data), &featureCollection)

		if err != nil {
			fmt.Println("Error:", err)
			return quakeFeatures
		}

		for _, item := range featureCollection.Features {

			IsTimeInRange := true
			if timeAfter != "" {
				timeAfterParsed := ParseTimeString(timeAfter, TimeFormatLayout)
				timeToValidateParsed := ParseTimeString(item.Properties.Time, TimeFormatLayout)

				if !(timeToValidateParsed.After(timeAfterParsed) ||
					timeToValidateParsed.Equal(timeAfterParsed)) {

					IsTimeInRange = false
				}
			}

			IsDepthAbove := true
			if depthAbove != -1 &&
				item.Properties.Depth < depthAbove {

				IsDepthAbove = false
			}

			IsMagnitudeAbove := true
			if magnitudeAbove != -1 &&
				item.Properties.Magnitude < magnitudeAbove {

				IsMagnitudeAbove = false
			}

			IsLocalityIn := true
			if localityIn != "" &&
				!strings.Contains(strings.ToLower(item.Properties.Locality),
					strings.ToLower(localityIn)) {

				IsLocalityIn = false
			}

			if (publicId == "" || publicId == item.Properties.PublicId) &&
				(quality == "" || quality == item.Properties.Quality) &&
				IsDepthAbove &&
				IsLocalityIn &&
				IsMagnitudeAbove &&
				IsTimeInRange {

				quakeFeatures = append(quakeFeatures, item)
			}
		}

	}

	return quakeFeatures
}
