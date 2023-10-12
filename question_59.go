package leetcode


func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i:=0; i<n; i++ {
		row := make([]int, n)
		matrix[i] = row
	}

	total := n * n
	cur := 1
	direction := "R"

	up := 0	
	bottom := n - 1
	left := 0
	right := n - 1

	i, j := 0, 0
	for i >= up && i <= bottom && j >= left && j <= right {
		matrix[i][j] = cur
		switch direction {
		case "R":
			if j == right {
				direction = "D"
				i++
				up++
			} else {
				j++
			}
		case "D":
			if i == bottom {
				direction = "L"
				j--
				right--
			} else {
				i++
			}
		case "L":
			if j == left {
				direction = "U"
				i--
				bottom--
			} else {
				j--
			}
		case "U":
			if i == up {
				direction = "R"
				j++
				left++
			} else {
				i--
			}
		}

		cur++
		if cur > total {
			break
		}
	}

	return matrix
}


//func main() {
//	n := 10
//	result := generateMatrix(n)
//	for i:=0; i<n; i++ {
//		fmt.Printf("%v\n", result[i])
//	}
//}
