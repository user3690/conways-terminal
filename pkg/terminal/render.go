package terminal

import (
	"fmt"
)

type RenderPipeline struct{}

func NewRenderPipeline() RenderPipeline {
	return RenderPipeline{}
}

func (pipe RenderPipeline) Render(windowArea Frame) {
	for y, _ := range windowArea {
		MoveTo(0, y+1)
		ClearLine()
	}

	for y, x := range windowArea {
		MoveTo(0, y+1)
		fmt.Print(string(x))
	}
}
