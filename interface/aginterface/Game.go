package aginterface

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