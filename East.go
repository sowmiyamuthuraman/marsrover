package main

type East struct {
	stepForXAxis int
	stepForYAxis int
}

func (east East) right() Direction {
	return Enum.S
}
func (east East) left() Direction {
	return Enum.N
}
func (east East) stepSizeForX() int {
	return east.stepForXAxis
}
func (east East) stepSizeForY() int {
	return east.stepForYAxis
}
func (east East) toString() string {
	return "East"
}
