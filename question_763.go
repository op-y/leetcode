package leetcode

// 官方题解 贪心算法
func partitionLabels(s string) []int {
    partitions := []int{}

    lastPos := [26]int{}
    for i, c := range s {
        lastPos[c-'a'] = i
    }
    
    start, end := 0, 0
    for i, c := range s {
        if lastPos[c-'a'] > end {
            end = lastPos[c-'a']
        }
        if i == end {
            partitions = append(partitions, end-start+1)
            start = end + 1
        }
    }
    return partitions
}
