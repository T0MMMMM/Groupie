package Groupie

import (
	"fmt"
	// "math/rand"
	"net/http"
	// "time"
)

type Engine struct {
	Port	string
	Artists []artists
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
}



func (E *Engine) Run() {
	E.Init()

	fs := http.FileServer(http.Dir("serv"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	http.HandleFunc("/", E.index)
	//http.HandleFunc("/hangman", E.Hangman)

	fmt.Println("(http://localhost:8080) - Serveur started on port", E.Port)

	http.ListenAndServe(E.Port, nil)
}