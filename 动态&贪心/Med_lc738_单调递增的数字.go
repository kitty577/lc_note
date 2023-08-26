package main

import "strconv"

// 如果n本身非递增 需要处理。大致思路为低位变9，前一位降1。
// 如95 ==> 低位5变为9，高位9-1=8 ===》89
// 如642 ==> 低位2变9，高位4-1=3 ==》639 ==》再一次低位3变9，高位6-1=5 ===》599
func monotoneIncreasingDigits(n int) int {
    // 首先将数字转化为string 方便使用下标处理位
    s := strconv.Itoa(n)
    // 由于string类型的值是常量，不可更改。因此转化为byte类型，方便更改
    ss:=[]byte(s)

    nl:=len(ss)
    if nl<=1{
        return n
    }
    
    for i:=nl-1;i>0;i--{
        if ss[i]<ss[i-1]{
            ss[i-1]-=1 // 高位-1
            for j:=i;j<nl;j++{
                ss[j]='9' // 低位变9
            }
        }
    }
    ans,_:=strconv.Atoi(string(ss))
    return ans
}
