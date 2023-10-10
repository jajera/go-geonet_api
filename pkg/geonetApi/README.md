# Geonet API

## Logging

VerboseLogging enables or disables verbose logging.

```go
const VerboseLogging = true
```

## News

### Parameters

**recordCount**

Accepts integer number.

**typeFilter**

Accepts one of these values.

	"All",
	"Data Blog",
	"Geohazard Events",
	"News",
	"Volcanic Activity Bulletin",

### Sample Usage
```go
geonetapi.GetGeonetNews(5, "Geohazard Events")
```
