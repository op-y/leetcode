package leetcode

type NumArray struct {
    s []int
}


func Constructor(nums []int) NumArray {
    na := NumArray{s: nums}
    return na
}


func (this *NumArray) SumRange(i int, j int) int {
    sum := 0
    for idx:=i; idx<=j; idx++ {
        sum += this.s[idx]
    }
    return sum
}
