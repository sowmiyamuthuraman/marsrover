package marsrover

type West struct{
	stepForXAxis int
	stepForYAxis int
}
func (west West)right() Direction{
	return Enum.N
}
func (west West) left() Direction{
	return Enum.S
}
func (west West) stepSizeForX() int{
	return west.stepForXAxis
}
func (west West) stepSizeForY() int{
	return west.stepForYAxis
}
func(west West) toString() string{
	return "West"
}
