package main

// 由于 x 平方根的整数部分ans 是满足k*k≤x 的最大k 值，因此我们可以二分查找
func mySqrt(x int) int {
    left,right:=0,x
    for left<=right{
        mid:=left+(right-left)/2
        if mid*mid==x{
            return mid
        }else if mid*mid>x{
            right=mid-1
        }else if mid*mid<x{
            left=mid+1
        }
    }
    return right
}