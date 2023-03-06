package main


// 贪心

// 局部最优：删除单调坡度上的节点（不包括单调坡度两端的节点），那么这个坡度就可以有两个局部峰值。

// 整体最优：整个序列有最多的局部峰值，从而达到最长摆动序列。


// 情况一：上下坡中有平坡
// 情况二：数组首尾两端
// 情况三：单调坡中有平坡

func wiggleMaxLength(nums []int) int {
    length:=len(nums)

    if length<2{
        return length
    }

    ans:=1
    preDiff:=nums[1]-nums[0]
    if preDiff!=0{
        ans=2
    }

    for i:=2;i<length;i++{
        diff:=nums[i]-nums[i-1]
        if diff>0&&preDiff<=0||diff<0&&preDiff>=0{
            ans++
            preDiff=diff
        }
    }
    return ans
}