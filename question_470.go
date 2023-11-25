package leetcode

func rand10() int {
    for {
        x := rand7()-1 + 7 * (rand7() -1)
        if x >= 1 && x <= 40 {
            return 1+x%10
        }
    }
}
