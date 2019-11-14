package main

import (
	"9board/game"
	"fmt"
)

func main() {
	bd := &game.Board{}

	moveSlot := -1

	for {
		fmt.Printf("%s\n", bd)
		moveSlot = game.ReadMove(bd, moveSlot)
		bd.MakeMove(moveSlot, game.Min)
		if bd.CheckWin(moveSlot) {
			fmt.Printf("Human wins\n")
			break
		}
		moveSlot = uct.FindMove(bd, 100000)
		bd.MakeMove(moveSlot, game.Max)
		if bd.CheckWin(moveSlot) {
			fmt.Printf("Computer wins\n")
			break
		}
	}
}
