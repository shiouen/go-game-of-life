package main

import "log"

func main() {
	var game = NewGame(16, 16)

	if err := game.Run(); err != nil {
		log.Fatalln(err)
	}
}
