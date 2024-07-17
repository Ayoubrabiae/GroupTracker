package data

type MainType struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type ArtistType struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationData"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationsType struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
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
	Artist    []ArtistType
	Locations LocationsType
	MainData  MainType
	Dates     DatesType
	Relations RelationsType
)
