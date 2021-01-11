package leetcode

func thirdMax(nums []int) int {
    status := 0
    first := 0
    second := 0
    third := 0

    for _, v := range nums {
        if status == 0 {
            first = v
            status = 1
            continue
        }

        if status == 1 {
            if v > first {
                second = first
                first = v
                status = 2
                continue
            } else if v == first {
                continue
            } else {
                second = v
                status = 2
                continue
            }
        }

        if status == 2 {
            if v > first {
                third = second
                second = first
                first = v
                status = 3
                continue
            } else if v == first {
                continue
            } else if v > second {
                third = second
                second = v
                status = 3
                continue
            } else if v == second {
                continue
            } else {
                third = v
                status = 3
                continue
            }
        }

        if status == 3 {
            if v > first {
                third = second
                second = first
                first = v
                continue
            } else if v == first {
                continue
            } else if v > second {
                third = second
                second = v
                continue
            } else if v == second {
                continue
            } else if v > third {
                third = v
                continue
            } else {
                continue
            }
        }
    }

    if status == 3 {
        return third
    } else {
        return first
    }
}
