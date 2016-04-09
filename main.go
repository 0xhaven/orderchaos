package main

import (
	"fmt"

	"github.com/jacobhaven/orderchaos/ai"
	"github.com/jacobhaven/orderchaos/engine"
)

func main() {
	game := engine.NewGame(ai.NewRandAI(), ai.NewRandAI())
	for game.Board().Winner() == engine.None {
		pos, c := game.Move()
		fmt.Printf("%d\t%s placed at (%d, %d)\n", game.MoveNum(), c, pos.X(), pos.Y())
	}
	fmt.Printf("%s Won!\n", game.Board().Winner())
}
