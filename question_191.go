package leetcode

//func hammingWeight(num uint32) int {
//    bit1 := 0
//    for num > 0 {
//        if num % 2 > 0 {
//            bit1++
//        }
//        num /= 2
//    }
//    return bit1
//}

func hammingWeight(num uint32) int {
    bit1 := 0
    for num != 0 {
        bit1++
        num = num & (num-1)
    }
    return bit1
}
