package main

import (
	"app/utils"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func playRound(rw http.ResponseWriter, r *http.Request) {
	buttonsSelected := r.URL.Query().Get("buttonsSelected")
	log.Println(buttonsSelected)
	result := utils.PlayRound(buttonsSelected)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(out)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/playRound", playRound)
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}
