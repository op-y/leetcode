package leetcode

//import (
//	"math"
//)
//
//func calculateMinimumHP(dungeon [][]int) int {
//	m, n := len(dungeon), len(dungeon[0])
//
//	ans := math.MinInt
//	var backoff func(int, int, int, int)
//	backoff = func(x, y, health, lose int) {
//		health += dungeon[x][y]
//		newlose := lose
//		if health < newlose {
//			newlose = health
//		}
//		// 剪枝 回退
//		if newlose < ans {
//			health -= dungeon[x][y]
//			return
//		}
//		// 已经走到尽头的到一个答案
//		if x == m-1 && y == n-1 {
//			ans = newlose
//			health -= dungeon[x][y]
//			return
//		}
//		if x+1 < m {
//			backoff(x+1, y, health, newlose)
//		}
//		if y+1 < n {
//			backoff(x, y+1, health, newlose)
//		}
//	}
//	backoff(0, 0, 0, 0)
//	if ans >= 0 {
//		return 1
//	}
//	return (-ans + 1)
//}
//
//func main() {
//	//dungeon := [][]int{
//	//	{-2, -3, 3},
//	//	{-5, -10, 1},
//	//	{10, 30, -5},
//	//}
//	dungeon := [][]int{
//		{1, 2, 3},
//	}
//	ans := calculateMinimumHP(dungeon)
//	fmt.Println(ans)
//}
