package main

import (
	"encoding/json"
	"fmt"
	"geonetApp/pkg/geonetApi"
)

func main() {
	// Print11LatestNews()
	// PrintLatestGeohazardEventsNews()
	// PrintVolcanoIds()
	// PrintVolcanoAlertLevel()
	PrintQuakesFiltered()
}

func PrintQuakesFiltered() {
	quakesFiltered := geonetApi.GetGeonetQuakes(
		"",                         // publicId
		"2023-10-15T04:24:31.290Z", // timeAfter
		1,                          // depthAbove
		"",                         // intensity
		1,                          // magnitudeAbove
		3,                          // mmi
		"",                         // localityIn
		"",                         // quality
	)

	formattedJSON, err := json.MarshalIndent(quakesFiltered, "", "  ")

	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("Quakes filtered:", "\n", string(formattedJSON))
}

func PrintVolcanoAlertLevel() {
	volcanoAlertLevel := geonetApi.GetGeonetVolcanoAlertLevel(
		"",    // volcanoID
		"",    // volcanoTitle
		-1,    // int
		"All", // acc
		"",    // activity
		"",    // hazards
	)

	formattedJSON, err := json.MarshalIndent(volcanoAlertLevel, "", "  ")

	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("Volcano IDs:", "\n", string(formattedJSON))
}

func PrintVolcanoIds() {
	volcanoIds := geonetApi.GetGeonetVolcanoIds()

	formattedJSON, err := json.MarshalIndent(volcanoIds, "", "  ")

	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("Volcano IDs:", "\n", string(formattedJSON))
}

func Print11LatestNews() {
	news := geonetApi.GetGeonetNews(11, "All")

	formattedJSON, err := json.MarshalIndent(news, "", "  ")

	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("11 latest news:", "\n", string(formattedJSON))
}

func PrintLatestGeohazardEventsNews() {
	news := geonetApi.GetGeonetNews(1, "Geohazard Events")

	formattedJSON, err := json.MarshalIndent(news, "", "  ")

	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}

	fmt.Println("Latest geohazard events news:", "\n", string(formattedJSON))
}
