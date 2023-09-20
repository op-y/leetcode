package leetcode

import (
    "fmt"
)

func permute(nums []int) [][]int {
    if nums == nil || len(nums) == 0 {
        return nil
    }
    if len(nums) == 1 {
       return [][]int{[]int{nums[0]}}
    }

    //fmt.Printf("nums: %v\n", nums)
    //sn := sort(nums)
    //fmt.Printf("sorted nums: %v\n", sn)

    ans := p(nums, 0)
    return dedup(ans)
}

func p(nums []int, pos int) [][]int  {
    //fmt.Printf("func p nums %v pos %d\n", nums, pos)
    ans := make([][]int, 0)
    if pos == len(nums) - 1 {
        //fmt.Printf("last pos: %v\n", nums)
        fmt.Printf("%v\n", nums)
        result := make([]int, len(nums))
        copy(result, nums)
        return  [][]int{result}
    } else {
        for i := pos; i < len(nums); i++ {
            //fmt.Printf("nums: %v\n", nums)
            //fmt.Printf("exchange num %d[%d] %d[%d]\n", pos, nums[pos], i, nums[i])
            if i == pos {
                res := p(nums, pos + 1)
                ans = append(ans, res...)
            } else {
                tmp := nums[pos]
                nums[pos] = nums[i]
                nums[i] = tmp
                //fmt.Printf("nums: %v go to next... \n", nums)
                res := p(nums, pos + 1)
                ans = append(ans, res...)
                //fmt.Printf("ans: %v\n", ans)
                //fmt.Printf("nums: %v\n", nums)
                //fmt.Printf("exchange num %d[%d] %d[%d]\n", pos, nums[pos], i, nums[i])
                tmp = nums[pos]
                nums[pos] = nums[i]
                nums[i] = tmp
                //fmt.Printf("nums: %v\n", nums)
            }
        }
        return ans
    }
}

func dedup(ans [][]int) [][]int {
    uniq := make([][]int, 0)
    for _, s := range ans {
        in := false
        for _, u := range uniq {
            if eq(s, u) == 0 {
                in = true
                break
            }
        }
        if !in {
            uniq = append(uniq, s)
        }
    }
    return uniq
}

func eq(a, b []int) int {
    for i := 0; i < len(a); i++ {
        if a[i] < b[i] {
            return -1
        }
        if a[i] > b[i] {
            return 1
        }
    }
    return 0
}

func sort(nums []int) []int {
    for i := 1; i < len(nums); i++ {
        for j := i; j > 0 && nums[j] < nums[j - 1]; j-- {
            tmp := nums[j]
            nums[j] = nums[j - 1]
            nums[j - 1] = tmp
        }
    }
    return nums
}
