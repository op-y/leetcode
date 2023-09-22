func divide(dividend int, divisor int) int {
    if dividend == 0 {
        return 0
    }
    if dividend == math.MinInt32 {
        if divisor == 1 {
            return math.MinInt32
        }
        if divisor == -1 {
            return math.MaxInt32
        }
    }
    if divisor == math.MinInt32 {
        if dividend == math.MinInt32 {
            return 1
        }
        return 0
    }

    rev := false
    if dividend > 0 {
        dividend = -dividend
        rev = !rev
    }
    if divisor > 0 {
        divisor = -divisor
        rev = !rev
    }

    ans := 0
    left, right := 1, math.MaxInt32
    for left <= right {
        mid := left + (right-left)>>1
        if quickAdd(divisor, mid, dividend) {
            ans = mid
            if mid == math.MaxInt32 {
                break
            }
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    if rev {
        return -ans
    } else {
        return ans
    }
}

func quickAdd(y, z, x int) bool {
    for result, add := 0, y; z > 0; z >>= 1 {
        if z&1 > 0 {
            // result + add >= x
            if result < x-add {
                return false
            }
            result += add
        }
        if z != 1 {
            // add + add >= x
            if add < x-add {
                return false
            }
            add += add
        }
    }
    return true
}
