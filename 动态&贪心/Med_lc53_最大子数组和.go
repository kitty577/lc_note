package main

// 方法一：贪心

// 局部最优：当前“连续和”为负数的时候立刻放弃，从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。

// 全局最优：选取最大“连续和”

// func maxSubArray(nums []int) int {
//     maxSum:=nums[0]
//     for i:=1;i<len(nums);i++{
//         if nums[i]+nums[i-1]>nums[i]{
//             nums[i]+=nums[i-1]
//         }
//         if nums[i]>maxSum{
//             maxSum=nums[i]
//         }
//     }
//     return maxSum
// }


// 方法二：动态规划
func maxSubArray(nums []int) int {
    dp:=make([]int,len(nums))  // 截止到i位置 的最大值
    maxSum:=nums[0]
    dp[0]=nums[0]
    for i:=1;i<len(nums);i++{
        dp[i]=max(dp[i-1]+nums[i],nums[i])
        maxSum=max(maxSum,dp[i])
    }
    return maxSum
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}