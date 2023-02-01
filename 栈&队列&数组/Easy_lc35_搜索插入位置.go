package main

// 二分查找
func searchInsert(nums []int, target int) int {
    left,right:=0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)/2

        if nums[mid]==target{
            return mid
        }else if nums[mid]>target{
            right=mid-1
        }else if nums[mid]<target{
            left=mid+1
        }
    }
    // 在nums所有数前
    // 在nums中
    // 在nums所有数后
    
    return right+1
}