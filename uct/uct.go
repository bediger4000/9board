package uct

import "9board/game"

type node struct {
	move            int
	parent          *node
	wins            int
	visits          int
	untriedMoves    []int
	playerJustMoved int
}

type state struct {
	game.Board
	playerJustMoved int
}

func FindMove(bd *game.Board, iterMax int) int {
	return 0
}
