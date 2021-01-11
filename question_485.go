package leetcode

func findMaxConsecutiveOnes(nums []int) int {
    if nums == nil {
       return 0
    }

    max := 0
    lastIsOne := false
    counter := 0
    for _, v := range nums {
        if !lastIsOne && v == 1{
            lastIsOne = true
            counter = 1
            continue
        }
        if !lastIsOne && v != 1 {
            continue
        }
        if lastIsOne && v == 1 {
            counter++
            continue
        }
        if lastIsOne && v != 1 {
            if counter > max {
                max = counter
            }
            counter = 0
            lastIsOne = false
            continue
        }
    }
    if counter > max {
        return counter
    } else {
        return max
    }
}
