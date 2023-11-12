package leetcode

import (
	"strconv"
)

func largestNumber(nums []int) string {
	sort179(nums)
	s := ""
	for i := 0; i < len(nums); i++ {
		s += strconv.Itoa(nums[i])
	}
	for i := 0; i < len(s); i++ {
		if s[i] != '0' {
			return s
		}
	}
	return "0"
}

func sort179(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			// 小的往后移动
			if compare179(nums[j], nums[j+1]) == -1 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func compare179(a, b int) int {
	if a == b {
		return 0
	}
	if a == 0 {
		return -1
	}

	na, nb := a, b
	la := []int{}
	lb := []int{}
	for na != 0 {
		qa := na / 10
		ra := na % 10
		la = append(la, ra)
		na = qa
	}
	for nb != 0 {
		qb := nb / 10
		rb := nb % 10
		lb = append(lb, rb)
		nb = qb
	}
	i, j := len(la)-1, len(lb)-1
	for i >= 0 && j >= 0 {
		if la[i] > lb[j] {
			return 1
		}
		if la[i] < lb[j] {
			return -1
		}
		i--
		j--
	}
	ab := a
	for i := 0; i < len(lb); i++ {
		ab *= 10
	}
	ab += b

	ba := b
	for i := 0; i < len(la); i++ {
		ba *= 10
	}
	ba += a

	if ab > ba {
		return 1
	}
	if ab < ba {
		return -1
	}
	return 0
}

//func main() {
//	nums := []int{3, 30, 34, 5, 9}
//	//nums := []int{3432, 34323}
//	//nums := []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}
//	s := largestNumber(nums)
//	fmt.Println(s)
//}
