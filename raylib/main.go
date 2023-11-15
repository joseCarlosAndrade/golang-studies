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
	
}