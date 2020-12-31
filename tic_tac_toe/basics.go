package tictactoe

import (
	"fmt"

	"github.com/Marvin9/AI-GTU-2180703/utils"
)

const (
	// NoMark - not marked on board
	NoMark = (iota + 1) << 1
	// Check - Right tick on board
	Check = iota << 2
	// Cross - Cross tick on board
	Cross
)

// Board is game board
type Board struct {
	board       [3][3]int
	turn        int
	totalMarked int
	bot         int
	human       int
}

func (b *Board) init() {
	for i := 0; i < blocks; i++ {
		for j := 0; j < blocks; j++ {
			b.board[i][j] = NoMark
		}
	}
	b.turn = Check
	b.bot = Cross
	b.human = Check
	b.totalMarked = 0
}

func (b *Board) display() {
	fmt.Println()
	for i := 0; i < blocks; i++ {
		if i == 1 {
			dashes()
			fmt.Println()
		}
		for j := 0; j < blocks; j++ {
			if j == 1 {
				fmt.Print("|")
			}
			displayMark := markSymbol(b.board[i][j])
			if b.board[i][j] == NoMark {
				displayMark = fmt.Sprintf("%v", indexToNumber(i, j))
			}
			fmt.Printf("\t%v\t", displayMark)
			if j == 1 {
				fmt.Print("|")
			}
		}
		if i == 1 {
			fmt.Println()
			dashes()
		}
		fmt.Println()
	}
}

func (b *Board) isMarked(at int) bool {
	i, j := numberToIndex(at)
	if b.board[i][j] == NoMark {
		return false
	}
	return true
}

func (b *Board) mark(at int) bool {
	if at < 1 || at > 9 {
		utils.Log("Invalid position to mark.")
		return false
	}
	if b.isMarked(at) {
		utils.Log("Already marked position.")
		return false
	}
	i, j := numberToIndex(at)
	b.board[i][j] = b.turn
	b.changePlayer()
	b.totalMarked++
	return true
}

func (b *Board) changePlayer() {
	if b.turn == Check {
		b.turn = Cross
	} else {
		b.turn = Check
	}
}

func dashes() {
	for dec := 0; dec < 50; dec++ {
		fmt.Printf("-")
	}
}

func indexToNumber(row, col int) int {
	return (row * blocks) + col + 1
}

func numberToIndex(num int) (int, int) {
	num--
	return num / blocks, num % blocks
}

func markSymbol(tp int) string {
	if tp == Check {
		return "O"
	}
	return "X"
}
