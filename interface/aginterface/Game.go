package aginterface

import rl "github.com/gen2brain/raylib-go/raylib"

// import (
// )

// screen
const (
	ScreenWidth int32 = 800
	ScreenHeight int32 = 700
	ScreenFPS int32 = 30
	ScreenTitle string = "Jogo do Asilo"

	// board
	BoardX int32 = 80
	BoardY int32 = 50
	BoardWidth int32 = 200
	BoardHeight int32 = 200
	BoardSpacing int32 = 10
	BoardGap int32 = 5 // gap from the borders

)

type Game struct {
	Boards []Board
	State int
	CurrentPlayer int
}

func NewGame() *Game {
	// base := NewBoard(Limits{BoardX, BoardY, BoardX+BoardWidth, BoardY+BoardHeight})
	boards:= make([]Board, 9)

	space := BoardWidth + BoardSpacing
	boards[0] = NewBoard(Limits{BoardX, BoardY, BoardX+BoardWidth, BoardY+BoardHeight}) 
	boards[1] = NewBoard(Limits{BoardX +space, BoardY, BoardX+BoardWidth+space, BoardY+BoardHeight}) 
	boards[2] = NewBoard(Limits{BoardX +2*space, BoardY, BoardX+BoardWidth +2*space, BoardY+BoardHeight}) 

	boards[3] = NewBoard(Limits{BoardX, BoardY+space, BoardX+BoardWidth, BoardY+BoardHeight+space}) 
	boards[4] = NewBoard(Limits{BoardX +space, BoardY+space, BoardX+BoardWidth+space, BoardY+BoardHeight+space}) 
	boards[5] = NewBoard(Limits{BoardX +2*space, BoardY+space, BoardX+BoardWidth +2*space, BoardY+BoardHeight+space}) 

	boards[6] = NewBoard(Limits{BoardX, BoardY+2*space, BoardX+BoardWidth, BoardY+BoardHeight+2*space}) 
	boards[7] = NewBoard(Limits{BoardX +space, BoardY+2*space, BoardX+BoardWidth+space, BoardY+BoardHeight+2*space}) 
	boards[8] = NewBoard(Limits{BoardX +2*space, BoardY+2*space, BoardX+BoardWidth +2*space, BoardY+BoardHeight+2*space}) 

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