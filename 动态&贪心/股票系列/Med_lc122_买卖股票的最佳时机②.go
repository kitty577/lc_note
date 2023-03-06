package main

func maxProfit(prices []int) int {
    maxProfit:=0
    
    for i:=1;i<len(prices);i++{
        if prices[i]-prices[i-1]>0{
            maxProfit+=prices[i]-prices[i-1]
        }
    }
    return maxProfit
}