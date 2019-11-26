package locations

type Country struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	TimeZone       string `json:"time_zone"`
	GeoInformation GeoInformation
	State          []State
}

type GeoInformation struct {
	location GeoLocations `json:"location"`
}

type GeoLocations struct {
	latitude  float64 `json:"latitude"`
	longitude float64 `json:"longitude"`
}

type State struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
