package leetcode

import (
    "fmt"
)

func threeSum(nums []int) [][]int {
    if nums == nil {
        return nil
    }
    N := len(nums)
    if N < 3 {
        return nil
    }
    InsertionSort(nums)
    fmt.Printf("after sort: %v\n", nums)
    ans := make([][]int, 0)
    var i, j, k int
    i = 0
    for i < N {
        j = i + 1
        k = N - 1
        for j < k {
            fmt.Printf("i, j, k: %d %d %d\n", i, j, k)
            ok := Check(nums[i], nums[j], nums[k])
            if ok == -1 {
                j++
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
            } else if ok == 1 {
                k--
                for j < k && nums[k] == nums[k + 1] {
                    k--
                }
            } else {
                fmt.Printf("Hit: %d[%d] %d[%d] %d[%d]\n", i, nums[i], j, nums[j], k, nums[k])
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
                j++
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
                k--
                for j < k && nums[k] == nums[k + 1] {
                    k--
                }
            }
        }
        // 移动首个元素
        i++
        for i < N && nums[i] == nums[i - 1] {
            i++
        }
    }
    return ans
}

func Check(a, b, c int) int {
    sum := a + b + c
    if sum < 0 {
        return -1
    } else if sum > 0 {
        return 1
    } else {
        return 0
    }
}

func InsertionSort(a []int) {
    N := len(a)
    for i := 0; i < N; i++ {
        for j := i; j > 0 && a[j] < a[j - 1]; j-- {
            tmp := a[j]
            a[j] = a[j - 1]
            a[j - 1] = tmp
        }
    }
}
