package leetcode

func constructRectangle(area int) []int {
    L := area
    W := 1
    result := []int{L, W}
    for L>W {
       L--
       W = area / L
       if area % L == 0 {
           result[0] = L
           result[1] = W
       }
    }
    return result
}
