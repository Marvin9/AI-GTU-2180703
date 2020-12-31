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

func markSymbol(tp int) string {
	if tp == Check {
		return "O"
	}
	return "X"
}

const blocks = 3

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
			var symbolicWinner string
			if winner == Check {
				symbolicWinner = "O"
			} else {
				symbolicWinner = "X"
			}
			utils.Log(fmt.Sprintf("%v won.", symbolicWinner))
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

func simulateAIMove(b Board) (int, int) {
	maxPoints, move := 0, -1
	for i := 0; i < blocks; i++ {
		for j := 0; j < blocks; j++ {
			if b.board[i][j] == NoMark {
				intoNum := indexToNumber(i, j)
				b.board[i][j] = b.turn
				b.totalMarked++
				winner := b.didAnyoneWon()
				if winner != -1 {
					if winner == b.bot {
						return 10, intoNum
					}
					return -10, intoNum

				} else if b.totalMarked == blocks*blocks {
					return 0, intoNum
				}
				b.changePlayer()
				nextPts, _ := simulateAIMove(b)
				b.changePlayer()
				if b.turn == b.human {
					if move == -1 || nextPts < maxPoints {
						// utils.Log(fmt.Sprintf("Turn: %v, suggested move %v", b.turn, intoNum))
						move = intoNum
						maxPoints = nextPts
					}
				} else {
					if move == -1 || nextPts > maxPoints {
						// utils.Log(fmt.Sprintf("Turn: %v, suggested move %v", b.turn, intoNum))
						move = intoNum
						maxPoints = nextPts
					}
				}
				b.board[i][j] = NoMark
				b.totalMarked--
			}
		}
	}
	return maxPoints, move
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
			col++
		}
		if init == b.board[0][2] {
			return init
		}
	}
	return -1
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
