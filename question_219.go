package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
    if nums == nil {
        return false
    }
    length := len(nums)
    if length < 2 {
        return false
    }
    dict := make(map[int][]int)
    for idx, value := range nums {
        if l, ok := dict[value]; ok {
            for _, pos := range l {
                if idx - pos <= k {
                    return true
                }
            }
            l = append(l, idx)
            dict[value] = l
        } else {
            s := make([]int, 0)
            s = append(s, idx)
            dict[value] = s
        }
    }
    return false
}
