package marsrover

import (
	. "github.com/golang/mock/gomock"
	"testing"
)

func TestLeftRotation(t *testing.T) {
	rover := getRover(direction{}.E)
	result, _ := rover.Run("L")
	expectedResult := "{3 3} North"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestRightRotation(t *testing.T) {
	rover := getRover(direction{}.E)
	result, _ := rover.Run("R")
	expectedResult := "{3 3} South"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}
func TestMove(t *testing.T) {
	rover := getRover(direction{}.E)
	result, _ := rover.Run("M")
	expectedResult := "{4 3} East"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestRunWithMultipleInstructions(t *testing.T) {
	t.Run("multiple instructions with move and right command", func(t *testing.T) {
		rover := getRover(direction{}.E)
		result, _ := rover.Run("MMRMMRMRRM")
		expectedResult := "{5 1} East"
		if result != expectedResult {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})
	t.Run("multiple instructions with combinations of move right and left command", func(t *testing.T) {
		controller := NewController(t) //tracks all  calls to the mock
		defer controller.Finish()
		rover := getRover(NewMockDirection(controller))
		result, _ := rover.Run("LMLMLMLMM")
		expectedResult := "{4 3} East"
		if result != expectedResult {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})

}

func TestUnitLeftRotation(t *testing.T) {
	controller := NewController(t) //tracks all  calls to the mock
	defer controller.Finish()
	mockDirection := NewMockDirection(controller)
	rover := getRover(mockDirection)
	mockDirection.EXPECT().left().Return(mockDirection).Times(1)
	mockDirection.EXPECT().toString().Return("MockDirection").Times(1)
	result, _ := rover.Run("L")
	expectedResult := "{3 3} MockDirection"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestRunReturnsErrorForInvalidCommand(t *testing.T) {
	rover := getRover(direction{}.E)
	_, err := rover.Run("C")
	if err == nil {
		t.Errorf("Expected error but no error was caught")
	}
}

func getRover(direction Direction) Rover {
	topRight := Coordinate{0, 0}
	bottomLeft := Coordinate{5, 5}
	plateau := GeneratePlateau(topRight, bottomLeft)
	coordinate := Coordinate{3, 3}
	return GenerateRover(plateau, coordinate, direction)
}
