package main

// 动态规划 
// dp[i] 表示拆分数i，得到的最大乘积
func integerBreak(n int) int {
    dp:=make([]int,n+1)
    dp[1]=1
    dp[2]=1 
    for i:=3;i<=n;i++{
        // i可以拆分成j和i-j
        // j从1开始！！如果从0开始的话，0与任何数相乘为0 
        for j:=1;j<i;j++{
            dp[i]=max(dp[i],max(j*dp[i-j],j*(i-j)))
        }
    }
    return dp[n]
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

// 2022.12.29