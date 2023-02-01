package main

// 与打家劫舍①不同的是房屋首尾相连 : 即有首无尾 或者无首有尾
// 可以转化为打家劫舍I的处理： nums[:len(nums)-1]  nums[1:]
func rob(nums []int) int {
    length:=len(nums)
    if length==0{
        return 0
    }
    if length==1{
        return nums[0]
    }
    if length==2{
        return max(nums[0],nums[1])
    }
    return max(_rob(nums[:length-1]),_rob(nums[1:]))
}

func _rob(nums []int)int{
    first,second:=nums[0],max(nums[0],nums[1])
    for i:=2;i<len(nums);i++{
        tmp:=second
        second=max(first+nums[i],second)
        first=tmp
    }
    return second
}

func max(a,b int)int{
    if a>b{
        return a
    }
    return b
}