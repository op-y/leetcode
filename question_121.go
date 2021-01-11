package leetcode

func maxProfit(prices []int) int {
    if prices == nil {
        return 0
    }
    length := len(prices)
    if length <= 1 {
        return 0
    }

    currentMin := 0
    maxIncome := 0
    for i:=1; i<length; i++ {
        if prices[i] - prices[currentMin] > maxIncome {
            maxIncome = prices[i] - prices[currentMin]
        }
        if prices[i] < prices[currentMin] {
            currentMin = i
        }
    }
    return maxIncome
}
