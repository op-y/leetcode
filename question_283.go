package leetcode

func moveZeroes(nums []int)  {
    if nums == nil {
        return
    }

    length := len(nums)
    if length <= 1 {
        return
    }

    head := 0
    tail := length - 1
    for nums[tail] == 0 && tail > 0 {
        tail--
    }

    for head < tail {
        if nums[head] != 0 {
            head++
            continue
        }
        for i:=head; i<tail; i++ {
            nums[i] = nums[i+1]
        }
        nums[tail] = 0
        tail--
    }

    return
}
