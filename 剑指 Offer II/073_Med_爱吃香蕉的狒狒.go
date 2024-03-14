package main

// 二分查找 最小吃1根，最多吃数组中的最大值
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, count := range piles {
		if count > right {
			right = count
		}
	}

	// for left<=right{
	//     mid:=left+(right-left)/2
	//     spend:=checkTime(piles,mid)
	//     if spend==h{
	//         return mid
	//     }else if spend>h{
	//         left=mid+1
	//     }else if spend<h{
	//         right=mid-1
	//     }
	// }
	for left < right {
		mid := left + (right-left)/2
		spend := checkTime(piles, mid)
		if spend > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func checkTime(piles []int, v int) int {
	spend := 0
	for _, p := range piles {
		d, mod := p/v, p%v
		if mod > 0 {
			d++
		}
		spend += d
	}
	return spend
}
