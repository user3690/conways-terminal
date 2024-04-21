package conways

func (game *GameOfLife) Controls() (done bool, err error) {
	var key []byte

	key, err = game.keyboard.KeyPress()
	if err != nil {
		return true, err
	}

	switch string(key) {
	case "w":
		game.MoveCursorUp()
	case "a":
		game.MoveCursorLeft()
	case "s":
		game.MoveCursorDown()
	case "d":
		game.MoveCursorRight()
	case "f":
		game.ToggleCell()
	case "x":
		game.RunSimulation()
	case "\x03":
		err = game.keyboard.Restore()

		return true, err
	}

	return false, nil
}

func (game *GameOfLife) MoveCursorDown() {
	if game.cursorPos.Y+1 >= height || game.simulate {
		return
	}

	// remove cursor from current cell
	currentCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	currentCell.hasCursor = false
	game.area[game.cursorPos.Y][game.cursorPos.X] = currentCell

	// move cursor pos
	game.cursorPos.Y = game.cursorPos.Y + 1

	// set new cursor
	newCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	newCell.hasCursor = true
	game.area[game.cursorPos.Y][game.cursorPos.X] = newCell
}

func (game *GameOfLife) MoveCursorRight() {
	if game.cursorPos.X+1 >= width || game.simulate {
		return
	}

	// remove cursor from current cell
	currentCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	currentCell.hasCursor = false
	game.area[game.cursorPos.Y][game.cursorPos.X] = currentCell

	// move cursor pos
	game.cursorPos.X = game.cursorPos.X + 1

	// set new cursor
	newCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	newCell.hasCursor = true
	game.area[game.cursorPos.Y][game.cursorPos.X] = newCell
}

func (game *GameOfLife) MoveCursorUp() {
	if game.cursorPos.Y-1 < 0 || game.simulate {
		return
	}

	// remove cursor from current cell
	currentCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	currentCell.hasCursor = false
	game.area[game.cursorPos.Y][game.cursorPos.X] = currentCell

	// move cursor pos
	game.cursorPos.Y = game.cursorPos.Y - 1

	// set new cursor
	newCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	newCell.hasCursor = true
	game.area[game.cursorPos.Y][game.cursorPos.X] = newCell
}

func (game *GameOfLife) MoveCursorLeft() {
	if game.cursorPos.X-1 < 0 || game.simulate {
		return
	}

	// remove cursor from current cell
	currentCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	currentCell.hasCursor = false
	game.area[game.cursorPos.Y][game.cursorPos.X] = currentCell

	// move cursor pos
	game.cursorPos.X = game.cursorPos.X - 1

	// set new cursor
	newCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	newCell.hasCursor = true
	game.area[game.cursorPos.Y][game.cursorPos.X] = newCell
}

func (game *GameOfLife) ToggleCell() {
	if game.simulate {
		return
	}

	currentCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	currentCell.isAlive = !currentCell.isAlive
	game.area[game.cursorPos.Y][game.cursorPos.X] = currentCell
}

func (game *GameOfLife) RunSimulation() {
	if game.simulate {
		return
	}

	game.simulate = true

	cursorCell := game.area[game.cursorPos.Y][game.cursorPos.X]
	cursorCell.hasCursor = false
	game.area[game.cursorPos.Y][game.cursorPos.X] = cursorCell
}
