package leetcode

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			alive := getAlive(board, i, j)
			if alive < 2 && board[i][j] == 1 {
				board[i][j] = -2
			} else if alive > 3 && board[i][j] == 1 {
				board[i][j] = -2
			} else if alive == 3 && board[i][j] == 0 {
				board[i][j] = -1
			} else if (alive == 2 || alive == 3) && board[i][j] == 1 {
				board[i][j] = 1
			}
		}
	}
	unified(board)
}

func getAlive(board [][]int, x, y int) int {
	alive := 0
	m, n := len(board), len(board[0])
	if x-1 >= 0 && y-1 >= 0 {
		if board[x-1][y-1] == 1 || board[x-1][y-1] == -2 {
			alive++
		}
	}
	if x-1 >= 0 {
		if board[x-1][y] == 1 || board[x-1][y] == -2 {
			alive++
		}
	}
	if x-1 >= 0 && y+1 < n {
		if board[x-1][y+1] == 1 || board[x-1][y+1] == -2 {
			alive++
		}
	}
	if y-1 >= 0 {
		if board[x][y-1] == 1 || board[x][y-1] == -2 {
			alive++
		}
	}
	if y+1 < n {
		if board[x][y+1] == 1 || board[x][y+1] == -2 {
			alive++
		}
	}
	if x+1 < m && y-1 >= 0 {
		if board[x+1][y-1] == 1 || board[x+1][y-1] == -2 {
			alive++
		}
	}
	if x+1 < m {
		if board[x+1][y] == 1 || board[x+1][y] == -2 {
			alive++
		}
	}
	if x+1 < m && y+1 < n {
		if board[x+1][y+1] == 1 || board[x+1][y+1] == -2 {
			alive++
		}
	}
	return alive
}

func unified(board [][]int) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			} else if board[i][j] == -2 {
				board[i][j] = 0
			}
		}
	}
}
