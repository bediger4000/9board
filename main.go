package main

import (
	"9board/game"
	"fmt"
)

func main() {
	bd := &game.Board{}
	bd.Cell[0] = game.Max
	bd.Cell[1] = game.Min
	bd.Cell[4] = game.Max
	bd.Cell[8] = game.Max

	bd.Cell[10] = game.Max
	bd.Cell[14] = game.Min
	bd.Cell[17] = game.Min

	fmt.Printf("%s\n", bd)
	fmt.Printf("%s\n", bd.SubboardString(21))
}
