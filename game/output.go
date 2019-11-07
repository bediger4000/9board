package game

import (
	"fmt"
)

var printOrder []int = []int{
	0, 1, 2, 10, 11, 12, 20, 21, 22,
	3, 4, 5, 13, 14, 15, 23, 24, 25,
	6, 7, 8, 16, 17, 18, 26, 27, 28,
	30, 31, 32, 40, 41, 42, 50, 51, 52,
	33, 34, 35, 43, 44, 45, 53, 54, 55,
	36, 37, 38, 46, 47, 48, 56, 57, 58,
	60, 61, 62, 70, 71, 72, 80, 81, 82,
	63, 64, 65, 73, 74, 75, 83, 84, 85,
	66, 67, 68, 76, 77, 78, 86, 87, 88,
}

var extras map[int]string = map[int]string{
	2: " | ", 12: " | ", 22: "\n",
	5: " | ", 15: " | ", 25: "\n",
	8: " | ", 18: " | ", 28: "\n---------------\n",
	32: " | ", 42: " | ", 52: "\n",
	35: " | ", 45: " | ", 55: "\n",
	38: " | ", 48: " | ", 58: "\n---------------\n",
	62: " | ", 72: " | ", 82: "\n",
	65: " | ", 75: " | ", 85: "\n",
	68: " | ", 78: " | ", 88: "\n",
}

func (b *Board) String() string {
	var stringrep string
	for _, slot := range printOrder {
		stringrep += b.Mark(slot)
		if extra, ok := extras[slot]; ok {
			stringrep += extra
		}
	}
	return stringrep
}

func (b *Board) SubboardString(lastSlot int) string {
	var stringrep string
	subboard := 10 * (lastSlot % 10)
	for i := 0; i < 9; i++ {
		slot := subboard + i
		stringrep += b.SubMark(slot, i)
		if (i % 3) == 2 {
			stringrep += "\n"
		}
	}
	return stringrep
}

func (b *Board) SubMark(slot int, cell int) string {
	if b.Cell[slot] == Empty {
		return fmt.Sprintf("%d ", cell)
	}
	return b.Mark(slot) + " "
}

func (b *Board) Mark(slot int) string {
	switch b.Cell[slot] {
	case Max:
		return "X"
	case Min:
		return "O"
	case Empty:
		return "_"
	default:
		return "?"
	}
}
