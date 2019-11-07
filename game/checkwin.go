package game

var winningTriplets [][3]int = [][3]int{
	[3]int{0, 1, 2}, [3]int{3, 4, 5}, [3]int{6, 7, 8},
	[3]int{0, 3, 6}, [3]int{1, 4, 7}, [3]int{2, 5, 8},
	[3]int{0, 4, 8}, [3]int{2, 4, 6},
}

func (b *Board) CheckWin(slot int) bool {
	subboard := slot - (slot % 10)
	for _, triplet := range winningTriplets {
		n := b.Cell[subboard+triplet[0]] + b.Cell[subboard+triplet[1]] + b.Cell[subboard+triplet[2]]
		if n == 3 || n == -3 {
			return true
		}
	}
	return false
}
