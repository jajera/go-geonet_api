package geonetApi_test

import (
	"fmt"
	"geonetApp/pkg/geonetApi"
	"testing"
)

func TestQuakeMmi3(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"", // publicId
		"", // timeAfter
		-1, // depthAbove
		"", // intensity
		-1, // magnitudeAbove
		3,  // mmi
		"", // localityIn
		"", // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeIntensityMinor(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"",      // publicId
		"",      // timeAfter
		-1,      // depthAbove
		"minor", // intensity
		-1,      // magnitudeAbove
		3,       // mmi
		"",      // localityIn
		"",      // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakePublicId(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"2023p728259", // publicId
		"",            // timeAfter
		-1,            // depthAbove
		"",            // intensity
		-1,            // magnitudeAbove
		3,             // mmi
		"",            // localityIn
		"",            // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeTimeAfter(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"",                         // publicId
		"2023-09-27T10:56:41.659Z", // timeAfter
		-1,                         // depthAbove
		"",                         // intensity
		-1,                         // magnitudeAbove
		3,                          // mmi
		"",                         // localityIn
		"",                         // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeLocalityIn(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"",            // publicId
		"",            // timeAfter
		-1,            // depthAbove
		"",            // intensity
		-1,            // magnitudeAbove
		3,             // mmi
		"Collingwood", // localityIn
		"",            // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeQualityBest(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"",     // publicId
		"",     // timeAfter
		-1,     // depthAbove
		"",     // intensity
		-1,     // magnitudeAbove
		3,      // mmi
		"",     // localityIn
		"best", // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeMagnitudeAbove3(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"", // publicId
		"", // timeAfter
		-1, // depthAbove
		"", // intensity
		3,  // magnitudeAbove
		3,  // mmi
		"", // localityIn
		"", // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}

func TestQuakeDepth5(t *testing.T) {
	quakeAll := geonetApi.GetGeonetQuakes(
		"", // publicId
		"", // timeAfter
		5,  // depthAbove
		"", // intensity
		-1, // magnitudeAbove
		3,  // mmi
		"", // localityIn
		"", // quality
	)

	expected := 0
	result := len(quakeAll)

	if result <= expected {
		t.Errorf("Expected %s is greater than %s",
			fmt.Sprint(result),
			fmt.Sprint(expected))
	}
}
