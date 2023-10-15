package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	if target < matrix[0][0] || target > matrix[m-1][n-1] {
		return false
	}

	row := -1
	low, high := 0, m - 1
	for low <= high {
		mid := (low + high) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] > target  {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}	
	row = high

	low, high = 0, n - 1
	for low <= high {
		mid := (low + high) / 2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] > target {
			high = mid - 1	
		} else {
			low = mid + 1
		}
	}
	return false
}
