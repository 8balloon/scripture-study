package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type SJVerse struct {
	Text  string
	Verse int
}

type SJChapter struct {
	Chapter int
	Verses  []SJVerse
}

type SJBook struct {
	Book     string
	Chapters []SJChapter
}

type ScriptureJSON struct {
	Books []SJBook
}

const POGPPath = "./scriptures-json/pearl-of-great-price.json"

func GetPOGPSJ() *ScriptureJSON {
	f, err := os.Open(POGPPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var pogpJSON ScriptureJSON
	json.Unmarshal(bytes, &pogpJSON)
	return &pogpJSON
}
