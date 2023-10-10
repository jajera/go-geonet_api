package geonetApi

const VerboseLogging = false

const (
	GeonetApiUrl = "https://api.geonet.org.nz"

	GeonetNewsApiPath  = "/news/geonet"
	GeonetQuakeApiPath = "/quake"
	GeonetValApiPath   = "/volcano/val"

	GeonetNewsApiUrl  = GeonetApiUrl + GeonetNewsApiPath
	GeonetQuakeApiUrl = GeonetApiUrl + GeonetQuakeApiPath
	GeonetValApiUrl   = GeonetApiUrl + GeonetValApiPath

	GeonetNewsAcceptHeader  = "application/json;version=2"
	GeonetQuakeAcceptHeader = "application/vnd.geo+json;version=2"
	GeonetValAcceptHeader   = "application/vnd.geo+json;version=2"
)
