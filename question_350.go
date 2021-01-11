package leetcode

func intersect(nums1 []int, nums2 []int) []int {
    if nums1 == nil || nums2 == nil {
        return nil
    }

    items := make(map[int][]int, 0)
    result := make([]int, 0)

    for _, v :=  range nums1 {
        if l, ok := items[v]; ok {
            l[0] = l[0] + 1
        } else {
            l := make([]int, 2)
            l[0] = 1
            l[1] = 0
            items[v] = l
        }
    }

    for _, v :=  range nums2 {
        if l, ok := items[v]; ok {
            l[1] = l[1] + 1
        } else {
            l := make([]int, 2)
            l[0] = 0
            l[1] = 1
            items[v] = l
        }
    }

    for k, l := range items {
        if l[0] == 0 || l[1] == 0 {
            continue
        } else {
            if l[0] <= l[1] {
                for i:=0; i<l[0]; i++ {
                    result = append(result, k)
                }
            } else {
                for i:=0; i<l[1]; i++ {
                    result = append(result, k)
                }
            }
        }
    }

    return result
}
