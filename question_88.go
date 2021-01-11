package leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := 0
    j := 0
    length := m
    for i < length && j < n {
        if nums1[i] <= nums2[j] {
            i++
        } else {
            length++
            for k:=length-1; k>i;k-- {
                nums1[k] = nums1[k-1]
            }
            nums1[i] = nums2[j]
            j++
        }
    }
    if j < n {
        for j < n {
            nums1[i] = nums2[j]
            i++
            j++
        }
    }
}
