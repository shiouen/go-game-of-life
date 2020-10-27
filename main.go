package main

import "log"

func main() {
	var game = NewGame(1000, 1000)

	if err := game.Run(); err != nil {
		log.Fatalln(err)
	}
}
