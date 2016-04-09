package engine

import (
	"math/rand"
)

type PlayerType int

const (
	None  PlayerType = iota
	Order            = iota
	Chaos            = iota
)

func (pl PlayerType) String() string {
	switch pl {
	case Order:
		return "Order"
	case Chaos:
		return "Chaos"
	}
	return "None"
}

type Color int

const (
	Empty Color = iota
	Black       = iota
	White       = iota
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return "Empty"
}

type Position uint8

func (p Position) X() Position {
	return p & 0x0f
}

func (p Position) Y() Position {
	return (p & 0xf0) >> 4
}

func (p Position) Valid() bool {
	return 0 < p.X() && p.X() <= 0x06 && 0 < p.Y() && p.Y() <= 0x60
}

type Board interface {
	Read(Position) Color
	Place(Position, Color)
	Open() map[Position]bool
	Winner() PlayerType
}

type board map[Position]Color

func NewBoard() Board {
	return make(board)
}

func (b board) Read(p Position) Color {
	c, ok := b[p]
	if !ok {
		c = Empty
	}
	return c
}

func (b board) Place(p Position, c Color) {
	b[p] = c
}

func (b board) Open() map[Position]bool {
	ret := make(map[Position]bool)
	var x, y uint8
	for x = 0x01; x <= 0x06; x++ {
		for y = 0x10; y <= 0x60; y++ {
			p := Position(y + x)
			if _, ok := b[p]; !ok {
				ret[p] = true
			}
		}
	}
	return ret
}

func (b board) Winner() PlayerType {
	neighbors := []uint8{0x01, 0x10, 0x11}
	var x, y uint8
	for x = 0x01; x <= 0x06; x++ {
		for y = 0x10; y <= 0x60; y++ {
			p := Position(y + x)
			if c := b.Read(p); c != Empty {
				for _, n := range neighbors {
					var inRow uint8 = 1
					for ; inRow < 5; inRow++ {
						if c != b.Read(Position(x+y+n*inRow)) {
							break
						}
					}
					if inRow == 5 {
						return Order
					}
				}
			}
		}
	}
	if len(b) == 6*6 {
		return Chaos
	}
	return None
}

type Player interface {
	SetPlayer(PlayerType)
	Move(Board) (Position, Color)
}

type Game interface {
	Board() Board
	MoveNum() int
	Move() (Position, Color)
}

type game struct {
	order   Player
	chaos   Player
	board   Board
	moveNum int
}

func NewGame(p1, p2 Player) Game {
	g := &game{board: NewBoard()}
	if rand.Intn(2) == 0 {
		g.order = p1
		g.chaos = p2
	} else {
		g.order = p2
		g.chaos = p1
	}
	g.order.SetPlayer(Order)
	g.chaos.SetPlayer(Chaos)
	return g
}

func (g *game) Board() Board {
	return g.board
}

func (g *game) MoveNum() int {
	return g.moveNum
}

func (g *game) Move() (Position, Color) {
	var p Player
	if g.MoveNum()%2 == 0 {
		p = g.order
	} else {
		p = g.chaos
	}
	pos, c := p.Move(g.board)
	g.board.Place(pos, c)
	g.moveNum++
	return pos, c
}
