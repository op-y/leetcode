package leetcode

func containsDuplicate(nums []int) bool {
    if nums == nil {
        return false
    }

    length := len(nums)
    if length < 1 {
        return false
    }
    for i:=0; i<length-1; i++ {
        j := i+1
        for j>0 && nums[j]<nums[j-1]{
            tmp := nums[j]
            nums[j] = nums[j-1]
            nums[j-1] = tmp
            j--
        }
        if j>0 && nums[j] == nums[j-1] {
            return true
        }
    }
    return false
}
