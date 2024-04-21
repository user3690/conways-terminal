package conways

import (
	"conways-terminal/internal"
	"fmt"
)

func (game *GameOfLife) Render() {
	// clear area
	for y := range game.area {
		internal.MoveTo(0, y+1)
		internal.ClearLine()
	}

	// render new area
	for y, x := range game.area {
		internal.MoveTo(0, y+1)
		for _, cell := range x {
			fmt.Print(cell.String())
		}
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
	fmt.Print("'f' - awake or kill cell")
	internal.MoveTo(0, len(game.area)+7)
	fmt.Print("'x' - start simulation")
}
