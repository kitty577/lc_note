package main

import "math"

// 背包问题：给定一组物品，每种物品都有自己的重量和价格，在限定的总重量内，我们如何选择，才能使得物品的总价格最高。
// 一般问题： 我们有 n 件物品和一个容量 (capacity) 为 C背包，记第 i件物品的重量 (weight) 为 w[i],价值 (value)为 v[i] ,求将哪些物品装入背包可使价值总和最大。
//
// [0-1背包]： 如果限定每件物品最多只能选取 1 次（即 0 或 1 次），则问题称为 0-1背包问题。
// dp[i][j]表示前i件物品放入一个容量为j的背包获得的最大价值；
// 转移过程为： 不选第i件物品，---> dp[i][j]=dp[i-1][j]
// 		      选第i件物品， ---> dp[i][j]=dp[i-1][j-wi]+wi  // 注意第i件物品能放入背包的前提是：wi<=j
// ====> dp[i][j]=max(dp[i-1][j],dp[i-1][j-wi]+wi)

// [完全背包]： 如果每件物品最多可以选取无限次，则问题称为 完全背包问题。
//	每件物品可以被选择多次，因此 dp[i][j]dp[i][j]dp[i][j] 应为以下所有可能方案中的最大值：
// 第i件物品选0个的最大价值 dp[i−1][j]
// 第i件物品选1个的最大价值 dp[i−1][j−wi]+vi
// 第i件物品选2个的最大价值 dp[i−1][j−2⋅wi]+2⋅vi
// ......
// 第i件物品选k个的最大价值 dp[i−1][j−k⋅wi]+k⋅vi  //注意，第i件物品能放入k件的前提为：k⋅wi≤j

// 定义二维数组dp[len(nums)][amount]. dp[i][j]表示：从前 i种硬币中组成金额 j所需最少的硬币数量
// 对于第一列,也即组成0金额，也即一个硬币都不选 ---> j==0时，dp[i][j]=0
// 对于第一行，也即只能选/不选第一个硬币  ---> i==0时，dp[0][j]=math.MaxInt64（当j不是第i枚硬币面值的倍数时) dp[0][j]=j/coins[i]（是倍数关系，能够凑到）
// i>0&&j>0 ---> 当前 number:=0;number*面值<=j;number++

func coinChange(coins []int, amount int) int {
	length := len(coins)
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32
		}
	}

	// j=0
	for i := 0; i < length; i++ {
		dp[i][0] = 0
	}

	// i=0
	for number := 0; number*coins[0] <= amount; number++ {
		dp[0][number*coins[0]] = number
	}

	// i>0&&j>0
	for i := 1; i < length; i++ {
		for j := 1; j <= amount; j++ {
			for number := 0; number*coins[i] <= j; number++ {
				dp[i][j] = min(dp[i][j], number+dp[i-1][j-number*coins[i]])
			}
		}
	}

	if dp[length-1][amount] == math.MaxInt32 {
		return -1
	}

	return dp[length-1][amount]
}
