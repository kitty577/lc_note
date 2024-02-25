package main

// 双指针
func twoSum(numbers []int, target int) []int {
    if numbers[0]>target{
        return []int{}
    }
    length:=len(numbers)

    left,right:=0,length-1
    for left<right{
        if numbers[left]+numbers[right]==target{
            return []int{left,right}
        }else if numbers[left]+numbers[right]>target{
            right--
        }else if numbers[left]+numbers[right]<target{
            left++
        }
    }
    return []int{}
}