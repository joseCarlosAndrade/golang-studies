package game

import "fmt"

/* 
# Board type

Hols an array of 9 Player elements. */
type board struct {
	board [9]Player

}

/* Generates a new board. */
func newBoard() *board {
	var b = [9]Player{EMPTY, EMPTY,EMPTY,EMPTY,EMPTY,EMPTY,EMPTY,EMPTY,EMPTY,}
	return &board{b} 
}

/* String() string interface implementation for board type */
func (b board) String() string {
	return fmt.Sprintf(" %c | %c | %c \n" +
					   "___|___|___\n" +
					   " %c | %c | %c \n" +
					   "___|___|___\n" +
					   " %c | %c | %c \n" +
					   "   |   |   \n", b.board[0], b.board[1], b.board[2],
							 	  b.board[3], b.board[4], b.board[5],
								  b.board[6], b.board[7], b.board[8] )
}

/* Checks wether the position is valid. If so, play the round on the pos putting the character Player. */
func (b *board) Play(p Player, pos int) error {
	// error handling
	if pos <0 || pos > 8 {
		return IllegalPosition{}
	} else if b.board[pos] != ' ' {
		return &PositionAlreadyPlayed{pos}
	}

	// actually plays
	b.board[pos] = p
	return nil
}

/* Condition checking for victory */
func (b board) CheckGame() bool {
	for i:=0 ; i < 3; i++ {
		// checking lines
		if b.board[3*i] == EMPTY {
			
		} else if b.board[3*i] == b.board[3*i+1] &&  b.board[3*i] == b.board[3*i+2] {
			return true
		}
	
		// checking rows
		if b.board[i] == EMPTY {
			
		} else if b.board[i] == b.board[i+3] && b.board[i] == b.board[i+6] {
			return true
		}		

	}
	// checking diagonals
	
	if (b.board[0] != EMPTY)&&b.board[0]==b.board[4] && b.board[0]==b.board[8] {
		return true
	} else if (b.board[2] != EMPTY) &&b.board[2]==b.board[4] && b.board[2]==b.board[6] {
		return true
	}
	return false
}


/////////////// errors /////////////////

/* Error that is returned when the chosen position is already taken. */
type PositionAlreadyPlayed struct {
	pos int
}

func (p *PositionAlreadyPlayed) Error() string {
	return fmt.Sprintf("position %d already played.", p.pos)
}

/* Error that is returned when the position is not valid. Eg: negative or greater than 8. */
type IllegalPosition struct {}

func (i IllegalPosition) Error() string {
	return "illegal position."
}


// implement functions: 
//	- board.Newboard() board
// 	- board.Play(pos int, ) bool
//


