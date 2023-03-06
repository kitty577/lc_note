package main

// 饼干 胃口排序
// 大饼干喂胃口大的
func findContentChildren(child []int, food []int) int {
    // 从小到大排序
    sort.Ints(child)
    sort.Ints(food)

    ans:=0
    childNums:=len(child)
    foodNums:=len(food)

    for i,j:=0,0;i<childNums&&j<foodNums;i++{
        for j<foodNums&&child[i]>food[j]{
            j++
        }
        if j<foodNums{
            ans++
            j++
        }
    }
    return ans
}