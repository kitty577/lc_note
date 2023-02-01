package main

// 遍历 将结点值存储到数组中 再判断回文
func isPalindrome(head *ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

// 2022.12.17
