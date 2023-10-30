package leetcode

func solve130(board [][]byte) {
	m, n := len(board), len(board[0])
	// 第一行
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			search130(0, i, m, n, board)
		}
	}
	// 第一列
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			search130(i, 0, m, n, board)
		}
	}
	// 最后一行
	for i := 0; i < n; i++ {
		if board[m-1][i] == 'O' {
			search130(m-1, i, m, n, board)
		}
	}
	// 最后一列
	for i := 0; i < m; i++ {
		if board[i][n-1] == 'O' {
			search130(i, n-1, m, n, board)
		}
	}

	// 处理标记
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
				continue
			}
			if board[i][j] == '*' {
				board[i][j] = 'O'
				continue
			}
		}
	}
}

func search130(x, y, m, n int, board [][]byte) {
	// 与边界的O相连的O全都标记*
	if board[x][y] == 'O' {
		board[x][y] = '*'
		if (x - 1) >= 0 {
			search130(x-1, y, m, n, board)
		}
		if (x + 1) < m {
			search130(x+1, y, m, n, board)
		}
		if (y - 1) >= 0 {
			search130(x, y-1, m, n, board)
		}
		if (y + 1) < n {
			search130(x, y+1, m, n, board)
		}
	}
	return
}
