package leetcode

func firstBadVersion(n int) int {
	low := 1
	high := n
	for {
		mid := (low + high) / 2
		if isBadVersion(mid) && !isBadVersion(mid-1) {
			return mid
		}
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

func isBadVersion(n int) bool {
	// TODO
	return true
}
