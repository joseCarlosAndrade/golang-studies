/*
# Pong Game in Golang

Made by Jose Carlos Andrade do Nascimento (github: @joseCarlosAndrade). Pong implementation using raylib package.
*/
package pong

import (
	"math"
	"math/rand"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)


/* 
Initializes everything from window to logic game. */
func InitEverything() *Game {
	rl.InitWindow(ScreenWidth, ScreenHeight, ScreenTitle)

	g:= NewGame()
	g.Init()

	return g
}

/* 
Resets the ball Position so it spawns randomly every time. */
func (b *Ball) Reset() {
	b.Position.X = float32(ScreenWidth)/2
	b.Position.Y = float32(ScreenHeight)/2

	vy := rand.Intn(6)
	vx := math.Sqrt(math.Pow(float64(BallVelocity), 2) + math.Pow(float64(vy), 2))
	if rand.Intn(2) == 1 {vx *= -1}
	if rand.Intn(2) == 1 {vy *= -1}
	b.Velocity.X = float32(vx)
	b.Velocity.Y = float32(vy)
}

/* 
Main loop that runs the game and handles input, object update, collision detection and drawing. */
func (g *Game)Run() {

	defer rl.CloseWindow() // pushing this function to stack to close window after
	
	rl.SetTargetFPS(int32(GameFPS))

	for !rl.WindowShouldClose() && !g.ShouldClose {
		g.Inputs()
		g.UpdateObjects()
		g.CheckBarBallCollision()
		g.DrawObjects()
	}
}

/* 
Handles input from terminal */
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

/* 
Updates all objects on the game by the GameObject interface that implements UpdatePosition() method. Each update also
handles screen border collision, returning a Collision type. */
func (g *Game) UpdateObjects() {
	for _, obj := range g.Objects {
		collision := obj.UpdatePosition()

		switch v:= obj.(type) {
		case *Ball:
			switch collision {
			case HeightCollision:
				v.Velocity.Y *= -1 // invert ball y velocity when it hits the floor or the top
			case RightCollision:
				// g.ShouldClose = true
				g.LeftScore += 1
				
				v.Reset()
			case LeftCollision:
				// g.ShouldClose = true
				g.RightScore += 1
				v.Reset()
			}
		}
	}
}

/* 
Checks the collision between the ball and the bar. This collision is checked separately from the screen border
collision because this one is a little more complex and it's better organized this way. */
func (g *Game)CheckBarBallCollision() {
	ball := g.Objects[0].(*Ball)
	barleft := g.Objects[1].(*Bar)
	barright := g.Objects[2].(*Bar)
	if ball.Position.X-ball.Radius <= barleft.RectangleBar.X+barleft.RectangleBar.Width { // left bar
		// check 
		if ball.Position.Y >= barleft.RectangleBar.Y &&
		 ball.Position.Y <= barleft.RectangleBar.Y+barleft.RectangleBar.Height {
			// ball.Velocity.X *= -1
			if d :=barleft.BarDirection; d == Stopped {
				ball.Velocity.X *= -1
			} else if d == Up {
				if ball.Velocity.Y >= BallVelocity*BallFactor { // max velocity
					ball.Velocity.X *= -1
					return
				} else {
					
					vy := ball.Velocity.Y
					vy += SpeedIncrement
					nvx := math.Sqrt(math.Pow(float64(BallVelocity), 2) + math.Pow(float64(vy), 2))

					ball.Velocity.X = float32(nvx)
					ball.Velocity.Y = vy

				}
			} else if d == Down {
				if ball.Velocity.Y <= -1*BallVelocity*BallFactor { // max velocity
					ball.Velocity.X *= -1
					return
				} else {
					vy := ball.Velocity.Y
					vy -= SpeedIncrement
					nvx := math.Sqrt(math.Pow(float64(BallVelocity), 2) + math.Pow(float64(vy), 2))

					ball.Velocity.X = float32(nvx)
					ball.Velocity.Y = vy
					
				}
			}


		}
		
	} else if ball.Position.X+ball.Radius >= barright.RectangleBar.X { // right bar

		// check
		if ball.Position.Y >= barright.RectangleBar.Y &&
		 ball.Position.Y <= barright.RectangleBar.Y+ barright.RectangleBar.Height {

			if d:=barright.BarDirection; d == Stopped {
				ball.Velocity.X *= -1
			} else if d == Up {
				if ball.Velocity.Y >= BallVelocity*BallFactor { // max velocity
					ball.Velocity.X *= -1
					return
				} else {
					vy := ball.Velocity.Y
					vy += SpeedIncrement
					nvx := math.Sqrt(math.Pow(float64(BallVelocity), 2) + math.Pow(float64(vy), 2))

					ball.Velocity.X = float32(-nvx)
					ball.Velocity.Y = vy

				}

			} else if d == Down {
				if ball.Velocity.Y <= -1*BallVelocity*BallFactor { // max velocity
					ball.Velocity.X *= -1
					return
				} else {
					vy := ball.Velocity.Y
					vy -= SpeedIncrement
					nvx := math.Sqrt(math.Pow(float64(BallVelocity), 2) + math.Pow(float64(vy), 2))

					ball.Velocity.X = float32(-nvx)
					ball.Velocity.Y = vy
					
				}
			}
		}
	}
}

/* 
Draws everything. It initializes rl.BeginDrawinG() and ends it here, so every draw method must be implemented here. */
func (g Game) DrawObjects() {
	
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black) // black background

	// scores
	rl.DrawText(fmt.Sprint(g.LeftScore), 25, ScreenHeight-30, 30, rl.White)
	rl.DrawText(fmt.Sprint(g.RightScore), ScreenWidth-40, ScreenHeight-30, 30, rl.White)

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