package pong

import (

	rl "github.com/gen2brain/raylib-go/raylib"
)



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
		g.CheckBarBallCollision()
		g.DrawObjects()
	}
}

func (g *Game) Inputs() {
	// right bar
	if rl.IsKeyDown(rl.KeyUp) {
		g.Objects[2].(*Bar).BarDirection = Up
	
	} else if rl.IsKeyDown(rl.KeyDown) {	
		g.Objects[2].(*Bar).BarDirection = Down
	} else {
		g.Objects[2].(*Bar).BarDirection = Stopped

	}

	// left bar
	if rl.IsKeyDown(rl.KeyW) {
		g.Objects[1].(*Bar).BarDirection = Up
	
	} else if rl.IsKeyDown(rl.KeyS) {	
		g.Objects[1].(*Bar).BarDirection = Down
	} else {
		g.Objects[1].(*Bar).BarDirection = Stopped

	}
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

func (g *Game)CheckBarBallCollision() {
	ball := g.Objects[0].(*Ball)
	barleft := g.Objects[1].(*Bar)
	barright := g.Objects[2].(*Bar)
	if ball.Position.X-ball.Radius <= barleft.RectangleBar.X+barleft.RectangleBar.Width { // right x
		// check 
		if ball.Position.Y >= barleft.RectangleBar.Y &&
		 ball.Position.Y <= barleft.RectangleBar.Y+barleft.RectangleBar.Height {
			ball.Velocity.X *= -1
		}
		
	} else if ball.Position.X+ball.Radius >= barright.RectangleBar.X {

		// check
		if ball.Position.Y >= barright.RectangleBar.Y &&
		 ball.Position.Y <= barright.RectangleBar.Y+ barright.RectangleBar.Height {
			ball.Velocity.X *= -1
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