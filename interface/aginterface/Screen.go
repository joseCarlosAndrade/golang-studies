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
	if x >= int(limit.xo) + 2 && x <= int(limit.xf) -2 {
		if y >= int(limit.yo)+2 && y <= int(limit.yf) -2 {
			return true
		}
	}
	return false
}

// verifies if the mouse x and y position is inside a piece. If so, change its state to selected
func (g*Game)insideLimitPiece(x, y int, b* Board) {
	for i:=0; i < len(b.Content); i++ {
		piece := &b.Content[i]
		if insideLimit(x, y, piece.Box) {
			piece.Selected = true
			// fmt.Println("changing color of piece", i, piece.Selected)

			// implement click here
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				g.PlayRound(piece)
			}
						
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
			g.insideLimitPiece(x, y, b)
			break // doesnt need to check all boards if one is detected
		}
	}
}

func (g* Game)drawEverything() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// drawBoard(g.Boards[0])
	for i := 0; i < len(g.Boards); i++ {
		drawBoard(&g.Boards[i])
	}
	g.PutPlayerOnScreen(60, 470, 20)

	rl.EndDrawing()
}

func drawBoard(b* Board) {
	// limits of board
	// rl.DrawRectangle(BoardX, BoardY, BoardWidth, BoardHeight, rl.Gray)

	// draw each rectangle for each board piece
	width := b.Content[0].Box.xf - b.Content[0].Box.xo
	height := b.Content[0].Box.yf - b.Content[0].Box.yo
	// gap := BoardGap + 5
	for i:=0; i < len(b.Content); i++ {
		b := &b.Content[i]
		var color rl.Color = rl.Green
		// fmt.Printf("piece: %v, selected: %v \n", i , b.Selected)
		if b.Selected {
			color.G += 25
			color.B += 90
			color.R += 90
			// b.Selected = false

			// fmt.Println("color is selected")
		}
		rl.DrawRectangle(b.Box.xo+BoardGap, b.Box.yo+BoardGap, width-2*BoardGap, height-2*BoardGap, color)
		b.Selected = false
	}

	// drawing lines (test so far)
	for  i  := 0; i < 2; i++ {
		
		x := b.BoardBox.xo + (int32(i)+1)*BoardWidth/3
		rl.DrawLine(x-1, b.BoardBox.yo, x-1, b.BoardBox.yf, rl.White)
		rl.DrawLine(x, b.BoardBox.yo, x, b.BoardBox.yf, rl.White)
		rl.DrawLine(x+1, b.BoardBox.yo, x+1, b.BoardBox.yf, rl.White)
	}

	for  i  := 0; i < 2; i++ {
		y := b.BoardBox.yo + (int32(i)+1)*BoardHeight/3
		rl.DrawLine(b.BoardBox.xo, y-1, b.BoardBox.xf, y-1, rl.White)
		rl.DrawLine(b.BoardBox.xo, y, b.BoardBox.xf, y, rl.White)
		rl.DrawLine(b.BoardBox.xo, y+1, b.BoardBox.xf, y+1, rl.White)
	}

	// drawing shapes
	for _, piece := range b.Content {
		if piece.Shape == NAS { continue }
		drawShape(piece, piece.Shape)
		
		// if piece.Shape == X {
		// 	drawShape(piece, X)
		// 	// rl.DrawLine(piece.Box.xo+1+gap, piece.Box.yo+gap, piece.Box.xf-gap, piece.Box.yf - 1 - gap, rl.Red)
		// 	// rl.DrawLine(piece.Box.xo+gap, piece.Box.yo+gap, piece.Box.xf-gap, piece.Box.yf - gap, rl.Red)
		// 	// rl.DrawLine(piece.Box.xo+gap, piece.Box.yo+1+gap, piece.Box.xf-1-gap, piece.Box.yf - gap, rl.Red)

		// 	// rl.DrawLine(piece.Box.xf-gap, piece.Box.yo+gap, piece.Box.xo+1+gap, piece.Box.yf - 1 - gap, rl.Red)
		// 	// rl.DrawLine(piece.Box.xf-gap, piece.Box.yo+gap, piece.Box.xo+gap , piece.Box.yf - gap, rl.Red)
		// 	// rl.DrawLine(piece.Box.xf-1-gap, piece.Box.yo+1+gap, piece.Box.xo+gap, piece.Box.yf - gap, rl.Red)
		// 	// rl.DrawRectangle(piece.Box.xo+BoardGap, piece.Box.yo+BoardGap, width-2*BoardGap, height-2*BoardGap, rl.Red)
		// } else {
		// 	// rl.DrawRectangle(piece.Box.xo+BoardGap, piece.Box.yo+BoardGap, width-2*BoardGap, height-2*BoardGap, rl.Blue)
		// 	drawShape(piece, O)
		// }
	}
}


func drawShape(b BoardPiece, s Shape) {
	width := b.Box.xf - b.Box.xo
	height := b.Box.yf - b.Box.yo
	gap := BoardGap + 10

	if s == X {
		for i:=int32(0); i < 100; i ++ {
			xl := b.Box.xo + gap + i*(width - 2*gap)/100
			xr := b.Box.xf - gap - i*(width - 2*gap)/100
			y := b.Box.yo + gap + i*(height - 2*gap)/100

			rl.DrawCircle(xl, y, 3, rl.Red)
			rl.DrawCircle(xr, y, 3, rl.Red)
		}
	} else if s == O {
		radius := (width - 2*gap)/2 -5
		for i:=0 ; i < 6; i++ {
			rl.DrawCircleLines((b.Box.xf + b.Box.xo)/2, (b.Box.yf + b.Box.yo)/2, float32(radius)+float32(i), rl.Blue)
		}
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