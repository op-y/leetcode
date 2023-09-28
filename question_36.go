package leetcode

func isValidSudoku(board [][]byte) bool {
    for i:=0; i<9; i++ {
        if !checkLine(board, i) {
            return false
        }
    }
    for i:=0; i<9; i++ {
        if !checkColumn(board, i) {
            return false
        }
    }
    for i:=0; i<9; i=i+3 {
        for j:=0; j<9; j=j+3 {
            if !checkBox(board, i, j) {
                return false
            }
        }
    }
    return true
}


func checkLine(board [][]byte, line int) bool {
    marks := newMarks()
    for i:=0; i<9; i++ {
        c := board[line][i]
        if c == '.' {
            continue
        }
        if marks[c] {
            return false
        } else {
            marks[c] = true
        }
    }
    return true
}

func checkColumn(board [][]byte, col int) bool {
    marks := newMarks()
    for i:=0; i<9; i++ {
        c := board[i][col]
        if c == '.' {
            continue
        }
        if marks[c] {
            return false
        } else {
            marks[c] = true
        }
    }
    return true
}

func checkBox(board [][]byte, line, col int) bool {
    marks := newMarks()
    for i:=line; i<line+3; i++ {
        for j:=col;j<col+3; j++ {
            c := board[i][j]
            if c == '.' {
                continue
            }
            if marks[c] {
                return false
            } else {
                marks[c] = true
            }
        }
    }
    return true
}

func newMarks() map[byte]bool {
    return map[byte]bool {
        '1': false,
        '2': false,
        '3': false,
        '4': false,
        '5': false,
        '6': false,
        '7': false,
        '8': false,
        '9': false,
    }
}
