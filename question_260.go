package leetcode

func singleNumber260(nums []int) []int {
    x := nums[0]
    for i:=1; i<len(nums); i++ {
        x ^= nums[i]
    }

    b := 1
    for x & b == 0 {
        b <<= 1
    }

    n1, n2 := 0, 0
    for i:=0; i<len(nums); i++ {
        if b & nums[i] == 0 {
            n1 ^= nums[i]
        } else {
            n2 ^= nums[i]
        }
    }

    return []int{n1, n2}
}
