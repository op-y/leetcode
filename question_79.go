package leetcode

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	flags := make([][]bool, m)
	for i := 0; i < m; i++ {
		row := make([]bool, n)
		flags[i] = row
	}

	var search func(int, int, int) bool
	search = func(x, y, idx int) bool {
		flags[x][y] = true
		if idx+1 == len(word) {
			return true
		}

		if x-1 >= 0 && !flags[x-1][y] && board[x-1][y] == word[idx+1] {
			if ok := search(x-1, y, idx+1); ok {
				return true
			}
		}
		if x+1 < m && !flags[x+1][y] && board[x+1][y] == word[idx+1] {
			if ok := search(x+1, y, idx+1); ok {
				return true
			}
		}
		if y-1 >= 0 && !flags[x][y-1] && board[x][y-1] == word[idx+1] {
			if ok := search(x, y-1, idx+1); ok {
				return true
			}
		}
		if y+1 < n && !flags[x][y+1] && board[x][y+1] == word[idx+1] {
			if ok := search(x, y+1, idx+1); ok {
				return true
			}
		}

		flags[x][y] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if ok := search(i, j, 0); ok {
					return true
				}
			}
		}
	}
	return false
}
