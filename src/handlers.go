package Groupie

import (
	"html/template"
	"net/http"
)




func (E *Engine) index(w http.ResponseWriter, r *http.Request) {

	sorting := r.FormValue("sorting")

	if sorting != "" {
		E.sorting(sorting)
	}

	E.templateRender(w, "index")

}

func (E *Engine) templateRender(w http.ResponseWriter, tmpl string) { // affiche la page passer en paramètre
	t, err := template.ParseFiles("./serv/html/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, E)
}
