package main

import (
	"fmt"
	"math"
)

type Game struct {
	Canvas *Canvas
	Cells  []*Cell
}

func NewGame(width int, height int) (game *Game) {
	var canvas *Canvas
	var cells []*Cell

	var cellTotal = width * height

	for i := 0; i < cellTotal; i++ {
		var coordinate = &Coordinate{
			X: i % width,
			Y: int(math.Floor(float64(i) / float64(width))),
		}

		fmt.Println(coordinate.X, coordinate.Y)
		cells = append(cells, NewCell(true, coordinate))
	}

	canvas = NewCanvas(width, height)

	return &Game{Canvas: canvas, Cells: cells}
}

func (game Game) Run() (err error) {
	fmt.Println(game.Cells)
	return
}
