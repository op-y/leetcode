package leetcode

type Node427 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node427
	TopRight    *Node427
	BottomLeft  *Node427
	BottomRight *Node427
}

func construct(grid [][]int) *Node427 {
	n := len(grid)
	return gen427(grid, 0, n, 0, n)
}

func gen427(grid [][]int, x1, x2, y1, y2 int) *Node427 {
	ok, val := check427(grid, x1, x2, y1, y2)
	if ok {
		return &Node427{
			Val:    val,
			IsLeaf: true,
		}
	}
	return &Node427{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     gen427(grid, x1, (x1+x2)/2, y1, (y1+y2)/2),
		TopRight:    gen427(grid, x1, (x1+x2)/2, (y1+y2)/2, y2),
		BottomLeft:  gen427(grid, (x1+x2)/2, x2, y1, (y1+y2)/2),
		BottomRight: gen427(grid, (x1+x2)/2, x2, (y1+y2)/2, y2),
	}

}

func check427(grid [][]int, x1, x2, y1, y2 int) (bool, bool) {
	count0, count1 := 0, 0
	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			if grid[i][j] == 0 {
				count0++
			} else {
				count1++
			}
			if count0 > 0 && count1 > 0 {
				return false, false
			}
		}
	}
	if count0 == 0 {
		return true, true
	} else {
		return true, false
	}
}
