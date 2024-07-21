package leetcode

func maximumSum(arr []int) int {
    dp0, dp1 := arr[0], 0
    result := arr[0]

    for i:=1; i<len(arr); i++ {
        dp0, dp1 = max(dp0, 0) + arr[i], max(dp1 + arr[i], dp0)
        result = max(result, max(dp0, dp1))
    }
    return result
}
