package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"notesApi/common"
)

type flagResponse struct {
	Flag string `json:"flag"`
}

func HandleFlag(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	flag := common.GetFlag()
	res := flagResponse{flag}

	byteData, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	_, err = w.Write(byteData)
	if err != nil {
		log.Println(err)
	}
}
