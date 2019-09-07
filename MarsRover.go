package main

import (
	"errors"
	"fmt"
	"regexp"
)

type Rover struct {
	Plateau    Plateau    `json:"plateau"`
	Direction  Direction  `json:"direction"`
	Coordinate Coordinate `json:"coordinates"`
}

func GenerateRover(plateau Plateau, coordinate Coordinate, direction Direction) Rover {
	return Rover{plateau, direction, coordinate}
}

func (rover *Rover) Run(s string) (string, error) {
	var isAValidCommand = regexp.MustCompile(`^[LMR]+$`).MatchString
	if !isAValidCommand(s) {
		return "", errors.New("enter a valid command")
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'L':
			rover.turnLeft()
		case 'R':
			rover.turnRight()
		case 'M':
			rover.move()
		default:
		}
	}
	return rover.toString(), nil

}

func (rover *Rover) turnLeft() {
	rover.Direction = rover.Direction.left()
}
func (rover *Rover) turnRight() {
	rover.Direction = rover.Direction.right()
}

func (rover *Rover) move() {
	coordinatesAfterMove := rover.Coordinate.GenerateWithStepSize(rover.Direction.stepSizeForX(), rover.Direction.stepSizeForY())
	if rover.Plateau.isInRange(coordinatesAfterMove) {
		rover.Coordinate = coordinatesAfterMove
	}
}

func (rover Rover) toString() string {
	return fmt.Sprint(rover.Coordinate, " ", rover.Direction.toString())
}
