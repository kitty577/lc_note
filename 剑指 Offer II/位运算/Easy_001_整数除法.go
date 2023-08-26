package main

import "math"

// 除法 可以转化为减法，被减的次数就等于商
// 可以将被除数成倍递增  如: a=15，b=2,c=1  ==> b=4 c=2 ==> b=8,c=4 ==》b=16 c=8 此时16>a 退出   当前剩余减数 15-8=7，累计除的次数=4
// 由于7>2,对其进行上述过程 a=7,b=2,c=1 ==> b=4,c=2 ==>b=8,c=4 此时8>a 退出 当前剩余减数7-4=3 累计除的次数=2
// 由于3>2 对其进行上述过程 a=3,b=2,c=1 ==> b=4,c=2 此时4>3 退出 当前剩余减数3-2=1
func divide(a int, b int) int {
    if a==0{
        return 0
    }

    if a==math.MinInt32{    // 被除数为最小值
        if b==1{    
            return math.MinInt32
        }
        if b==-1{
            return math.MaxInt32
        }
    }

    var needRever bool
    if a>0&&b<0||a<0&&b>0{
        needRever=true
    }
    
    if a<0{
        a=-a
    }
    if b<0{
        b=-b
    }
    
    res:=div(a,b)
    if needRever{
        return -res
    }
    return res
}

func div(a,b int)(ans int){
    for a>=b{
        d:=b
        c:=1
        for d+d<=a{
            d+=d
            c+=c
        }
        a=a-d
        ans+=c
    }
    return ans
}

