package leetcode

func reverse(x int32) int32 {
    if x == 0 {
        return 0
    }
    sign := int32(1)
    if x < 0 {
        sign = -1
        x = sign * x
    }
    result := int32(0)
    for {
       remainder := x % 10
       x = x / 10
       if x > 0 {
            result = (result + remainder) * 10
       } else {
            result = result + remainder
            break
       }
    }
    result *= sign
    if sign == 1 && result < 0 {
        return int32(0)
    }
    if sign == -1 && result > 0 {
        return int32(0)
    }
    return result
}
