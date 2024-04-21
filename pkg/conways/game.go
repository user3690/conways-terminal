package conways

import (
	"conways-terminal/internal"
)

const (
	height = 30
	width  = 60
)

type PlayingField = [][]Cell

type GameOfLife struct {
	keyboard  *internal.KeyboardReader
	cursorPos internal.Position
	area      PlayingField
	simulate  bool
}

func NewGameOfLife() (GameOfLife, error) {
	var (
		newPlayingField  = make(PlayingField, height)
		initialCursorPos = internal.Position{}
	)

	var x, y int
	for y = 0; y < height; y++ {
		newPlayingField[y] = make([]Cell, width)
		for x = 0; x < width; x++ {
			newPlayingField[y][x] = NewCell(x, y)
			if initialCursorPos == newPlayingField[y][x].position {
				newCell := newPlayingField[y][x]
				newCell.hasCursor = true
				newPlayingField[y][x] = newCell
			}
		}
	}

	reader, err := internal.NewKeyboard()
	if err != nil {
		return GameOfLife{}, err
	}

	return GameOfLife{
		keyboard:  reader,
		cursorPos: initialCursorPos,
		area:      newPlayingField,
	}, nil
}

func (game *GameOfLife) CurrentPlayingField() PlayingField {
	return game.area
}

func (game *GameOfLife) countAliveNeighboursForEveryCell() {
	var (
		aliveNeighbours uint8
		x, y            int
	)

	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			aliveNeighbours = 0
			cell := game.area[y][x]

			// upper
			if cell.position.Y+1 < height {
				if game.area[cell.position.Y+1][cell.position.X].isAlive {
					aliveNeighbours++
				}
			}

			// upper right
			if cell.position.Y+1 < height && cell.position.X+1 < width {
				if game.area[cell.position.Y+1][cell.position.X+1].isAlive {
					aliveNeighbours++
				}
			}

			// right
			if cell.position.X+1 < width {
				if game.area[cell.position.Y][cell.position.X+1].isAlive {
					aliveNeighbours++
				}
			}

			// bottom right
			if cell.position.Y-1 >= 0 && cell.position.X+1 < width {
				if game.area[cell.position.Y-1][cell.position.X+1].isAlive {
					aliveNeighbours++
				}
			}

			// bottom
			if cell.position.Y-1 >= 0 {
				if game.area[cell.position.Y-1][cell.position.X].isAlive {
					aliveNeighbours++
				}
			}

			// bottom left
			if cell.position.Y-1 >= 0 && cell.position.X-1 >= 0 {
				if game.area[cell.position.Y-1][cell.position.X-1].isAlive {
					aliveNeighbours++
				}
			}

			// left
			if cell.position.X-1 >= 0 {
				if game.area[cell.position.Y][cell.position.X-1].isAlive {
					aliveNeighbours++
				}
			}

			// upper left
			if cell.position.Y+1 < height && cell.position.X-1 >= 0 {
				if game.area[cell.position.Y+1][cell.position.X-1].isAlive {
					aliveNeighbours++
				}
			}

			cell.aliveNeighbours(aliveNeighbours)
			game.area[y][x] = cell
		}
	}
}

func (game *GameOfLife) newLifeCycle() {
	var (
		x, y int
	)

	for y = 0; y < height; y++ {
		for x = 0; x < width; x++ {
			cell := game.area[y][x]
			cell.evalState()
			game.area[y][x] = cell
		}
	}
}

func (game *GameOfLife) Simulate() {
	if game.simulate {
		game.countAliveNeighboursForEveryCell()
		game.newLifeCycle()
	}
}
