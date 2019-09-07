package main

type Direction interface {
	right() Direction
	left() Direction
	stepSizeForX() int
	stepSizeForY() int
	toString() string
}

type direction struct {
	W Direction
	E Direction
	N Direction
	S Direction
}

var Enum = &direction{
	W: West{-1, 0},
	E: East{1, 0},
	N: North{0, 1},
	S: South{0, -1},
}
