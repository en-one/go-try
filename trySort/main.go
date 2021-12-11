package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	//fmt.Println(quickSort(arr, 0, len(arr)-1))
	fmt.Println(headSort(arr))
}

// 快排
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

	for i := 0; i <= len(arr); i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(quickSort1(left), append([]int{pivot}, quickSort1(right)...)...)
}

// 堆排序

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
