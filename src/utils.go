package Groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
					date1, err1 := strconv.Atoi(E.Artists[i].FirstAlbum[6:])
					date2, err2 := strconv.Atoi(E.Artists[j].FirstAlbum[6:])
					if (err1 != nil || err2 != nil) {
						fmt.Println("Error during conversion")
						return
					}
					move = date1 > date2
				case "member":
					move = len(E.Artists[i].Members) > len(E.Artists[j].Members)
			}
			if (move) {
				E.Artists[i], E.Artists[j] = E.Artists[j], E.Artists[i]
			}
		}
	}
}



