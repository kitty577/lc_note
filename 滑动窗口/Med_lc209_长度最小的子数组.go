package main

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
    start,end,sum:=0,0,0
    minLen:=len(nums)+1

    for end<len(nums){
        sum+=nums[end]
        for sum>=target{
            sum-=nums[start]
            minLen=min(minLen,end-start+1)
            start++
        }
        end++
    }

    if minLen==len(nums)+1{
        return 0
    }


    return minLen

}

func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}