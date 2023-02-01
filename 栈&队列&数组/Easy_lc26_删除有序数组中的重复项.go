package main

func removeDuplicates(nums []int) int {
    //双指针 快指针遍历数组到达的下标位置；慢指针表示下一个不同的元素要填入的下标位置
    n:=len(nums)
    if n==0{
        return 0
    }
    slow:=1
    for fast:=1;fast<n;fast++{
        if nums[fast]!=nums[fast-1]{
            nums[slow]=nums[fast]
            slow++
        }
    }
    return slow
}

