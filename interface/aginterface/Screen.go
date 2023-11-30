package aginterface

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)


func (g * Game)InitScreen() {
	rl.InitWindow(ScreenWidth, ScreenHeight, ScreenTitle)
	rl.SetTargetFPS(ScreenFPS)
}

func (g * Game)Run() {

	defer rl.WindowShouldClose()
	
	for !rl.WindowShouldClose() {
		g.input()
		// update
		// draw
		g.drawEverything()
	}
	
}

// returns true if x ,y is within give limits
func insideLimit(x, y int, limit Limits) bool {
	// fmt.Printf("checking: x: %v y: %v, limit: %v\n", x, y, limit)
	if x >= int(limit.xo) && x <= int(limit.xf) {
		if y >= int(limit.yo) && y <= int(limit.yf) {
			return true
		}
	}
	return false
}

// verifies if the mouse x and y position is inside a piece. If so, change its state to selected
func insideLimitPiece(x, y int, b* Board) {
	for i:=0; i < len(b.Content); i++ {
		piece := &b.Content[i]
		if insideLimit(x, y, piece.Box) {
			piece.Selected = true
			// fmt.Println("changing color of piece", i, piece.Selected)
		} else {
			piece.Selected = false
		}
	}
}

func (g* Game)input() {

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		fmt.Println(rl.GetMousePosition())
	}

	x,y := int(rl.GetMousePosition().X), int(rl.GetMousePosition().Y)

	for i:=0 ; i < len(g.Boards); i++ {

		b := &g.Boards[i]
		// check first if mouse is inside board b
		if insideLimit(x, y, b.BoardBox) {
			// if it is, then checks the board piece
			insideLimitPiece(x, y, b)
			break // doesnt need to check all boards if one is detected
		}
	}
}

func (g* Game)drawEverything() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	drawBoard(g.Boards[0])

	rl.EndDrawing()
}

func drawBoard(b Board) {
	// limits of board
	// rl.DrawRectangle(BoardX, BoardY, BoardWidth, BoardHeight, rl.Gray)

	// draw each rectangle for each board piece
	width := b.Content[0].Box.xf - b.Content[0].Box.xo
	height := b.Content[0].Box.yf - b.Content[0].Box.yo
	for i:=0; i < len(b.Content); i++ {
		b := &b.Content[i]
		var color rl.Color = rl.Green
		// fmt.Printf("piece: %v, selected: %v \n", i , b.Selected)
		if b.Selected {
			color.G += 25
			color.B += 90
			color.R += 90
			b.Selected = false

			// fmt.Println("color is selected")
		}
		rl.DrawRectangle(b.Box.xo+BoardGap, b.Box.yo+BoardGap, width-2*BoardGap, height-2*BoardGap, color)
	}

	// drawing lines (test so far)
	for  i  := 0; i < 2; i++ {
		x := BoardX + (int32(i)+1)*BoardSpacing
		rl.DrawLine(x-1, BoardY, x-1, BoardY+BoardHeight, rl.White)
		rl.DrawLine(x, BoardY, x, BoardY+BoardHeight, rl.White)
		rl.DrawLine(x+1, BoardY, x+1, BoardY+BoardHeight, rl.White)
	}

	for  i  := 0; i < 2; i++ {
		y := BoardY + (int32(i)+1)*BoardSpacing
		rl.DrawLine(BoardX, y-1, BoardX+BoardWidth, y-1, rl.White)
		rl.DrawLine(BoardX, y, BoardX+BoardWidth, y, rl.White)
		rl.DrawLine(BoardX, y+1, BoardX+BoardWidth, y+1, rl.White)
	}

}

/*

..X...X..
...X.X...
....X....
...X.X...
..X...X..

...OOO...
.OO...OO.
.O.....O.
.OO...OO.
...OOO...

*/