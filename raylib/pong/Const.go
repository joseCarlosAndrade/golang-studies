package pong

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collision int

type Direction int

const (
	// screen
	ScreenWidth int32 = 850
	ScreenHeight int32 = 600
	ScreenTitle string = "Pong!"
	GameFPS int = 30

	// ball
	BallRadius float32 = 10
	BallVelocity float32 = 20

	// enums for collision types
	NoCollision Collision = 0x00
	LeftCollision Collision = 0x01
	RightCollision Collision = 0x02
	HeightCollision Collision = 0x03

	// bar 
	BarVelocity float32 = 3
	BarWidth float32 = 10
	Barheight float32 = 120
	
	// direction
	Up Direction = 0x00
	Down Direction = 0x01
	Stopped Direction = 0x03
)

type GameObject interface {
	UpdatePosition() Collision
}

type Ball struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Radius float32 

	Color rl.Color
}

// implementing GameObject interface on game objects

// update ball position based on collision
func (b *Ball) UpdatePosition() Collision {
	vx, vy := b.Velocity.X, b.Velocity.Y

	b.Position = rl.NewVector2(b.Position.X + vx, b.Position.Y + vy)

	x, y := b.Position.X, b.Position.Y

	// horizontal collision TODO: BAR COLLISION
	if x-BallRadius <= 0 { return LeftCollision}
	if x+BallRadius >= float32(ScreenWidth) { return RightCollision}

	// vertical collision 
	if y-BallRadius <= 0 || y+BallRadius >= float32(ScreenHeight) {return HeightCollision}

	return NoCollision // no scenario collision
}

type Bar struct {
	RectangleBar rl.Rectangle
	// MaxYVelocity float64

	BarDirection Direction
	BarColor rl.Color
}

func (b *Bar) UpdatePosition() Collision {
	// if going on a collision direction, returns
	if b.BarDirection == Up {
		if b.RectangleBar.Y <= 0  { return HeightCollision}

		b.RectangleBar.Y -= BarVelocity

	} else if b.BarDirection ==  Down{
		if  b.RectangleBar.Y + b.RectangleBar.Height >= float32(ScreenHeight) { return HeightCollision}

		b.RectangleBar.Y += BarVelocity
	}
	
	// height collision
	

	return NoCollision
}	

type Game struct {
	Ball Ball
	Bar1 * Bar
	Bar2 *Bar

	Objects [3]GameObject

	GameOver bool // game stops

	ShouldClose bool // window should close

}

func NewGame() (g *Game) {
	g = &Game{}
	g.Init()

	return
} 

func (g *Game) Init() {
	// creating the main game object
	var ball GameObject  = 
			&Ball{rl.NewVector2(float32(ScreenWidth)/2, float32(ScreenHeight)/2), 
			rl.NewVector2(0, BallVelocity), 
			BallRadius, 
			rl.NewColor(255, 255, 255, 255)}
	var barleft GameObject = 
			&Bar{rl.NewRectangle(10, float32(ScreenHeight)/2 - Barheight/2, BarWidth, Barheight), 
			Stopped, 
			rl.NewColor(255,255,255,255)}
	var barright GameObject = 
			&Bar{rl.NewRectangle(float32(ScreenWidth)-10-BarWidth, float32(ScreenHeight)/2 - Barheight/2, BarWidth, Barheight), 
			Stopped, 
			rl.NewColor(255,255,255,255)}

	g.Objects[0] = ball
	g.Objects[1] = barleft
	g.Objects[2] = barright
}

