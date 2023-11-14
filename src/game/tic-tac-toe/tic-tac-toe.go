/*
Made by José Carlos Andrade do Nascimento - Github: @joseCarlosAndrade
Tic tac toe on golang
*/

package main

import (
	"fmt"
	"game"
)

func main() {
	fmt.Println("Tic Tac Toe")

	tictactoe := game.NewGame()

	var state game.State
	var err error
	
	for {
		tictactoe.ShowBoard()

		var pos int
		fmt.Scanln(&pos)
		fmt.Println("Position chose: ", pos)

		for {
			state, err = tictactoe.Play(pos)
			
			if err== nil {
				break
			}
			fmt.Println("Error! - ", err)
			fmt.Scanln(&pos)
			fmt.Println("Position chose: ", pos)
		}

		if state != game.GOING {
			break
		}
	}
	tictactoe.ShowBoard()

	switch state {
	case game.TIE:
		fmt.Println("Tie!")
		

	case game.VICTORY1:
		fmt.Println("Player 1 wins! (X)")
	
	case game.VICTORY2:
		fmt.Println("Player 2 wins! (O)")
	}
}

