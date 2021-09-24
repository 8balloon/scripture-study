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

type Meta struct {
	Path string
	Name string
}

var Metas = []Meta{
	{
		Path: "./scriptures-json/old-testament.json",
		Name: "The Old Testament",
	},
	{
		Path: "./scriptures-json/new-testament.json",
		Name: "The New Testament",
	},
	{
		Path: "./scriptures-json/book-of-mormon.json",
		Name: "The Book of Mormon",
	},
	{
		Path: "./scriptures-json/doctrine-and-covenants.json",
		Name: "The Doctrine and Covenants",
	},
	{
		Path: "./scriptures-json/pearl-of-great-price.json",
		Name: "The Pearl of Great Price",
	},
}

type Scripture struct {
	Meta Meta
	JSON ScriptureJSON
}

func GetScriptures() []Scripture {
	var scriptures []Scripture
	for _, meta := range Metas {
		f, err := os.Open(meta.Path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		bytes, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
		var scripture Scripture
		scripture.Meta = meta
		json.Unmarshal(bytes, &scripture.JSON)
		scriptures = append(scriptures, scripture)
	}
	return scriptures
}
