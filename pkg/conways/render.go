package conways

import (
	"conways-terminal/internal"
	"fmt"
	"strings"
)

func (game *GameOfLife) Render() {
	var lineToRender []string

	// render new area
	for y, cells := range game.area {
		internal.MoveTo(0, y+1)

		lineToRender = make([]string, width)
		for x, cell := range cells {
			lineToRender[x] = cell.String()
		}

		fmt.Print(strings.Join(lineToRender, ""))
		// clear last bits
		internal.ClearToEnd()
	}

	internal.MoveTo(0, len(game.area)+2)
	fmt.Print("'w' - up")
	internal.MoveTo(0, len(game.area)+3)
	fmt.Print("'d' - right")
	internal.MoveTo(0, len(game.area)+4)
	fmt.Print("'s' - down")
	internal.MoveTo(0, len(game.area)+5)
	fmt.Print("'a' - left")
	internal.MoveTo(0, len(game.area)+6)
	fmt.Print("'f' - create or kill cell")
	internal.MoveTo(0, len(game.area)+7)
	fmt.Print("'x' - start simulation")
}
