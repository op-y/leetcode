package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max4s(a ...int) int {
    res := a[0]
    for _, v := range a[1:] {
        if v > res {
            res = v
        }
    }
    return res
}

func gcd(a, b int) int {
    for a != 0 {
        a, b = b%a, a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

