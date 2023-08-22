package main

import (
	"log"
	"net/http"
	"notesApi/handlers"
)

func main() {

	http.HandleFunc("/api/auth", handlers.AuthPostHandler)
	http.HandleFunc("/api/notes", handlers.NotesGetHandler)
	http.HandleFunc("/api/flag", handlers.HandleFlag)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}

}
