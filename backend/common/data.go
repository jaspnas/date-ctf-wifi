package common

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Data struct {
	Password string `json:"password"`
	Flag     string `json:"flag"`
	Notes    []Note `json:"notes"`
}

func readData() Data {
	jsonFile, err := os.Open("./data.json")
	if err != nil {
		log.Println(err)
	}

	byteData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
	}
	err = jsonFile.Close()
	if err != nil {
		log.Println(err)
	}
	var data Data
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		log.Println(err)
	}
	return data
}

func GetNotes() []Note {
	return readData().Notes
}

func GetPassword() string {
	return readData().Password
}

func GetFlag() string {
	return readData().Flag
}
