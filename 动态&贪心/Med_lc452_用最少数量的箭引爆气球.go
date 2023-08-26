package main

import "sort"

func findMinArrowShots(points [][]int) int {
    length:=len(points)
    if length<2{
        return length 
    }
    // 先排序
    sort.Slice(points,func(i,j int)bool{
        return points[i][1]<points[j][1]
    })
    ans:=1
    maxRight:=points[0][1]
    for _,p:=range points{
        if p[0]>maxRight{
            ans++
            maxRight=p[1]
        }
    }
    
    return ans
}
