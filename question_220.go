package leetcode

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
    m := map[int]int{}
    for i, x := range nums {
        id := getID(x, valueDiff + 1)
        if _, ok := m[id]; ok {
            return true
        }
        if y, ok := m[id-1]; ok && abs(x-y) <= valueDiff {
            return true
        }
        if y, ok := m[id+1]; ok && abs(x-y) <= valueDiff {
            return true
        }
        m[id] = x
        if i >= indexDiff {
            delete(m, getID(nums[i-indexDiff], valueDiff + 1)) // 只关注一个窗口之内的数
        }
    }
    return false
}

func getID(x, w int) int {
    if x >= 0 {
        return x / w
    }
    return (x + 1) / w - 1
}

