package main

import (
	. "github.com/golang/mock/gomock"
	"testing"
)

func TestUnitLeftRotation(t *testing.T) {
	/*Arrange*/
	controller := NewController(t) //tracks all  calls to the mock
	defer controller.Finish()
	mockDirection := NewMockDirection(controller)
	rover := getRover(mockDirection)
	mockDirection.EXPECT().left().Return(mockDirection).Times(1)
	mockDirection.EXPECT().toString().Return("MockDirection").Times(1)
	/*Act*/
	result, _ := rover.Run("L")
	/*Assert*/
	expectedResult := "{3 3} MockDirection"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestUnitRightRotation(t *testing.T) {
	/*Arrange*/
	controller := NewController(t) //tracks all  calls to the mock
	defer controller.Finish()
	mockDirection := NewMockDirection(controller)
	rover := getRover(mockDirection)
	mockDirection.EXPECT().right().Return(mockDirection).Times(1)
	mockDirection.EXPECT().toString().Return("MockDirection").Times(1)
	/*Act*/
	result, _ := rover.Run("R")
	/*Assert*/
	expectedResult := "{3 3} MockDirection"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestUnitMove(t *testing.T) {
	/*Arrange*/
	controller := NewController(t) //tracks all  calls to the mock
	defer controller.Finish()
	mockDirection := NewMockDirection(controller)
	rover := getRover(mockDirection)
	mockDirection.EXPECT().stepSizeForX().Return(-1).Times(1)
	mockDirection.EXPECT().stepSizeForY().Return(0).Times(1)
	mockDirection.EXPECT().toString().Return("MockDirection").Times(1)
	/*Act*/
	result, _ := rover.Run("M")
	/*Assert*/
	expectedResult := "{2 3} MockDirection"
	if result != expectedResult {
		t.Errorf("Expected %s but got %s", expectedResult, result)
	}
}

func TestRunWithMultipleInstructions(t *testing.T) {
	/*Arrange*/
	rover := getRover(Enum.E)
	t.Run("multiple instructions with move and right command", func(t *testing.T) {
		/*Act*/
		result, _ := rover.Run("MMRMMRMRRM")
		/*Assert*/
		expectedResult := "{5 1} East"
		if result != expectedResult {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})
	t.Run("multiple instructions with combinations of move right and left command", func(t *testing.T) {
		/*Act*/
		result, _ := rover.Run("LMLMLMLMM")
		expectedResult := "{4 3} East"
		/*Assert*/
		if result != expectedResult {
			t.Errorf("Expected %s but got %s", expectedResult, result)
		}
	})

}

func TestRunReturnsErrorForInvalidCommand(t *testing.T) {
	/*Arrange*/
	controller := NewController(t) //tracks all  calls to the mock
	defer controller.Finish()
	rover := getRover(NewMockDirection(controller))
	/*Act*/
	_, err := rover.Run("C")
	/*Assert*/
	if err == nil {
		t.Errorf("Expected error but no error was caught")
	}
}

func getRover(direction Direction) Rover {
	var topRight Coordinate = Coordinate{0, 0}
	bottomLeft := Coordinate{5, 5}
	plateau := GeneratePlateau(topRight, bottomLeft)
	coordinate := Coordinate{3, 3}
	return GenerateRover(plateau, coordinate, direction)
}
