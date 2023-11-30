package leetcode

func countDigitOne(n int) int {
    ans := 0
    k, mulk := 0, 1
    for n >= mulk {
        ans += (n/(mulk*10))*mulk + min(max(n%(mulk*10)-mulk+1, 0), mulk) // 关键是这里的分析过程
        mulk *= 10
        k++
    }
    return ans
}
