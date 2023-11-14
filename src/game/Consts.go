package game

type Player rune 

/* One of the ways of representing enums in Go */
const (
	PLAYER1 Player  = 'X'
	PLAYER2 Player  = 'O'
	EMPTY   Player  = ' '
)

type State int

/* State of the game  */
const (
	GOING    State = 0 // game going
	TIE      State = 1 // game tied (deu velha)
	VICTORY1 State = 2 // game finished with a victory
	VICTORY2 State = 3 // game finished with a victory
)