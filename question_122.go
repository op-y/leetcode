package leetcode

func maxProfit(prices []int) int {
    if prices == nil {
        return 0
    }
    length := len(prices)
    if length <= 1 {
        return 0
    }

    maxIncome := 0
    buy := 0
    sale := 0
    for i:=1; i<length; i++ {
        if prices[i] > prices[sale] {
            sale = i
            if i == length - 1 {
                maxIncome += prices[sale] - prices[buy]
            }
        } else {
            if buy == sale {
                buy = i
                sale = i
            } else {
                maxIncome += prices[sale] - prices[buy]
                buy = i
                sale = i
            }
        }
    }
    return maxIncome
}
