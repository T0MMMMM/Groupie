package Groupie

import (
	"html/template"
	"net/http"
)


func (E *Engine) index(w http.ResponseWriter, r *http.Request) {

	sorting := r.FormValue("sorting")
	numberOfMember := r.FormValue("number-of-member")
	albumDate  := r.FormValue("album-date")
	albumDateFilter := r.FormValue("album-date-filter")

	name := r.FormValue("search")
	reload := r.FormValue("reload")

	if name != "" { // BARRE DE RECHERCHE
		E.search(name)
	}

	if sorting != "" { // TRI
		E.sorting(sorting)
	}
	if numberOfMember != "" {
		E.filterNumberOfMember(E.atoi(numberOfMember))
	}
	if albumDateFilter != "" {
		E.filterAlbumDate(E.atoi(albumDate))
	}

	if (reload != "" || (sorting == "" && numberOfMember == "" && name == "" && albumDate == "")) { // REMET LA LIST COMPLET DES ARTISTES
		E.Artists = E.ArtistsList
	}
	
	//E.templateRender(w, "artist", artist)
	E.templateRender(w, "index", "none")

}

func (E *Engine) artist(w http.ResponseWriter, r *http.Request) {
	artist := r.FormValue("artist")
	E.templateRender(w, "artist", artist)
}

func (E *Engine) templateRender(w http.ResponseWriter, tmpl string, artist string) { // affiche la page passer en paramètre
	t, err := template.ParseFiles("./serv/html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if (artist == "none") {
		t.Execute(w, E)
	} else if (artist != "") {
		for i := 0; i < len(E.ArtistsList); i++ {
			if (E.ArtistsList[i].Name == artist) {
				t.Execute(w, E.ArtistsList[i])
			}
		}
	} else {
		t.Execute(w, E.ArtistsList[0])
	}
	
}
