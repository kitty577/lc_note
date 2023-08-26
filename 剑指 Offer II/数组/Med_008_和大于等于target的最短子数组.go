package main

func minSubArrayLen(target int, nums []int) int {
    ans:=len(nums)+1

    start,end,sum:=0,0,0

    for end<len(nums){
        sum+=nums[end]
        for sum>=target{
            ans=min(ans,end-start+1)
            sum-=nums[start]
            start++
        }
        
        end++
    }

    if ans==len(nums)+1{
        return 0
    }
    return ans       
}

func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}