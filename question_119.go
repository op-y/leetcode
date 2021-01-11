package leetcode

func getRow(rowIndex int) []int {
    if rowIndex < 0 {
        return nil
    }
    result := make([]int, rowIndex + 1)
    for i:=0; i<=rowIndex; i++ {
        for j:=i; j>=0; j-- {
            // 首尾为1
            if j == 0 || j == i {
                result[j] = 1
            } else {
                result[j] = result[j] + result[j-1]
            }
        }
    }
    return result
}
