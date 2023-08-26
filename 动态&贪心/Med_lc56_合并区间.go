package main

import "sort"

func merge(intervals [][]int) [][]int {
    length:=len(intervals)
    if length<2{
        return intervals
    }

    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0]<intervals[j][0]
    })

    ans:=[][]int{intervals[0]}
    left,right:=0,0

    for i:=1;i<length;i++{
        pre:=ans[len(ans)-1]
        if intervals[i][0]>pre[1]{
            ans=append(ans,intervals[i])
        }else{
            ans=ans[:len(ans)-1]
            left=min(pre[0],intervals[i][0])
            right=max(pre[1],intervals[i][1])
            ans=append(ans,[]int{left,right})
        }
    }
    return ans   
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

func min(a,b int)int{
    if a>b{
        return b
    }
    return a
}