package conways

import (
	"fmt"
	"strings"
)

func (game *GameOfLife) Render() {
	var (
		lineToRender []string
	)

	// render new area
	for y, cells := range game.area {
		game.renderBuffer.MoveTo(0, y+1)

		lineToRender = make([]string, width)
		for x, cell := range cells {
			lineToRender[x] = cell.String()
		}

		game.renderBuffer.WriteTo(strings.Join(lineToRender, ""))
		// clear last bits
		game.renderBuffer.ClearToEnd()
	}

	game.renderBuffer.MoveTo(0, len(game.area)+2)
	game.renderBuffer.WriteTo("'w' - up")
	game.renderBuffer.MoveTo(0, len(game.area)+3)
	game.renderBuffer.WriteTo("'d' - right")
	game.renderBuffer.MoveTo(0, len(game.area)+4)
	game.renderBuffer.WriteTo("'s' - down")
	game.renderBuffer.MoveTo(0, len(game.area)+5)
	game.renderBuffer.WriteTo("'a' - left")
	game.renderBuffer.MoveTo(0, len(game.area)+6)
	game.renderBuffer.WriteTo("'f' - create or kill cell")
	game.renderBuffer.MoveTo(0, len(game.area)+7)
	game.renderBuffer.WriteTo("'x' - start simulation")

	// write from buffer to actual terminal
	fmt.Print(game.renderBuffer.Read())

	game.renderBuffer.Reset()
}
