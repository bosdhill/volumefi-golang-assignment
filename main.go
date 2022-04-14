package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	dataJson := `[["IND", "EWR"],["SFO", "ATL"],["GSO", "IND"],["ATL", "GSO"]]`
	var flights [][]string
	err := json.Unmarshal([]byte(dataJson), &flights)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	flightGraph := make(map[string]string)
	destAirports := make(map[string]string)
	for _, flight := range flights {
		src, dest := flight[0], flight[1]
		flightGraph[src] = dest
		destAirports[dest] = ""
	}

	// find all the airports with zero in bound flights
	var sources []string
	for src := range flightGraph {
		if _, ok := destAirports[src]; !ok {
			sources = append(sources, src)
		}
	}

	if len(sources) > 1 {
		log.Fatalf("flight path is disconnected")
		os.Exit(1)
	}

	var sortedFlights [][]string
	src := sources[0]
	for len(sortedFlights) != len(flights) {
		dest := flightGraph[src]
		sortedFlights = append(sortedFlights, []string{src, dest})
		if _, ok := flightGraph[dest]; ok {
			src = dest
		}
	}

	fmt.Println(sortedFlights[0][0], sortedFlights[len(flights)-1][1])
}
