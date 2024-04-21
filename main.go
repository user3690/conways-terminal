package main

import (
	"conways-terminal/internal"
	"conways-terminal/pkg/conways"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	internal.HideCursor()
	internal.ClearScreen()

	keystrokeTicker := time.NewTicker(1 * time.Millisecond)
	gameTicker := time.NewTicker(500 * time.Millisecond)
	renderTicker := time.NewTicker(40 * time.Millisecond)

	game, err := conways.NewGameOfLife()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			<-renderTicker.C
			game.Render()
		}
	}()

	go func() {
		for {
			<-gameTicker.C
			game.Simulate()
		}
	}()

	go func() {
		for {
			<-keystrokeTicker.C
			done, err := game.Controls()
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
			internal.ShowCursor()
			internal.ClearScreen()

			fmt.Println("exit")
			os.Exit(0)
		default:
			// do nothing
		}
	}
}
