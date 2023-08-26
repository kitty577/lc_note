package main

// 单调栈

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    m:=make(map[int]int,0)  // key为nums2中元素值，value为下一个更大的元素值
    ans:=make([]int,len(nums1))
    stack:=[]int{}  // 模拟栈 存下一个更大元素的值
    for _,v:=range nums2{
        for len(stack)!=0&&v>stack[len(stack)-1]{
            m[stack[len(stack)-1]]=v
            stack=stack[:len(stack)-1]
        }
        stack=append(stack,v)
    }
    // fmt.Println(m)

    for i,v:=range nums1{
        if _,ok:=m[v];ok{
            ans[i]=m[v]
        }else{
            ans[i]=-1
        }
    }
    return ans
}