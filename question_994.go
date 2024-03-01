package leetcode

import (
	"fmt"
)

func orangesRotting(grid [][]int) int {
	fresh := 0
	rotten := map[string][]int{}

	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			}
			if grid[i][j] == 2 {
				k := fmt.Sprintf("%d+%d", i, j)
				rotten[k] = []int{i, j}
			}
		}
	}

	if len(rotten) > 0 && fresh == 0 {
		return 0
	}

	if len(rotten) == 0 && fresh > 0 {
		return -1
	}

	if len(rotten) == 0 && fresh == 0 {
		return 0
	}

	loop := 0
	for len(rotten) != 0 {
		next := map[string][]int{}
		for _, v := range rotten {
			x, y := v[0], v[1]

			if x > 0 && grid[x-1][y] == 1 {
				k := fmt.Sprintf("%d+%d", x-1, y)
				next[k] = []int{x - 1, y}
				grid[x-1][y] = 2
				fresh--
			}
			if x < m-1 && grid[x+1][y] == 1 {
				k := fmt.Sprintf("%d+%d", x+1, y)
				next[k] = []int{x + 1, y}
				grid[x+1][y] = 2
				fresh--
			}
			if y > 0 && grid[x][y-1] == 1 {
				k := fmt.Sprintf("%d+%d", x, y-1)
				next[k] = []int{x, y - 1}
				grid[x][y-1] = 2
				fresh--
			}
			if y < n-1 && grid[x][y+1] == 1 {
				k := fmt.Sprintf("%d+%d", x, y+1)
				next[k] = []int{x, y + 1}
				grid[x][y+1] = 2
				fresh--
			}
		}
		if len(next) > 0 {
			loop++
		}
		rotten = next
	}

	if fresh != 0 {
		return -1
	} else {
		return loop
	}
}
