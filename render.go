package main

import (
	"encoding/json"
	"net/http"
)

type Libro struct {
	Título string `json:"título"`
	Autor  string `json:"autor"`
}

func main() {
	http.HandleFunc("/", MuestraLibros)
	http.ListenAndServe(":8080", nil)
}

func MuestraLibros(w http.ResponseWriter, r *http.Request) {
	libro := Libro{"Construyendo Aplicaciones Web con Go", "Jeremy Saenz"}

	js, err := json.Marshal(libro)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
