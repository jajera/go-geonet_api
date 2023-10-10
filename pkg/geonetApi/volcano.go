package geonetApi

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"
)

type VFeatureCollection struct {
	Type     string     `json:"type"`
	Features []VFeature `json:"features"`
}

type VFeature struct {
	Type       string      `json:"type"`
	Geometry   VGeometry   `json:"geometry"`
	Properties VProperties `json:"properties"`
}

type VGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type VProperties struct {
	Acc          string `json:"acc"`
	Activity     string `json:"activity"`
	Hazards      string `json:"hazards"`
	Level        int    `json:"level"`
	VolcanoID    string `json:"volcanoID"`
	VolcanoTitle string `json:"volcanoTitle"`
}

type VolcanoID string

var volcanoAlertLevel = []int{-1, 0, 1, 2, 3, 4, 5}

var volcanoAviationColourCodes = []string{"All", "Green", "Orange", "Red", "Yellow"}

func GetGeonetVolcanoAlertLevel(
	volcanoID string,
	volcanoTitle string,
	level int,
	acc string,
	activity string,
	hazards string,
) []VFeature {

	var valData HTTPGetResult
	var featureCollection VFeatureCollection
	var valFeatures []VFeature

	if !slices.Contains(volcanoAviationColourCodes, acc) {
		fmt.Println("Invalid acc:", acc)
		fmt.Println("Allowed acc:", strings.Join(volcanoAviationColourCodes, ", "))
		return valFeatures
	}

	if !slices.Contains(volcanoAlertLevel, level) {
		fmt.Println("Invalid level:", level)
		fmt.Println("Allowed level:",
			strings.Trim(strings.Replace(
				fmt.Sprint(volcanoAlertLevel), " ", ", ", -1), "[]"))
		return valFeatures
	}

	valData = HTTPGet(GeonetValApiUrl, GeonetValAcceptHeader)
	err := json.Unmarshal([]byte(valData.Data), &featureCollection)
	if err != nil {
		fmt.Println("Error:", err)
		return valFeatures
	}

	for _, item := range featureCollection.Features {
		if (volcanoID == "" || volcanoID == item.Properties.VolcanoID) &&
			(volcanoTitle == "" || volcanoTitle == item.Properties.VolcanoTitle) &&
			(level == -1 || level == item.Properties.Level) &&
			(acc == "All" || acc == item.Properties.Acc) &&
			(activity == "" || activity == item.Properties.Activity) &&
			(hazards == "" || hazards == item.Properties.Hazards) {

			valFeatures = append(valFeatures, item)
		}
	}

	return valFeatures
}

func GetGeonetVolcanoIds() []VolcanoID {
	var volcanoIDs []VolcanoID

	val := GetGeonetVolcanoAlertLevel(
		"",    // volcanoID
		"",    // volcanoTitle
		-1,    // int
		"All", // acc
		"",    // activity
		"",    // hazards
	)

	for _, item := range val {
		volcanoID := VolcanoID(item.Properties.VolcanoID)
		volcanoIDs = append(volcanoIDs, volcanoID)

		sort.Slice(volcanoIDs, func(i, j int) bool {
			return volcanoIDs[i] < volcanoIDs[j]
		})
	}

	return volcanoIDs
}
