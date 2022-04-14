package flight

import (
	"encoding/json"
	"errors"
)

const (
	pathDisconnectedErr = "flight path is disconnected"
	pathLoopErr         = "flight path has a loop"
	pathCycleErr        = "flight path has a cycle"
)

// SortPath topoligically sorts the flights and returns the source and
// destination of the sorted flight path as a json array.
func SortPath(flights [][]string) ([]byte, error) {
	// build flight graph and inverse mapping of dest airports
	flightGraph := make(map[string]string)
	destAirports := make(map[string]byte)
	for _, flight := range flights {
		src, dest := flight[0], flight[1]
		if src == dest {
			return nil, errors.New(pathLoopErr)
		}
		if _, ok := flightGraph[src]; ok {
			return nil, errors.New(pathDisconnectedErr)
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
		return nil, errors.New(pathCycleErr)
	} else if len(sources) > 1 {
		return nil, errors.New(pathDisconnectedErr)
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

	if len(sortedFlights) < len(flights) {
		return nil, errors.New(pathDisconnectedErr)
	}
	src, dest := sortedFlights[0][0], sortedFlights[len(sortedFlights)-1][1]
	return json.Marshal([]string{src, dest})
}
