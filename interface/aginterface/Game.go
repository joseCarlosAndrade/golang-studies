package aginterface

import rl "github.com/gen2brain/raylib-go/raylib"

// import (
// )

// screen
const (
	ScreenWidth int32 = 500
	ScreenHeight int32 = 500
	ScreenFPS int32 = 30
	ScreenTitle string = "Jogo do Asilo"

	// board
	BoardX int32 = 50
	BoardY int32 = 50
	BoardWidth int32 = 400
	BoardHeight int32 = 400
	BoardSpacing int32 = 400/3
	BoardGap int32 = 5 // gap from the borders

)

type Game struct {
	Boards []Board
	State int
	CurrentPlayer int
}

func NewGame() *Game {
	base := NewBoard(Limits{BoardX, BoardY, BoardX+BoardWidth, BoardY+BoardHeight})
	boards:= make([]Board, 9)

	boards[0] = base // TESTS
	
	game := &Game{
		Boards: boards,
		State: 0,
		CurrentPlayer: 1, // start as X
	}
	return game
}

func (g * Game)PlayRound(b * BoardPiece) bool {
	if b.Shape != NAS { return false} // if not empty, doesnt play

	if g.CurrentPlayer ==1 {
		b.Shape = X
	} else {
		b.Shape = O
	}
	g.CurrentPlayer *= -1 // invert current player

	return true
}

func (g Game)PutPlayerOnScreen(x, y, fontSize int32) {
	var text string
	if g.CurrentPlayer == 1 {
		text = "Player 1 (X)!"
		
	} else {
		text = "Player 2 (O)!"
	}
	rl.DrawText(text, x, y, fontSize, rl.White)
}