package leetcode

func jump(nums []int) int {
    l := len(nums)
    
    steps := make([]int, l)
    steps[0] = 0
    for i:=1; i<l; i++ {
        steps[i] = -1
    }

    for i:=0; i<l; i++ {
        v := nums[i]
        for j:=1; j<=v; j++ {
            pos := i + j
            // 跳出范围了
            if pos >= l {
                break
            }
            if steps[pos] == -1 {
                steps[pos] = steps[i] + 1
                continue
            }
            if steps[pos] > steps[i] + 1 {
                steps[pos] = steps[i] + 1
                continue
            }
        }
    }

    return steps[l-1]
}
