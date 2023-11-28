package leetcode

func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}

	nums := make([]int, 9)
	for i := 1; i < 10; i++ {
		nums[i-1] = i
	}

	min, max := 0, 0
	for i := 0; i < k; i++ {
		min += nums[i]
		max += nums[8-i]
	}

	if n < min || n > max {
		return ans
	}

	nl := []int{}
	sum := 0
	count := 0
	var backoff func(int)
	backoff = func(pos int) {
		if count == k {
			if sum == n { // 一个满足条件的组合
				one := make([]int, k)
				copy(one, nl)
				ans = append(ans, one)
			}
			return
		}
		for i := pos; i < 9; i++ {
			nl = append(nl, nums[i])
			sum += nums[i]
			count++
			backoff(i + 1)
			count--
			sum -= nums[i]
			nl = nl[:len(nl)-1]
		}
		return
	}
	backoff(0)
	return ans
}

//func main() {
//	k, n := 3, 9
//	ans := combinationSum3(k, n)
//	fmt.Printf("%+v\n", ans)
//}
