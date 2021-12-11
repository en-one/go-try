package main

// selectSort 选择--------------------------------------------------
// 手里一把扑克牌，把最小的放在最左边，然后从下一张开始，接着找最小的。
func selectSort(arr []int) []int {
	var result []int
	len := len(arr)
	for i := 0; i < len; i++ {
		smallest_index := findSmallest(arr)
		result = append(result, arr[smallest_index])
		arr = append(arr[:smallest_index], arr[smallest_index+1:]...)
	}

	return result
}

func findSmallest(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	smallest := arr[0]
	smallest_index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index
}

// quickSort 快排---------------------------------------------------
//
// 以第一个为基准，大于第一个放右边，小于第一个放左边，分别对左右递归，然后集合
func quickSort(nums []int, head, tail int) []int {
	if head >= tail {
		return nums
	}

	l, r := head, tail
	mid := (r-l)/2 + l
	pivot := nums[mid]

	for l <= r {
		for l <= r && nums[l] < pivot {
			l++
		}
		for l <= r && nums[r] > pivot {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	quickSort(nums, head, r)
	quickSort(nums, l, tail)

	return nums

}

func quickSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort1(left)
	right = quickSort1(right)
	return append(left, append([]int{pivot}, right...)...)
}

// 堆排序---------------------------------------------
// O(nlogn)

func headSort(num []int) []int {
	// 构建二叉堆
	len := len(num)
	index := (len - 1) / 2
	for i := index; i >= 0; i-- {
		heapfiy(i, len-1, num)
	}
	// 排序
	for i := len - 1; i >= 0; i-- {
		if num[0] < num[i] {
			num[0], num[i] = num[i], num[0]
			heapfiy(0, i-1, num)
		}
	}

	return num
}

func heapfiy(index, end int, arr []int) []int {
	for {
		child := 2*index + 1
		if child > end {
			break
		}
		if child+1 < end && arr[child] > arr[child+1] {
			child++
		}
		if child < end && arr[index] > arr[child] {
			arr[index], arr[child] = arr[child], arr[index]
			index = child
		} else {
			break
		}
	}
	return arr

}
