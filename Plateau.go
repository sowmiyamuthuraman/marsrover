package main

type Plateau struct {
	TopRightCoordinate   Coordinate `json: topRightCoordinate`
	BottomLeftCoordinate Coordinate `json: bottomLeftCoordinate`
}

func GeneratePlateau(topRightCoordinate Coordinate, bottomLeftCoordinate Coordinate) Plateau {
	return Plateau{topRightCoordinate, bottomLeftCoordinate}
}

func (plateau *Plateau) isInRange(coordinate Coordinate) bool {
	return coordinate.isLesserThan(plateau.BottomLeftCoordinate) && coordinate.isGreaterThan(plateau.TopRightCoordinate)
}
