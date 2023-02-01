package main

import "math"

// 动态规划
func coinChange(coins []int, amount int) int {
    if len(coins)==0{
        return -1
    }

    // memo[n]： 表示凑成总金额为n所需要的最少硬币个数
    memo:=make([]int,amount+1)
    memo[0]=0
    for i:=1;i<=amount;i++{
        min:=math.MaxInt32
        for j:=0;j<len(coins);j++{
            if i-coins[j]>=0&&memo[i-coins[j]]<min{
                min=memo[i-coins[j]]+1
            }
        }
        memo[i]=min
    }

    if memo[amount]==math.MaxInt32{
        return -1
    }
    return memo[amount]
}



