package leetcode

func hammingDistance(x int, y int) int {
     or := x ^ y
     count := 1
     for or != 0 {
         count++
         or = or & (or-1)
     }
     return count
}
