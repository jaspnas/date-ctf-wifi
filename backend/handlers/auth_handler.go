package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"notesApi/common"
)

type AuthReq struct {
	Password string `json:"password"`
}

func AuthPostHandler(w http.ResponseWriter, r *http.Request) {
	var authReq AuthReq

	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(byteData, &authReq)
	if err != nil {
		log.Println(err)
	}

	// Very secure storage of password here, I don't fucking care
	if authReq.Password == common.GetPassword() {
		http.Redirect(w, r, "/notes", http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}
