package main

// 滑动窗口
func minWindow(s string, t string) string {
    if len(t)==0||len(s)==0||len(t)>len(s){
        return ""
    }

    //left窗口左边界，right窗口有边界，start是有效窗口的起始位置
    left,right,start:=0,0,0
    // size 记录覆盖子串的最小长度
    size:=len(s)+1
    needMap:=map[byte]int{}
    // 需要的字符长度
    needCount:=len(t)

    for i:=0;i<len(t);i++{
        needMap[t[i]]++
    }

    for right<len(s){
        if needMap[s[right]]>0{   // 该字符是需要的
            needCount--
        }
        needMap[s[right]]--

        if needCount==0{  // 已经收集齐了所有需要的字符
            // 消除不需要的字符以达到最小覆盖子串
            for left<right&&needMap[s[left]]<0{
                needMap[s[left]]++
                left++
            }

            // 更新长度
            if right-left+1<size{
                size=right-left+1
                start=left
            }

            //经过以上的消除操作，此时的窗口 左边界和右边界都是需要的字符。进行for循环计算下一个窗口，即需要左边界右移动
            needMap[s[left]]++
            needCount++
            left++
        }

        right++
    }

    if size==len(s)+1{
        return ""
    }

    return s[start:start+size]
}

// 2023.01.10