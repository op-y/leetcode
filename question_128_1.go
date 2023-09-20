package leetcode

func longestConsecutive(nums []int) int {
    if nums == nil {
        return 0
    }
    if len(nums) <= 1  {
        return len(nums)
    }

    threewayqsort(nums, 0, len(nums) - 1)
    maxLen := 1
    Len := 1
    for i:=1; i<len(nums); i++ {
        if nums[i] == nums[i - 1] {
            continue
        } else if nums[i] == nums[i - 1] + 1 {
            Len++
        } else {
            Len = 1
        }
        if Len > maxLen {
            maxLen = Len
        }
    }
    return maxLen
}


func threewayqsort(a []int, lo, hi int) {
    if hi <= lo {
        return
    }

    flag := a[lo]
    lt := lo
    i := lo + 1
    gt := hi
    for i <= gt {
        if a[i] < flag {
            tmp := a[lt]
            a[lt] = a[i]
            a[i] = tmp
            lt++
            i++
        } else if a[i] > flag {
            tmp := a[gt]
            a[gt] = a[i]
            a[i] = tmp
            gt--
        } else {
            i++
        }
    }
    threewayqsort(a, lo, lt - 1)
    threewayqsort(a, gt + 1, hi)
}
