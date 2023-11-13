package leetcode

func numIslands(grid [][]byte) int {
	islands := 0
	m, n := len(grid), len(grid[0])

	var backoff func(int, int)
	backoff = func(x, y int) {
		grid[x][y] = 'X' // 标记
		if x-1 >= 0 && grid[x-1][y] == '1' {
			backoff(x-1, y) // 上
		}
		if x+1 < m && grid[x+1][y] == '1' {
			backoff(x+1, y) // 下
		}
		if y-1 >= 0 && grid[x][y-1] == '1' {
			backoff(x, y-1) // 左
		}
		if y+1 < n && grid[x][y+1] == '1' {
			backoff(x, y+1) // 右
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				islands++
				backoff(i, j)
			}
		}
	}

	return islands
}
