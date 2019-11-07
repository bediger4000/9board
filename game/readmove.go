package game

import (
	"fmt"
	"io"
	"os"
)

func ReadMove(bd *Board, lastSlot int) int {
	moveSubboard := (lastSlot % 10) * 10
	var cellNo int
READMOVE:
	for {
		fmt.Printf("%sYour move: ", bd.SubboardString(lastSlot))
		_, err := fmt.Scanf("%d\n", &cellNo)
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			fmt.Printf("Failed to read: %v\n", err)
			os.Exit(1)
		}
		switch {
		case cellNo < 0 || cellNo > 8:
			fmt.Printf("Choose a number between 0 and 8, try again\n")
		case bd.Cell[moveSubboard+cellNo] != Empty:
			fmt.Printf("Cell already chosen, try again\n")
		default:
			break READMOVE
		}
	}
	return moveSubboard + cellNo
}
