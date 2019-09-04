package marsrover

import (
	"errors"
	"fmt"
	"regexp"
)

type Rover struct {
	plateau     Plateau
	direction   Direction
	coordinates Coordinate
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
	rover.direction = rover.direction.left()
}
func (rover *Rover) turnRight() {
	rover.direction = rover.direction.right()
}

func (rover *Rover) move() {
	coordinatesAfterMove := rover.coordinates.GenerateWithStepSize(rover.direction.stepSizeForX(), rover.direction.stepSizeForY())
	if rover.plateau.isInRange(coordinatesAfterMove) {
		rover.coordinates = coordinatesAfterMove
	}
}

func (rover Rover) toString() string {
	return fmt.Sprint(rover.coordinates, " ", rover.direction.toString())
}
