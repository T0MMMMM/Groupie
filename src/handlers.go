package Groupie


import (
	"fmt"
	"net/http"
)


func (E *Engine) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
