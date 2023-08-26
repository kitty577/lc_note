package main

// 对于判断两字符串是否含有相同字符 可以使用位运算
func maxProduct(words []string) int {
    masks:=make([]int,len(words))
    for i,word:=range words{
        for _,ch:=range word{
            masks[i]|=1<<(ch-'a')
        }
    }
    var ans int

    for i,x:=range masks{
        for j:=i+1;j<len(masks);j++{
            if x&masks[j]==0{   // 每个位上的数不相同==> 不含相同字符
                ans=max(ans,len(words[i])*len(words[j]))
            }
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