package tictactoe

import (
	"fmt"

	"github.com/Marvin9/AI-GTU-2180703/utils"
)

const blocks = 3

// Start - start game
func (b *Board) Start() {
	b.init()
	for {
		b.display()
		var placeToMark int
		if b.turn == Check {
			fmt.Print("\nPlayer 1's turn, select number to mark O: ")
			fmt.Scanf("%d", &placeToMark)
		} else {
			fmt.Print("\nPlayer 2's turn, select number to mark X: ")
			_, placeToMark = simulateAIMove(*b)
			fmt.Print(placeToMark)
		}
		validMark := b.mark(placeToMark)
		if !validMark {
			continue
		}

		winner := b.didAnyoneWon()
		if winner != -1 {
			utils.Log(fmt.Sprintf("%v won.", markSymbol(winner)))
			b.display()
			return
		}
		if b.totalMarked == blocks*blocks {
			utils.Log("Draw")
			b.display()
			return
		}
	}
}

func (b *Board) didAnyoneWon() int {
	// Nobody could win until 5th turn
	if b.totalMarked < 5 {
		return -1
	}

	// check row wise
	for row := 0; row < blocks; row++ {
		init := b.board[row][0]
		if init == NoMark {
			continue
		}
		// utils.Log(fmt.Sprintf("init before: %v", init))
		for col := 1; col < blocks; col++ {
			// utils.Log(fmt.Sprintf("init&=%v", b.board[row][col]))
			init &= b.board[row][col]
		}
		// utils.Log(fmt.Sprintf("init after: %v", init))
		if init == b.board[row][0] {
			return init
		}
	}

	// check column wise
	for col := 0; col < blocks; col++ {
		init := b.board[0][col]
		if init == NoMark {
			continue
		}
		for row := 1; row < blocks; row++ {
			init &= b.board[row][col]
		}
		if init == b.board[0][col] {
			return init
		}
	}

	// check diagonals
	// left to right
	init := b.board[0][0]
	if init != NoMark {
		row, col := 1, 1
		for row < blocks {
			init &= b.board[row][col]
			row++
			col++
		}
		if init == b.board[0][0] {
			return init
		}
	}

	init = b.board[0][2]
	if init != NoMark {
		// right to left
		row, col := 1, 1
		for row < blocks {
			init &= b.board[row][col]
			row++
			col--
		}
		if init == b.board[0][2] {
			return init
		}
	}
	return -1
}
