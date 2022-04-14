package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"volumefi-golang-assignment/flight"
)

// decodeJson reads in a json from stdin and decodes it.
func decodeJson() ([][]string, error) {
	var jsonData [][]string
	err := json.NewDecoder(os.Stdin).Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func main() {
	flights, err := decodeJson()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	srcDest, err := flight.SortPath(flights)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(string(srcDest))
	os.Exit(0)
}
