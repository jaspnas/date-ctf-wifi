package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"notesApi/common"
)

func NotesGetHandler(w http.ResponseWriter, r *http.Request) {
	notes := common.GetNotes()

	byteData, err := json.Marshal(notes)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(byteData)
	if err != nil {
		log.Println(err)
	}

}
