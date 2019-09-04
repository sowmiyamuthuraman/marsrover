package main

import (
	"fmt"
	"marsRover"
)

func main() {
	topRight := marsrover.GenerateCoordinate(0, 0)
	bottomLeft := marsrover.GenerateCoordinate(5, 5)
	plateau := marsrover.GeneratePlateau(topRight, bottomLeft)
	coordinate := marsrover.GenerateCoordinate(3, 3)
	rover := marsrover.GenerateRover(plateau, coordinate, marsrover.Enum.E)
	commands := "MMRMMRMRRM"
	result, err := rover.Run(commands)
	if (err != nil) {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)
}
