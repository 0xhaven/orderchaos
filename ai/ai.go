package ai

import (
	"fmt"
	"math/rand"

	"github.com/jacobhaven/orderchaos/engine"
)

type randomAI struct {
	playerType engine.PlayerType
}

func NewRandAI() engine.Player {
	return randomAI{}
}

func (r randomAI) SetPlayer(playerType engine.PlayerType) {
	fmt.Println("RandAI is playing as:", playerType)
	r.playerType = playerType
}

func (r randomAI) Move(board engine.Board) (p engine.Position, c engine.Color) {
	for p = range board.Open() {
		c = randColor()
		return
	}
	return
}

func randColor() engine.Color {
	return engine.Color(1 + rand.Intn(2))
}
