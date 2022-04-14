package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"volumefi-golang-assignment/flight"
)

const defaultJsonPath = "testdata/input.json"

// parseJsonFile reads in a json file and parses it.
func parseJsonFile(jsonFilePath string) ([][]string, error) {
	jsonFile, err := os.Open(jsonFilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	b, _ := io.ReadAll(jsonFile)
	var flights [][]string
	err = json.Unmarshal(b, &flights)
	if err != nil {
		return nil, err
	}
	return flights, nil
}

func main() {
	var jsonFile string
	flag.StringVar(&jsonFile, "jsonFile", defaultJsonPath, "input json file path")

	flights, err := parseJsonFile(jsonFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	srcDest, err := flight.SortPath(flights)
	// flights.SortPath(flights)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(string(srcDest))
}
