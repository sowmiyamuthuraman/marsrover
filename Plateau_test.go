package main

import (
	"testing"
)

var testCases = []struct {
	coordinate     Coordinate
	expectedResult bool
}{{
	//coordinate is within bound
	Coordinate{1, 2},
	true,
},
	{
		Coordinate{-1, -2},
		false,
	},
	{
		Coordinate{6, 5},
		false,
	},
}

func TestIsinRange(t *testing.T) {
	topRight := Coordinate{0, 0}
	bottomLeft := Coordinate{5, 5}
	plateau := GeneratePlateau(topRight, bottomLeft)
	for _, testCase := range testCases {
		result := plateau.isInRange(testCase.coordinate)
		if result != testCase.expectedResult {
			t.Error("expected ", testCase.expectedResult, "got ", result)
		}
	}
}
