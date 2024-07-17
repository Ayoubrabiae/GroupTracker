package data

type MainType struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type DatesType struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type RelationsType struct {
	Index []struct {
		Id             int               `json:"id"`
		DatesLocations map[string]string `json:"datesLocations"`
	} `json:"index"`
}

var (
	MainData        MainType
	DatesHolder     DatesType
	RelationsHolder RelationsType
)
