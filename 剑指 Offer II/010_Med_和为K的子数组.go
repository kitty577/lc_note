package main

// 前缀和就是 从区间起始位置到现在位置之间的区间元素之和
// 这题可以利用前缀和
// 把前缀和看成一个线段。(位置0到i)pre = (位置0到i之间的某一段区间)a+(所求区间)k
func subarraySum(nums []int, k int) int {
    count:=0
    m:=make(map[int]int,0)
    pre:=0
    m[0]=1

    for _,v:=range nums{
        pre+=v
        if _,ok:=m[pre-k];ok{
            count+=m[pre-k]
        }

        m[pre]++
    }
    return count
}