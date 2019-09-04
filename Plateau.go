package marsrover

type Plateau struct {
	topRightCoordinate   Coordinate
	bottomLeftCoordinate Coordinate
}


func GeneratePlateau(topRightCoordinate Coordinate, bottomLeftCoordinate Coordinate) Plateau{
	return Plateau{topRightCoordinate,bottomLeftCoordinate}
}

func (plateau *Plateau) isInRange(coordinate Coordinate) bool {
	return coordinate.isLesserThan(plateau.bottomLeftCoordinate) && coordinate.isGreaterThan(plateau.topRightCoordinate)
}
