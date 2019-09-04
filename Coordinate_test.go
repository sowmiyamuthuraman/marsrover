package marsrover

import "testing"

var isLesserThantestCases = []struct {
	coordinate      Coordinate
	otherCoordinate Coordinate
	expectedResult  bool
}{
	{
		Coordinate{1, 1},
		Coordinate{2, 1},
		true,
	}, {
		Coordinate{3, 1},
		Coordinate{2, 1},
		false,
	},
}
var isGreaterThantestCases = []struct {
	coordinate      Coordinate
	otherCoordinate Coordinate
	expectedResult  bool
}{
	{
		Coordinate{3, 3},
		Coordinate{2, 1},
		true,
	}, {
		Coordinate{1, 1},
		Coordinate{2, 1},
		false,
	},
}

func TestIsLessThan(t *testing.T) {
	for _, testCase := range isLesserThantestCases {
		result := testCase.coordinate.isLesserThan(testCase.otherCoordinate)
		if result != testCase.expectedResult {
			t.Error("expected ", testCase.expectedResult, "got ", result)
		}
	}
}

func TestIsGreaterThan(t *testing.T) {
	for _, testCase := range isGreaterThantestCases {
		result := testCase.coordinate.isGreaterThan(testCase.otherCoordinate)
		if result != testCase.expectedResult {
			t.Error("expected ", testCase.expectedResult, "got ", result)
		}
	}
}
