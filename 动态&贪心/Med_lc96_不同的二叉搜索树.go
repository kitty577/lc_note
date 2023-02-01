package main

// 动态规划
// dp[i]：表示i个不同的数 组成的二叉搜素数的个数

// 利用好二叉搜索树的特点

// 例如 假设i=3;则： 1 2 3 
// 当根结点为1时，左子树的数个数为0个，右边2个  dp[0]*dp[2]
// 当根结点为2时，左右各1个数  dp[1]*dp[1]
// 当根结点为3时，左边2个，右边0个  dp[2]*dp[0]

// =====》 如果以j为根结点，那边左边有j-1个数，右边有i-j个数
func numTrees(n int) int {
    dp:=make([]int,n+1)
    dp[0]=1
    dp[1]=1
    for i:=2;i<=n;i++{
        for j:=1;j<=i;j++{
            left:=dp[j-1]
            right:=dp[i-j]
            dp[i]+=left*right
        }
    }
    return dp[n]
} 