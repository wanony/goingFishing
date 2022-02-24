package main

import (
	"awesomeProject/api"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp := api.GetAllFish()
	//fmt.Println(resp)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		bodyStr := string(bodyBytes)
		log.Println(bodyStr)
	}
	specResp := api.GetFishBySpecies("Pacific Cod")
	fmt.Println(specResp)
}
