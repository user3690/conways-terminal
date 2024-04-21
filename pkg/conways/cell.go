package conways

import "conways-terminal/internal"

type Cell struct {
	neighbours         uint8
	hasCursor, isAlive bool
	position           internal.Position
}

const (
	dead   = "." // . dead
	alive  = "x" // x alive
	cursor = "$" // + cursor
)

func NewCell(x, y int) Cell {
	return Cell{
		position: internal.Position{
			X: x,
			Y: y,
		},
	}
}

func (cell *Cell) String() string {
	if cell.isAlive && !cell.hasCursor {
		return alive
	}

	if cell.hasCursor {
		return cursor
	}

	return dead
}

func (cell *Cell) aliveNeighbours(neighbours uint8) {
	cell.neighbours = neighbours
}

func (cell *Cell) evalState() {
	if cell.neighbours > 3 {
		cell.isAlive = false
	}

	if cell.neighbours == 3 {
		cell.isAlive = true
	}

	if cell.neighbours < 2 {
		cell.isAlive = false
	}
}
