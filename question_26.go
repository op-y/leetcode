package leetcode

func removeDuplicates(nums []int) int {
   if len(nums) == 0 {
        return -1
   }
   length := len(nums)
   pos := 1
   for i:=1; i < length; i++ {
       if nums[i] != nums[pos-1] {
           nums[pos] = nums[i]
           pos++
       }
   }
   return pos
}
