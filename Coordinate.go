package main

type Coordinate struct {
	XPosition int `json: xPosition`
	YPosition int `json: yPosition`
}

func (coordinate Coordinate) GenerateWithStepSize(x int, y int) Coordinate {
	return GenerateCoordinate(coordinate.XPosition+x, coordinate.YPosition+y)
}

func GenerateCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}

func (coordinate Coordinate) isLesserThan(otherCoordinate Coordinate) bool {
	return (coordinate.XPosition <= otherCoordinate.XPosition) && (coordinate.YPosition <= otherCoordinate.YPosition)
}

func (coordinate Coordinate) isGreaterThan(otherCoordinate Coordinate) bool {
	return (coordinate.XPosition >= otherCoordinate.XPosition) && (coordinate.YPosition >= otherCoordinate.YPosition)
}
