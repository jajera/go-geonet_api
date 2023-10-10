package geonetApi_test

import (
	"fmt"
	"geonetApp/pkg/geonetApi"
	"testing"
)

func TestVolcanoAlertLevelAll(t *testing.T) {
	volcanoAlertLevel := geonetApi.GetGeonetVolcanoAlertLevel(
		"",    // volcanoID
		"",    // volcanoTitle
		1,     // int
		"All", // acc
		"",    // activity
		"",    // hazards
	)

	expected := 0
	result := len(volcanoAlertLevel)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestVolcanoAlertLevelFilterId(t *testing.T) {
	volcanoAlertLevel := geonetApi.GetGeonetVolcanoAlertLevel(
		"whiteisland", // volcanoID
		"",            // volcanoTitle
		-1,            // int
		"All",         // acc
		"",            // activity
		"",            // hazards
	)

	expected := "whiteisland"
	result := volcanoAlertLevel[0].Properties.VolcanoID

	if result != expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestVolcanoAlertLevelFilterAcc(t *testing.T) {
	volcanoAlertLevel := geonetApi.GetGeonetVolcanoAlertLevel(
		"",      // volcanoID
		"",      // volcanoTitle
		-1,      // int
		"Green", // acc
		"",      // activity
		"",      // hazards
	)

	expected := 0
	result := len(volcanoAlertLevel)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestVolcanoIds(t *testing.T) {
	volcanoIds := geonetApi.GetGeonetVolcanoIds()

	expected := 0
	result := len(volcanoIds)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}
