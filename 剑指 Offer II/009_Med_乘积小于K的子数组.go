package main

func numSubarrayProductLessThanK(nums []int, k int) int {
    start,multi:=0,1

    ans:=0

    for end,v:=range nums{
        multi*=v
        for ;start<=end&&multi>=k;start++{
            multi=multi/nums[start]
        }

        // 由于其中nums中的每个数都>=1
        ans+=end-start+1
    }
    return ans
}