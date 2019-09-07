package main

type South struct {
	stepForXAxis int
	stepForYAxis int
}

func (south South) right() Direction {
	return Enum.W
}
func (south South) left() Direction {
	return Enum.E
}
func (south South) stepSizeForX() int {
	return south.stepForXAxis
}
func (south South) stepSizeForY() int {
	return south.stepForYAxis
}
func(south South) toString() string{
	return "South"
}

