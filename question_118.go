package leetcode

func generate(numRows int) [][]int {
    result := [][]int{}
    if numRows <= 0 {
        return result
    }
    for i:=1; i<=numRows; i++ {
        layer := make([]int, i)
        for j:=0; j<i; j++ {
            // 首尾为1
            if j == 0 || j == i-1 {
                layer[j] = 1
            } else {
                layer[j] = result[i-2][j-1] + result[i-2][j]
            }
        }
        result = append(result, layer)
    }
    return result
}
