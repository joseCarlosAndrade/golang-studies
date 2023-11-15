package pong

import rl "github.com/gen2brain/raylib-go/raylib"



func InitEverything() *Game {
	rl.InitWindow(ScreenWidth, ScreenHeight, ScreenTitle)

	g:= NewGame()
	g.Init()

	return g
}

func (g *Game)Run() {

	defer rl.CloseWindow() // pushing this function to stack to close window after
	
	rl.SetTargetFPS(int32(GameFPS))

	for !rl.WindowShouldClose() {
		g.Inputs()
		g.UpdateObjects()
		g.DrawObjects()
	}
}

func (g *Game) Inputs() {
	
}

func (g *Game) UpdateObjects() {
	for _, obj := range g.Objects {
		collision := obj.UpdatePosition()

		switch v:= obj.(type) {
		case *Ball:
			switch collision {
			case HeightCollision:
				v.Velocity.Y *= -1 // invert ball y velocity when it hits the floor or the top
			}
		}
	}
}


func (g Game) DrawObjects() {
	
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	for _, obj := range g.Objects {
		switch v := obj.(type) {
			case *Ball:
				
				rl.DrawCircle(int32(v.Position.X), int32(v.Position.Y), v.Radius, v.Color)	
				

			case *Bar:
			
				rl.DrawRectangle(v.RectangleBar.ToInt32().X, v.RectangleBar.ToInt32().Y, v.RectangleBar.ToInt32().Width, v.RectangleBar.ToInt32().Height, v.BarColor)
			
		}
	}

	rl.EndDrawing()
}