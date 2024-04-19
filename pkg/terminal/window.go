package terminal

type Position struct {
	x int // width
	y int // height
}
type Frame [][]byte
type Window struct {
	keyboard  *KeyboardReader
	cursorPos Position
	area      Frame
}

const frameHeight = 10
const frameWidth = 20

func NewWindow() (*Window, error) {
	var (
		newFrame         = make(Frame, frameHeight)
		initialCursorPos = Position{0, 0}
	)

	var x, y int
	for y = 0; y < frameHeight; y++ {
		newFrame[y] = make([]byte, frameWidth)
		for x = 0; x < frameWidth; x++ {
			pos := Position{x: x, y: y}

			newFrame[y][x] = '.'
			if initialCursorPos == pos {
				newFrame[y][x] = 'x'
			}
		}
	}

	reader, err := NewKeyboard()
	if err != nil {
		return &Window{}, err
	}

	return &Window{
		keyboard:  reader,
		cursorPos: initialCursorPos,
		area:      newFrame,
	}, nil
}

func (win *Window) Area() Frame {
	return win.area
}

func (win *Window) Move() (done bool, err error) {
	var key []byte

	key, err = win.keyboard.keyPress()
	if err != nil {
		return true, err
	}

	switch string(key) {
	case "w":
		win.MoveCursorUp()
	case "a":
		win.MoveCursorLeft()
	case "s":
		win.MoveCursorDown()
	case "d":
		win.MoveCursorRight()
	case "\x03":
		err = win.keyboard.restore()

		return true, err
	}

	return false, nil
}

func (win *Window) MoveCursorDown() {
	if win.cursorPos.y+1 >= frameHeight {
		return
	}

	win.area[win.cursorPos.y][win.cursorPos.x] = '.'
	win.cursorPos.y = win.cursorPos.y + 1
	win.area[win.cursorPos.y][win.cursorPos.x] = 'x'
}

func (win *Window) MoveCursorRight() {
	if win.cursorPos.x+1 >= frameWidth {
		return
	}

	win.area[win.cursorPos.y][win.cursorPos.x] = '.'
	win.cursorPos.x = win.cursorPos.x + 1
	win.area[win.cursorPos.y][win.cursorPos.x] = 'x'
}

func (win *Window) MoveCursorUp() {
	if win.cursorPos.y-1 < 0 {
		return
	}

	win.area[win.cursorPos.y][win.cursorPos.x] = '.'
	win.cursorPos.y = win.cursorPos.y - 1
	win.area[win.cursorPos.y][win.cursorPos.x] = 'x'
}

func (win *Window) MoveCursorLeft() {
	if win.cursorPos.x-1 < 0 {
		return
	}

	win.area[win.cursorPos.y][win.cursorPos.x] = '.'
	win.cursorPos.x = win.cursorPos.x - 1
	win.area[win.cursorPos.y][win.cursorPos.x] = 'x'
}
