package main

// 二分查找
func searchRange(nums []int, target int) []int {
    if len(nums)==0||nums[0]>target{
        return []int{-1,-1}
    }

    targetmid:=-1

    left,right:=0,len(nums)-1
    for left<=right{
        mid:=left+(right-left)/2
        if nums[mid]==target{
            targetmid=mid
            break
        }else if nums[mid]>target{
            right=mid-1
        }else if nums[mid]<target{
            left=mid+1
        }
    }

    if targetmid==-1{  // nums中不存在target
        return []int{-1,-1}
    }

    // 找targetmid的左右两边 确认开始位置和结束位置

    start,end:=targetmid,targetmid
    for i:=targetmid;i>=0;i--{
        if nums[i]==target{
            start=i
        }else{
            break
        }
    }

    for i:=targetmid;i<len(nums);i++{
        if nums[i]==target{
            end=i
        }else{
            break
        }
    }

    return []int{start,end}
}