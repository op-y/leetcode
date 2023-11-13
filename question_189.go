package leetcode

//func rotate189(nums []int, k int)  {
//    if k == 0 {
//        return
//    }
//
//    l := len(nums)
//    if l == 1 {
//        return
//    }
//
//    // l的倍数会回到原地
//    k = k % l
//
//    nums1 := nums[:l-k]
//    nums2 := nums[l-k:]
//    newnums := []int{}
//    newnums = append(newnums, nums2...)
//    newnums = append(newnums, nums1...)
//    copy(nums, newnums)
//}

func rotate189(nums []int, k int)  {
    if k == 0 {
        return
    }

    l := len(nums)
    if l == 1 {
        return
    }

    // l的倍数会回到原地
    k = k % l

    for i,j := 0, l-1; i<j; i,j=i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    for i,j := 0, k-1; i<j; i,j=i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    for i,j := k, l-1; i<j; i,j=i+1, j-1 {
        nums[i], nums[j] = nums[j], nums[i]
    }
}
