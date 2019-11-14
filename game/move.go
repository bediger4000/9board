package game

// MakeMove takes a slot, values 0 - 88 inclusive,
// and marks the board-of-boards appropriately
func (b *Board) MakeMove(slot int, player Player) {
	b.Cell[slot] = player
}
