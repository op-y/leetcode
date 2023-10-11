package leetcode

func spiralOrder(matrix [][]int) []int {
	up := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	count := 0
	total := len(matrix) * len(matrix[0])
	direction := "R"
	result := []int{}
	i, j := 0, 0
	for i >= up && i <= bottom && j >= left && j <= right {
		switch direction {
		case "R":
			result = append(result, matrix[i][j])
			if j == right {
				direction = "D"
				i++
				up++
			} else {
				j++
			}
		case "D":
			result = append(result, matrix[i][j])
			if i == bottom {
				direction = "L"
				j--
				right--
			} else {
				i++
			}
		case "L":
			result = append(result, matrix[i][j])
			if j == left {
				direction = "U"
				i--
				bottom--
			} else {
				j--
			}
		case "U":
			result = append(result, matrix[i][j])
			if i == up {
				direction = "R"
				j++
				left++
			} else {
				i--
			}
		}
		count++
		if count == total {
			break
		}
	}
	return result
}

//func main() {
//	matrix := [][]int{}
//	matrix = append(matrix, []int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
//	r := spiralOrder(matrix)
//	fmt.Printf("%v\n", r)
//}
