package api

import (
	"log"
	"net/http"
	"strings"
)

const (
	fishWatchURL string = "https://www.fishwatch.gov/api/species"
)

func GetAllFish() *http.Response {
	resp, err := http.Get(fishWatchURL)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func GetFishBySpecies(s string) *http.Response {
	res := strings.ToLower(strings.ReplaceAll(s, " ", "-"))
	resp, err := http.Get(fishWatchURL + "/" + res)
	if err != nil {
		//	assume the fish doesn't exist
		log.Fatalln(err)
	}
	return resp
}
