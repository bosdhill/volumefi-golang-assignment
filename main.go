package main

import "fmt"

func main() {
	flightGraph := make(map[string]string)
	flights := [][]string{
		{"IND", "EWR"},
		{"SFO", "ATL"},
		{"GSO", "IND"},
		{"ATL", "GSO"},
	}

	destAirports := make(map[string]string)
	for _, flight := range flights {
		src, dest := flight[0], flight[1]
		flightGraph[src] = dest
		destAirports[dest] = ""
	}

	// find the src, or the one with zero in bound edges
	var src string
	for s := range flightGraph {
		if _, ok := destAirports[s]; !ok {
			src = s
		}
	}

	var sortedFlights [][]string
	for len(sortedFlights) != len(flights) {
		dest := flightGraph[src]
		sortedFlights = append(sortedFlights, []string{src, dest})
		if _, ok := flightGraph[dest]; ok {
			src = dest
		}
	}

	fmt.Println(sortedFlights[0][0], sortedFlights[len(flights)-1][1])
}
