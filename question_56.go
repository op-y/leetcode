package leetcode

func merge56(intervals [][]int) [][]int {
	l := len(intervals)
	if l == 1 {
		return intervals
	}
	sort56(intervals)

	merged := [][]int{intervals[0]}
	for i:=1; i<l; i++ {
		if intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		} else {
			if intervals[i][1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
		}
	}

	return merged
}

func sort56(intervals [][]int) {
	for i := 1; i < len(intervals); i++ {
		for j := i; j > 0 && intervals[j][0] < intervals[j-1][0]; j-- {
			intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
		}
	}
}

//func main() {
//	intervals := [][]int{
//		{1,4},
//		{4,5},
//	}
//	result := merge(intervals)
//	fmt.Printf("%v\n", result)
//}

