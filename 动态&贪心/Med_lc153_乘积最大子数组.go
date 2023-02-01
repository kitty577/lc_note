package main 

// 动态规划
// dp[i]:表示截止到i位置的最大乘积
// 与求和不同，乘积要注意正负;因为可以有负数，最大值可能由之前的最小值乘当前的负数得到

// 如果这个数为负数，那前面负的越多越好
// 如果这个数是正数，那前面正的越多越好

const inf int = 0x3f3f3f
func maxProduct(nums []int) int {
    preMax, preMin, ans := 1, 1, -inf
    for _, num := range nums {
        preMax, preMin = max(preMax * num, preMin * num, num), min(preMax * num, preMin * num, num)
        ans = max(preMax, ans)
    }
    return ans
}

func max(vals ...int) int {
    ans := -inf
    for _, v := range vals {
        if v > ans {
            ans = v
        }
    }
    return ans
}

func min(vals ...int) int {
    ans := inf
    for _, v := range vals {
        if v < ans {
            ans = v
        }
    }
    return ans
} 


