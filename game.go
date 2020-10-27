package main

type Game struct {
	Canvas *Canvas
	Grid   *Grid
}

func NewGame(width int, height int) (game *Game) {
	var canvas = NewCanvas(width, height)
	var grid = NewGrid(width, height)
	return &Game{Canvas: canvas, Grid: grid}
}

func (game Game) Run() (err error) {
	//fmt.Println(game.Cells)
	return
}
