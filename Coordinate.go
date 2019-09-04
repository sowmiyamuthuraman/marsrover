package marsrover

type Coordinate struct {
	xPosition int
	yPosition int
}

func (coordinate Coordinate) GenerateWithStepSize(x int, y int) Coordinate {
	return GenerateCoordinate(coordinate.xPosition+x, coordinate.yPosition+y)
}

func GenerateCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}

func (coordinate Coordinate) isLesserThan(otherCoordinate Coordinate) bool {
	return (coordinate.xPosition <= otherCoordinate.xPosition) && (coordinate.yPosition <= otherCoordinate.yPosition)
}

func (coordinate Coordinate) isGreaterThan(otherCoordinate Coordinate) bool {
	return (coordinate.xPosition >= otherCoordinate.xPosition) && (coordinate.yPosition >= otherCoordinate.yPosition)
}
