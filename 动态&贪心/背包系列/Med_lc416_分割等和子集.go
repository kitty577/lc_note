package main

// 01背包

func canPartition(nums []int) bool {
    sum:=0
    for _,v:=range nums{
        sum+=v
    }
    if sum%2!=0{  // 和为奇数，无法拆分成两个数组
        return false
    }

    target:=sum/2   // 容量
    
    dp:=make([]int,target+1)
    for _,num:=range nums{
        for j:=target;j>=num;j--{
            if dp[j]<dp[j-num]+num{
                dp[j]=dp[j-num]+num
            }
        }
    }
    return dp[target]==target
}

 // 二维dp写法：
func canPartition2(nums []int) bool {
    /**
    动态五部曲：
        1.确定dp数组和下标含义
        2.确定递推公式
        3.dp数组初始化
        4.dp遍历顺序
        5.打印
    **/
    //确定和
    var sum int
    for _,v:=range nums{
        sum+=v
    }
    if sum%2!=0{   //如果和为奇数，则不可能分成两个相等的数组
        return false
    }
    sum/=2
    //确定dp数组和下标含义
    var dp [][]bool //dp[i][j] 表示： 前i个石头是否总和不大于J
    //初始化数组
    dp=make([][]bool,len(nums)+1)
    for i:=range dp{
        dp[i]=make([]bool,sum+1)
        dp[i][0]=true
    }
    for i:=1;i<=len(nums);i++{
        for j:=1;j<=sum;j++{//j是固定总量
            if j>=nums[i-1]{//如果容量够用则可放入背包
                dp[i][j]=dp[i-1][j]||dp[i-1][j-nums[i-1]]
            }else{//如果容量不够用则不拿，维持前一个状态
                dp[i][j]=dp[i-1][j]
            }
        }
    }
    return dp[len(nums)][sum]
}