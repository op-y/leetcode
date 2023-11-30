package leetcode

func productExceptSelf(nums []int) []int {
    l := len(nums)
    ans := make([]int, l)
    ans[0] = 1
    for i:=1; i<l; i++ {
        ans[i] = ans[i-1] * nums[i-1]
    }
    r := 1
    for i:=l-1; i>=0; i-- {
        ans[i] = ans[i] * r
        r *= nums[i]
    }
    return ans
}

