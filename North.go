package main

type North struct{
	stepForXAxis int
	stepForYAxis int
}
func (north North)right() Direction{
	return Enum.E
}
func (north North)left() Direction{
	return Enum.W
}
func (north North) stepSizeForX() int{
	return north.stepForXAxis
}
func (north North) stepSizeForY() int{
	return north.stepForYAxis
}
func(north North) toString() string{
	return "North"
}
