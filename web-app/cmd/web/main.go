package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"web-app/internal/httpapp"
	"web-app/internal/store"
)

func main() {
	store := store.NewMemoryStore()

	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	handlers := &httpapp.Handlers{
		Store: store,
		Tmpl:  tmpl,
	}

	router := httpapp.NewRouter(handlers)

	server := &http.Server{
		Addr:              ":3000",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("Web app running at http://localhost:3000")
	log.Fatal(server.ListenAndServe())
}
