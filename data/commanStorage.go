package data

type MainType struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type ArtistType struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      string `json:"members"`
	CreationDate int    `json:"creationData"`
	FirstAlbum   string `json:"firstAlbum"`
}

type LocationsType struct {
	Locations string `json:"locations"`
	Dates string `json:"dates"`
}

var (

	Artist   ArtistType
	Locations LocationsType
	MainData MainType
)
