package main

import (
	"9board/game"
	"fmt"
)

func main() {
	bd := &game.Board{}

	moveSlot := 44
	for {
		fmt.Printf("%s\n", bd)
		moveSlot = game.ReadMove(bd, moveSlot)
		fmt.Printf("Read %d\n", moveSlot)
		bd.MakeMove(moveSlot, game.Min)
	}
}
