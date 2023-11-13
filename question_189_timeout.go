package leetcode

//func rotate(nums []int, k int)  {
//	if k == 0 {
//		return
//	}
//
//	l := len(nums)
//	if l == 1 {
//		return
//	}
//
//	// l的倍数会回到原地
//	k = k % l
//
//	for i:=0; i<k; i++ {
//		ele := nums[l-1]
//		for j:=l-1; j>0; j-- {
//			nums[j] = nums[j-1]
//		}
//		nums[0] = ele
//	}
//
//	return
//}
