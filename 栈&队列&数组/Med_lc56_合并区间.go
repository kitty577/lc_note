package main

import "sort"

// 区间a,b能合并成一个区间的条件：
// a[1]>=b[0]

func merge(intervals [][]int) [][]int {
    // 首先要对intervals 从小到大排序（according to start）
    sort.Slice(intervals,func(i,j int)bool{
        return intervals[i][0]<intervals[j][0]
    })

    var res [][]int
    for _,interval:=range intervals{
        var length=len(res)

        if length==0||res[length-1][1]<interval[0]{   // 没有交集
            res=append(res,interval)
        }else{
            res[length-1][1]=max(res[length-1][1],interval[1])  // 右闭 取较大值
        }
    }
    return res
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}

// 2023.01.03