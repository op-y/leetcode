package leetcode

func threeSumClosest(nums []int, target int) int {
    // 提示存在唯一答案所以不做边界判断了
    N := len(nums)
    // 先排序
    InsertionSort(nums)
    // 先将差值设置成一个很大的值
    distMin := 1000000
    ans := 0
    var i, j, k int
    i = 0
    for i < N {
        j = i + 1
        k = N - 1
        for j < k {
            dist, sum, small := Check(nums[i], nums[j], nums[k], target)
            if dist < distMin {
                distMin = dist
                ans = sum
            }
            if small {
                j++
                for j < k && nums[j] == nums[j - 1] {
                    j++
                }
            } else {
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

func Check(a, b, c, target int) (int, int, bool) {
    sum := a + b + c
    if sum > target {
        return sum - target, sum, false
    } else {
        return target - sum, sum, true
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
