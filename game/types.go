package game

type Player int

const (
	Empty Player = 0
	Max   Player = 1
	Min   Player = -1
)

type Board struct {
	Cell [89]Player
}
