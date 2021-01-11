package leetcode

//func climbStairs(n int) int {
//    if n == 0 {
//        return 0
//    } else if n == 1 {
//        return 1
//    } else if n == 2 {
//        return 2
//    } else {
//        return climbStairs(n - 1) + climbStairs(n - 2)
//    }
//}

func climbStairs(n int) int {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
    if n == 2 {
        return 2
    }
    b2 := 1
    b1 := 2
    for i := 3; i<=n; i++ {
        tmp := b1
        b1 = b1 + b2
        b2 = tmp
    }
    return b1
}
