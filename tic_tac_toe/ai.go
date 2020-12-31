package tictactoe

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
