package main

import (
	"fmt"
	"raylib/pong"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	fmt.Println("Hello")
	
	g := pong.InitEverything()
	
	g.Run()
	// rl.InitWindow(600, 400, "aaaaaaaaaaaaaaaaaaaaaaaa")
	// defer rl.CloseWindow()

	// rl.SetTargetFPS(60)

	// for !rl.WindowShouldClose() {
	// 	rl.BeginDrawing()
	// 	rl.ClearBackground(rl.Green)
	// 	rl.DrawCircle(0, 0, 30, rl.Red)

	// 	rl.EndDrawing()
	// }
}