package leetcode

func maxProfit124(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	l := len(prices)
	// 可能的4种状态转移
    for i:=1; i<l; i++ {
        buy1 = max(buy1, -prices[i])
        sell1 = max(sell1, buy1+prices[i])
        buy2 = max(buy2, sell1-prices[i])
        sell2 = max(sell2, buy2+prices[i])
    }
	return sell2
}
