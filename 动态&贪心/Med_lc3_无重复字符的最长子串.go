package main

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
    visited:=make(map[byte]int)
    start,end,ans:=0,0,0
    if len(s)==0{
        return 0
    }

    for end<len(s){ 
        if visited[s[end]]==0{  // 没有出现重复，那么加进去，扩充窗口
            visited[s[end]]++
            end++
            ans=max(ans,end-start)
        }else{  // 出现重复字符了，收缩窗口
            visited[s[start]]--
            start++
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