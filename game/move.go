package game

import (
	"fmt"
	"os"
)

func (b *Board) MakeMove(slot int, player Player) {
	fmt.Fprintf(os.Stderr, "enter Board.MakeMove\n")
	b.Cell[slot] = player
}
