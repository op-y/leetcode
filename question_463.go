package leetcode

func islandPerimeter(grid [][]int) int {
    if grid == nil {
        return 0
    }
    row := len(grid)
    column := len(grid[0])
    perimeter := 0
    for i:=0; i<row; i+=1 {
        for j:=0; j<column; j+=1 {
            if grid[i][j] == 1 {
                increment := 4
                if i-1 >= 0 && grid[i-1][j] == 1 {
                    increment--
                }
                if i+1 < row && grid[i+1][j] == 1 {
                    increment--
                }
                if j-1 >= 0 && grid[i][j-1] == 1 {
                    increment--
                }
                if j+1 < column && grid[i][j+1] == 1 {
                    increment--
                }
                perimeter += increment
            }
        }
    }
    return perimeter
}
