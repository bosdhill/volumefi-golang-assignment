package flight

import (
	"bytes"
	"testing"
)

func TestSortPath(t *testing.T) {
	tests := []struct {
		name    string
		flights [][]string
		want    []byte
	}{
		// Happy paths :D
		{
			name: "one flight",
			flights: [][]string{
				{"SFO", "ATL"},
			},
			want: []byte(`["SFO","ATL"]`),
		},
		{
			name: "two sorted connected flights",
			flights: [][]string{
				{"SFO", "ATL"},
				{"ATL", "EWR"},
			},
			want: []byte(`["SFO","EWR"]`),
		},
		{
			name: "two unsorted connected flights",
			flights: [][]string{
				{"ATL", "EWR"},
				{"SFO", "ATL"},
			},
			want: []byte(`["SFO","EWR"]`),
		},
		{
			name: "connected flights with no cycles",
			flights: [][]string{
				{"IND", "EWR"},
				{"SFO", "ATL"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			want: []byte(`["SFO","EWR"]`),
		},
		// Unhappy paths :C
		{
			name: "one flight with a loop",
			flights: [][]string{
				{"SFO", "SFO"},
			},
			want: nil,
		},
		{
			name: "connected flights with a loop",
			flights: [][]string{
				{"JFK", "LAX"},
				{"SFO", "SFO"},
				{"LAX", "ATL"},
				{"ATL", "SFO"},
			},
			want: nil,
		},
		{
			name: "two unsorted disconnected flights",
			flights: [][]string{
				{"ATL", "EWR"},
				{"SFO", "JFK"},
			},
			want: nil,
		},
		{
			name: "connected flights with a cycle",
			flights: [][]string{
				{"EWR", "JFK"},
				{"SFO", "EWR"},
				{"JFK", "SFO"},
			},
			want: nil,
		},
		{
			name: "disconnected flights with multiple unique zero inbound sources",
			flights: [][]string{
				{"IND", "EWR"},
				{"ATL", "SFO"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			want: nil,
		},
		{
			name: "disconnected flights with multiple non-unique zero inbound sources",
			flights: [][]string{
				{"IND", "EWR"},
				{"ATL", "SFO"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
			},
			want: nil,
		},
		{
			name: "two connected flights with a cycle",
			flights: [][]string{
				{"JFK", "HNL"},
				{"HNL", "JFK"},
			},
			want: nil,
		},
		{
			name: "disconnected flights with a cycle",
			flights: [][]string{
				{"IND", "EWR"},
				{"GSO", "IND"},
				{"ATL", "GSO"},
				{"JFK", "HNL"},
				{"HNL", "JFK"},
			},
			want: nil,
		},
		{
			name: "disconnected flights with non-unique destinations",
			flights: [][]string{
				{"JFK", "ATL"},
				{"SFO", "LAX"},
				{"LAX", "ATL"},
			},
			want: nil,
		},
	}

	for _, tc := range tests {
		got, _ := SortPath(tc.flights)

		if bytes.Compare(got, tc.want) != 0 {
			t.Fatalf("%v expected: %v, got: %v", tc.name, string(tc.want), string(got))
		}
	}
}
