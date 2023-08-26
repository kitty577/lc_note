package main
import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][1]<intervals[j][1]
    })
    ans:=0
    pre:=intervals[0]
    for i:=1;i<len(intervals);i++{
        if intervals[i][0]<pre[1]{
            ans++
        }else{
            pre=intervals[i]
        }
    }
    return ans
}