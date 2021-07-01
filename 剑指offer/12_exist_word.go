package offer

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if find(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, r, c int, word string) bool {
	rows, cols := len(board), len(board[0])
	if len(word) == 0 {
		return true
	}
	if r < 0 || c < 0 || r >= rows || c >= cols {
		return false
	}
	if board[r][c] != word[0] {
		return false
	}
	board[r][c] = -board[r][c]
	result := find(board, r-1, c, word[1:]) || find(board, r, c+1, word[1:]) || find(board, r+1, c, word[1:]) || find(board, r, c-1, word[1:])
	board[r][c] = -board[r][c]
	return result
}
