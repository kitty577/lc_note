package main

// 记录下每个字符出现的最远边界
func partitionLabels(s string) []int {
    ans:=[]int{}
    var mark [26]int
    size,left,right:=len(s),0,0
    for i:=range s{
        mark[s[i]-'a']=i
    }
    for i:=0;i<size;i++{
        right=max(right,mark[s[i]-'a'])
        if i==right{
            ans=append(ans,right-left+1)
            left=right+1
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