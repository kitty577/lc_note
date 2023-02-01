# 背包问题

## 01背包  （只能放入一次）

### 题
有n件物品和一个最多能背重量为w 的背包。第i件物品的重量是weight[i]，得到的价值是value[i] 。每件物品只能用一次，求解将哪些物品装入背包里物品价值总和最大。

### 解
1、确定dp数组以及下标含义：
   dp[i][j]：表示从下标为[0~i]的物品里任意取，放进容量为j的背包，价值总和最大是多少

   ps:: 也可一维数组，dp[j]：容量为j的背包的最大价值

2、递推公式
    对每个物品都有选和不选
    不选：dp[i-1][j]
    选：dp[i-1][j-weight[i]]+value[i]

    dp[i][j]=max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])

3、dp初始化
    首先，如果背包容量j为0，即dp[i][0]=0
    dp[0][j]:即存放编号为0的物品，各个容量的背包所能存放的最大价值
    如j>=weight[0],dp[0][j]=value[0]
    如j< weight[0],,即存放不了编号为0的物品，dp[0][j]=0
    
    for j:=0;j< weight[0];j++{
        dp[0][j]=0
    }
    for j:=weight[0];j<=bagweight;j++{
        dp[0][j]=value[0]
    }

4、确定遍历顺序
    dp[i][j]:取编号为i的物品，放进容量为j的背包，最大价值

    明显，有两个遍历维度，物品、背包容量

func bag(weight,value []int,bagweight int)int{
    // 定义dp数组
    dp:=make([][]int,len(weight))
    for i:=0;i< len(weight);i++{
        dp[i]=make([]int,bagweight+1)
    }

    // dp初始化
    for j:=weight[0];j<=bagweight;j++{
        dp[0][j]=value[0]
    }

    for i:=1;i<len(weight);i++{   // 遍历物品
        for j:=0;j<=bagweight;j++{  // 遍历背包容量
            if j<weight[i]{   // 装不下
                dp[i][j]=dp[i-1][j]
            }else{  // 装的下，可以考虑要不要装
                dp[i][j]=max(dp[i-1][j],dp[i-1][j-weight[i]]+value[i])
            }
        }
    }
    return dp[len(weight)-1][bagweight]
}