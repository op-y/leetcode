package leetcode

func majorityElement(nums []int) int {
    v := nums[0]
    c := 0

    for _, n := range nums {
        if c == 0 {
            v = n
            c = 1
            continue
        }
        if v == n {
            c++
        } else {
            c--
        }
    }
    return v
}
