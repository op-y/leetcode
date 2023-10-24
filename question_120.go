package leetcode

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	sum := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, i+1)
		// 第一行
		if i == 0 {
			row[0] = triangle[0][0]
			sum[0] = row
			continue
		}
		// 其它行
		for j := 0; j <= i; j++ {
			if j == 0 {
				row[j] = sum[i-1][j] + triangle[i][j]
			} else if j == i {
				row[j] = sum[i-1][j-1] + triangle[i][j]
			} else {
				if sum[i-1][j-1] < sum[i-1][j] {
					row[j] = sum[i-1][j-1] + triangle[i][j]
				} else {
					row[j] = sum[i-1][j] + triangle[i][j]
				}
			}
		}
		sum[i] = row
	}
	min := sum[n-1][0]
	for _, v := range sum[n-1] {
		if v < min {
			min = v
		}
	}
	return min
}
