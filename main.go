package main

import (
	"conways-terminal/pkg/terminal"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	terminal.HideCursor()
	terminal.ClearScreen()

	keystrokeTicker := time.NewTicker(1 * time.Millisecond)
	renderTicker := time.NewTicker(40 * time.Millisecond)

	renderPipeline := terminal.NewRenderPipeline()
	win, err := terminal.NewWindow()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			<-renderTicker.C
			renderPipeline.Render(win.Area())
		}
	}()

	go func() {
		for {
			<-keystrokeTicker.C
			done, err := win.Move()
			if err != nil {
				panic(err.Error())
			}

			if done {
				keystrokeTicker.Stop()

				break
			}
		}
	}()

	for {
		select {
		case <-sigChan:
			terminal.ShowCursor()
			terminal.ClearScreen()

			fmt.Println("exit")
			os.Exit(0)
		default:
			// do nothing
		}
	}
}
