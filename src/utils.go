package Groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (E *Engine) getAPI(link string) []byte {

	response, err := http.Get(link)

	if err != nil {
        log.Fatal(err)
    }

    data, err := ioutil.ReadAll(response.Body)
    
	if err != nil {
        log.Fatal(err)
    }

	return data

}

func (E *Engine) setArtists(APIData []byte) {
		json.Unmarshal(APIData, &E.Artists)
		for i := 0; i < len(E.Artists); i++ {
			json.Unmarshal(E.getAPI(E.Artists[i].LocationsAPI), &E.Artists[i].Locations)
			json.Unmarshal(E.getAPI(E.Artists[i].ConcertDatesAPI), &E.Artists[i].ConcertDates)
			json.Unmarshal(E.getAPI(E.Artists[i].RelationsAPI), &E.Artists[i].Relations)
		}
}


func (E *Engine) sorting(parameter string) {
	count := len(E.Artists)
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			move := false
			switch parameter {
				case "creation":
					move = E.Artists[i].CreationDate > E.Artists[j].CreationDate
				case "album":
					move = E.atoi(E.Artists[i].FirstAlbum[6:]) > E.atoi(E.Artists[j].FirstAlbum[6:])
				case "member":
					move = len(E.Artists[i].Members) > len(E.Artists[j].Members)
			}
			if (move) {
				E.Artists[i], E.Artists[j] = E.Artists[j], E.Artists[i]
			}
		}
	}
}

func (E *Engine) search(parameter string) {
	var newList []artists
	for i := 0; i < len(E.ArtistsList); i++ {
		if (strings.Contains(strings.ToLower(E.ArtistsList[i].Name), strings.ToLower(parameter))) {
			newList = append(newList, E.ArtistsList[i])
		}
	}
	E.Artists = newList
}

func (E *Engine) filterNumberOfMember(parameter int) {
	var newList []artists
	for i := 0; i < len(E.Artists); i++ {
		if (len(E.Artists[i].Members) == parameter) {
			newList = append(newList, E.Artists[i])
		}
	}
	E.Artists = newList
}

func (E *Engine) filterAlbumDate(parameter int) {
	var newList []artists
	for i := 0; i < len(E.Artists); i++ {
		if (E.atoi(E.Artists[i].FirstAlbum[6:]) == parameter) {
			newList = append(newList, E.Artists[i])
		}
	}
	E.Artists = newList
}

func (E *Engine) atoi(str string) int {
	s, err := strconv.Atoi(str)
	if (err != nil) {
		fmt.Println(err)
		return 0
	}
	return s
}
