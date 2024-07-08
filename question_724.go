package leetcode

func pivotIndex(nums []int) int {
    l := len(nums)
    // 只有一个元素边界情况
    if l == 1 {
        return 0
    }

    sum := 0
    for _, elem := range nums {
        sum += elem
    }

    leftSum := 0
    for i, elem := range nums {
        rightSum := sum - leftSum - elem
        if leftSum == rightSum {
            return i
        } else {
            leftSum += elem
        }
    }

    return -1
}
