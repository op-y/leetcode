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

// 官方体检 更简洁一点
//class Solution:
//    def maxProfit(self, prices: List[int]) -> int:
//        inf = int(1e9)
//        minprice = inf
//        maxprofit = 0
//        for price in prices:
//            maxprofit = max(price - minprice, maxprofit)
//            minprice = min(price, minprice)
//        return maxprofit
