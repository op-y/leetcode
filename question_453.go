package leetcode

func minMoves(nums []int) int {
    moves := 0
    min := 0x7fffffffffffffff
    for i:=0; i<len(nums); i++ {
        if nums[i] > min {
            min = nums[i]
        }
    }

    for i:=0; i<len(nums); i++ {
        moves += nums[i] - min
    }

    return moves
}
