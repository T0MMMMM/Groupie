package Groupie

import (
	"fmt"
	// "math/rand"
	"net/http"
	"strings"
	"sort"
)

type Engine struct {
	Port		string
	Artists 	[]artists
	ArtistsList []artists
	Info		APIinfo
}

type APIinfo struct {
	MaxAlbumDate int
	MinAlbumDate int
	MaxMemberNumber int
	MaxMemberNumberList []int // List de nombre de la taille du nombre max de membres
	MaxCreationDate int
	MinCreationDate int
	ConcertLocations []string
	Country string
}

type artists struct {
	Id		int			`json:"id"`
	Image	string		`json:"image"`
	Name	string		`json:"name"`
	Members	[]string	`json:"members"`
	CreationDate int	`json:"creationDate"`
	FirstAlbum string	`json:"firstAlbum"`
	LocationsAPI string	`json:"locations"`
	ConcertDatesAPI string	`json:"concertDates"`
	RelationsAPI string	`json:"relations"`
	Locations locations
	ConcertDates dates
	Relations relations
}

type locations struct {
	Id			int				`json:"id"`
	Locations	[]string		`json:"locations"`
	Dates		string			`json:"dates"`
	Country []string
	ConcertCityLocations []string
}

type dates struct {
	Id			int				`json:"id"`
	Dates		[]string		`json:"dates"`
}

type relations struct {
	Id			int						`json:"id"`
	Relations	map[string][]string		`json:"datesLocations"`
}


func (E *Engine) Init() {
	//rand.Seed(time.Now().UnixNano())

	E.Port = ":8080"

	E.setArtists(E.getAPI("https://groupietrackers.herokuapp.com/api/artists"))

	E.ArtistsList = E.Artists

	E.Info.MinAlbumDate = E.atoi(E.Artists[0].FirstAlbum[6:])
	E.Info.MinCreationDate = E.Artists[0].CreationDate

	for i := 0; i < len(E.Artists); i++ {
		if E.Info.MinAlbumDate > E.atoi(E.Artists[i].FirstAlbum[6:]) {
			E.Info.MinAlbumDate = E.atoi(E.Artists[i].FirstAlbum[6:])
		}
		if E.Info.MaxAlbumDate < E.atoi(E.Artists[i].FirstAlbum[6:]) {
			E.Info.MaxAlbumDate = E.atoi(E.Artists[i].FirstAlbum[6:])
		}
		if E.Info.MaxMemberNumber < len(E.Artists[i].Members) {
			E.Info.MaxMemberNumber = len(E.Artists[i].Members)
		}
		if E.Info.MinCreationDate > E.Artists[i].CreationDate {
			E.Info.MinCreationDate = E.Artists[i].CreationDate
		}
		if E.Info.MaxCreationDate < E.Artists[i].CreationDate {
			E.Info.MaxCreationDate = E.Artists[i].CreationDate
		}
		for j := 0; j < len(E.Artists[i].Locations.Locations); j++ {
			E.Info.ConcertLocations = append(E.Info.ConcertLocations, strings.Title(strings.ReplaceAll(strings.Split(E.Artists[i].Locations.Locations[j], "-")[1], "_", " ")))
			E.Artists[i].Locations.Country = append(E.Artists[i].Locations.Country, strings.Title(strings.ReplaceAll(strings.Split(E.Artists[i].Locations.Locations[j], "-")[1], "_", " ")))
			E.Artists[i].Locations.ConcertCityLocations = append(E.Artists[i].Locations.ConcertCityLocations, strings.Split(E.Artists[i].Locations.Locations[j], "-")[0] ) 
		} // LAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
	}

	for i := 1; i <= E.Info.MaxMemberNumber; i++ {
		E.Info.MaxMemberNumberList = append(E.Info.MaxMemberNumberList, i)
	}

	E.Info.ConcertLocations = removeDuplicateValues(E.Info.ConcertLocations)

	sort.Strings(E.Info.ConcertLocations)

}



func (E *Engine) Run() {
	E.Init()

	fs := http.FileServer(http.Dir("serv"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	http.HandleFunc("/", E.index)
	http.HandleFunc("/artist", E.artist)

	fmt.Println("(http://localhost:8080) - Serveur started on port", E.Port)

	http.ListenAndServe(E.Port, nil)
}