package main

type Canvas struct {
	Width  int
	Height int
}

func NewCanvas(width int, height int) (canvas *Canvas) {
	return &Canvas{Width: width, Height: height}
}
