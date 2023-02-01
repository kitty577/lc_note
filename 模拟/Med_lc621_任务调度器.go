package main

// 根据tasks建立[26]int{}数组，用于记录任务的次数
// 对次数数组进行排序，
// 优先处理个数多的任务，result=(n+1)*(count-1)+1
//       ==>A->X->X->A->X->X->A (X为其他任务或者待命)
// 再排序下一个任务，如果下一任务与最大任务数一致，result++
//       ==>A->B->X->A->B->X->A->B (X为其他任务或者待命)
// 如果空位都插满后，还有任务，则直接插入

func leastInterval(tasks []byte, n int) int {
	if len(tasks) <= 1 || n < 1 {
		return len(tasks)
	}
	nums := [26]int{}
	for _, task := range tasks {
		nums[task-'A']++
	}

	// 找到执行次数最多的任务
	maxTask, maxNums := 0, 0
	for _, c := range nums {
		if c > maxTask {
			maxTask, maxNums = c, 0
		} else if c == maxTask {
			maxNums++
		}
	}

	minTime := (n+1)*(maxTask-1) + 1 + maxNums

	return max(minTime, len(tasks))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func quickSort(arr [26]int) [26]int {
	// 找一个基准 大于的放右边 小于的放左边
	return _quickSort(arr, 0, len(arr)-1)
}
func _quickSort(arr [26]int, left, right int) [26]int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}
	return arr
}

func partition(arr [26]int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left], arr[right] = arr[right], arr[left]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	return left
}

// 2022.10.21
