package leetcode

func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])

	// 判断第一行中是否有0
	flag := false
	for j:=0; j<n; j++ {
		if matrix[0][j] == 0 {
			flag = true
			break
		}	
	}

	// 从最后一行往上处理
	for i:=m-1; i> 0; i-- {
		for j:=0; j<n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0 // 发现0 在第一行做列标记
				matrix[i][0] = 0 // 同时用本行第一列做行标记因为从左往右遍历
			}
		}
		if matrix[i][0] == 0 {
			for j:=0; j<n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据标记处理列
	for j:=0; j<n; j++ {
		if matrix[0][j] == 0 {
			for i:=0; i<m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据flag标记处理第一行
	if flag {
		for j:=0; j<n; j++ {
			matrix[0][j] = 0
		}
	}
}
