package geonetApi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FeedItem struct {
	Title     string `json:"title"`
	Type      string `json:"type"`
	Tag       string `json:"tag"`
	Val       int    `json:"val"`
	Published string `json:"published"`
	Link      string `json:"link"`
	MLink     string `json:"mlink"`
}

type Feed struct {
	Page  int        `json:"page"`
	Total int        `json:"total"`
	Feed  []FeedItem `json:"feed"`
}

var AllowedTypedFilter = []string{
	"All",
	"Data Blog",
	"Geohazard Events",
	"News",
	"Volcanic Activity Bulletin",
}

func GetGeonetNews(recordCount int, typeFilter string) []FeedItem {
	traverseCount := 0
	pageNoCount := 1
	var newsData HTTPGetResult
	var feedItems []FeedItem

	if !strings.Contains(strings.Join(AllowedTypedFilter, ","), typeFilter) {
		fmt.Println("Invalid typeFilter:", typeFilter)
		fmt.Println("Allowed typeFilter:", strings.Join(AllowedTypedFilter, ", "))
		return feedItems
	}

	for traverseCount < recordCount {
		newGeonetNewsApiUrl := GeonetNewsApiUrl + "?page=" +
			fmt.Sprintf("%d", pageNoCount)
		newsData = HTTPGet(newGeonetNewsApiUrl, GeonetNewsAcceptHeader)

		var feedData Feed
		err := json.Unmarshal([]byte(newsData.Data), &feedData)
		if err != nil {
			fmt.Println("Error:", err)
			return feedItems
		}

		for _, item := range feedData.Feed {
			if traverseCount >= recordCount {
				break
			}

			if typeFilter == "All" ||
				typeFilter == item.Tag {
				feedItem := FeedItem{
					Title:     item.Title,
					Type:      item.Type,
					Tag:       item.Tag,
					Val:       item.Val,
					Published: item.Published,
					Link:      item.Link,
					MLink:     item.MLink,
				}
				feedItems = append(feedItems, feedItem)

				traverseCount++
			}
		}
		if pageNoCount > feedData.Total ||
			pageNoCount >= 100 {
			break
		}
		pageNoCount++
	}

	return feedItems
}
