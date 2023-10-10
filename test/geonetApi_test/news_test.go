package geonetApi_test

import (
	"fmt"
	"geonetApp/pkg/geonetApi"
	"testing"
)

func TestNewsAll(t *testing.T) {
	news := geonetApi.GetGeonetNews(11, "All")

	result := fmt.Sprintf("%d", len(news))
	expected := "11"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestNewsDataBlog(t *testing.T) {
	news := geonetApi.GetGeonetNews(1, "Data Blog")

	result := fmt.Sprintf("%d", len(news))
	expected := "1"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	result = news[0].Tag
	expected = "Data Blog"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestNewsGeohazardEvents(t *testing.T) {
	news := geonetApi.GetGeonetNews(1, "Geohazard Events")

	result := fmt.Sprintf("%d", len(news))
	expected := "1"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	result = news[0].Tag
	expected = "Geohazard Events"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestNewsNews(t *testing.T) {
	news := geonetApi.GetGeonetNews(1, "News")

	result := fmt.Sprintf("%d", len(news))
	expected := "1"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	result = news[0].Tag
	expected = "News"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestNewsVolcanicActivityBulletin(t *testing.T) {
	news := geonetApi.GetGeonetNews(1, "Volcanic Activity Bulletin")

	result := fmt.Sprintf("%d", len(news))
	expected := "1"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	result = news[0].Tag
	expected = "Volcanic Activity Bulletin"

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func BenchmarkNewsAll(b *testing.B) {
	geonetApi.GetGeonetNews(100, "All")
}

func BenchmarkNewsDataBlog(b *testing.B) {
	geonetApi.GetGeonetNews(5, "Data Blog")
}

func BenchmarkNewsGeohazardEvents(b *testing.B) {
	geonetApi.GetGeonetNews(1, "Geohazard Events")
}

func BenchmarkNewsNews(b *testing.B) {
	geonetApi.GetGeonetNews(5, "News")
}

func BenchmarkNewsVolcanicActivityBulletin(b *testing.B) {
	geonetApi.GetGeonetNews(5, "Volcanic Activity Bulletin")
}
