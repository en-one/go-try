package main

// insertSort 插入--------------------------------------------------
// 手里一把扑克牌，从左手边第二张（0是第一张）开始，向左比较；然后从左手边第三张开始.....
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

// mergeSort 归并--------------------------------------------
func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var middle = len(arr) / 2

	var left = mergeSort(arr[:middle])
	var right = mergeSort(arr[middle:])

	return merge(left, right)
}

// 借鉴合并两个有序列链表
func merge(left, right []int) []int {
	var res = make([]int, len(left)+len(right))
	var i, j = 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			//res = append(res, left[i])  不行
			res[i+j] = left[i]
			i++
		} else {
			res[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		res[i+j] = left[i]
		i++
	}

	for j < len(right) {
		res[i+j] = right[j]
		j++
	}

	return res
}
