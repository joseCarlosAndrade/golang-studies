package main

import (
	"fmt"
	"raylib/pong"
)

func main() {
	fmt.Println("Hello")
	
	g := pong.InitEverything()
	
	g.Run()
	
}