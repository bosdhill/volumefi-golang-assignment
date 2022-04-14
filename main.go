package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func sortFlights(flights [][]string) ([]byte, error) {
	// build flight graph and inverse mapping of dest airports
	flightGraph := make(map[string]string)
	destAirports := make(map[string]byte)
	for _, flight := range flights {
		src, dest := flight[0], flight[1]
		if src == dest {
			return nil, errors.New("flight path has a loop")
		}
		if _, ok := flightGraph[src]; ok {
			return nil, errors.New("flight path is disconnected")
		}
		flightGraph[src] = dest
		destAirports[dest] = 0
	}

	// find all the airports with zero in bound flights
	var sources []string
	for src := range flightGraph {
		if _, ok := destAirports[src]; !ok {
			sources = append(sources, src)
		}
	}

	if len(sources) == 0 {
		return nil, errors.New("flight path contains a cycle")
	}
	if len(sources) > 1 {
		return nil, errors.New("flight path is disconnected")
	}

	// "visit" all the airports starting from the source airport
	var sortedFlights [][]string
	src := sources[0]
	for {
		dest := flightGraph[src]
		sortedFlights = append(sortedFlights, []string{src, dest})
		if _, ok := flightGraph[dest]; ok {
			src = dest
		} else {
			break
		}
	}

	if len(sortedFlights) != len(flights) {
		return nil, errors.New("flight path contains a cycle")
	}

	src, dest := sortedFlights[0][0], sortedFlights[len(sortedFlights)-1][1]
	return json.Marshal([]string{src, dest})
}

func main() {
	// dataJson := `[["IND", "EWR"],["SFO", "ATL"],["GSO", "IND"],["ATL", "GSO"]]`
	dataJson := `[["EWR", "JFK"],["SFO", "EWR"],["JFK", "SFO"]]`
	var flights [][]string
	err := json.Unmarshal([]byte(dataJson), &flights)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	b, err := sortFlights(flights)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(string(b))
}
