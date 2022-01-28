package main

// bubbleSort 冒泡--------------------------------------------
// 手里一把扑克牌，第一轮，把最大的放在右边；第二轮，排除最右边一张，接着从0开始找最大，放在最右边...
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ { // 控制范围，每次排除最右边已经排好序的内容
		for j := 0; j < len(arr)-1-i; j++ { // 确定好最右边内容，然后从最左边开始排序
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// insertSort 插入--------------------------------------------------

// 手里一把扑克牌，从左手边第二张（0是第一张）开始，向左比较；然后从左手边第三张开始.....
func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

// mergeSort 归并--------------------------------------------------
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

// countingSort 计数-----------------------------------------------
// 局限性：1、不适用于非整数排序
// 找出数组的最大值和最小值，创建新数组[max+1]， 原数组值等于新数组下标时， 新数组值加一，最后把新数组下标按值的个数输出即排序完成

func countingSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}

	min, max := CountMaxMin(arr)

	temp := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		temp[arr[i]]++
	}

	var index int
	for i := min; i < len(temp); i++ {
		for j := temp[i]; j > 0; j-- {
			arr[index] = i
			index++
		}
	}

	return arr

}

func CountMaxMin(data []int) (int, int) {
	min, max := data[0], data[0]
	for i := 0; i < len(data); i++ {
		if data[i] < min {
			min = data[i]
		}
		if data[i] > max {
			max = data[i]
		}
	}
	return min, max
}

// radixSort 基数-----------------------------------------------
// 调用然后收集

func radixSort(arr []int) []int {
	key := maxLimit(arr)
	tmp := make([]int, len(arr), len(arr))
	count := new([10]int)
	radix := 1
	var i, j, k int
	for i = 0; i < key; i++ { // 进行key次排序
		// 初始化桶
		for j = 0; j < 10; j++ {
			count[j] = 0
		}

		for j = 0; j < len(arr); j++ {
			k = (arr[j] / radix) % 10
			count[k]++
		}
		// 叠加之后下标减一就是该元素在新数组中位置
		for j = 1; j < 10; j++ {
			count[j] = count[j-1] + count[j]
		}
		// 倒叙 保持当相同的桶中的元素后面的元素在后面
		for j = len(arr) - 1; j >= 0; j-- {
			k = (arr[j] / radix) % 10
			tmp[count[k]-1] = arr[j]
			count[k]--
		}
		for j = 0; j < len(arr); j++ {
			arr[j] = tmp[j]
		}

		radix = radix * 10
	}
	return arr
}

// 计算最大的数一共有几位
func maxLimit(arr []int) int {
	ret := 1
	var key int = 10
	for i := 0; i < len(arr); i++ {
		for arr[i] >= key {
			key = key * 10
			ret++
		}
	}
	return ret
}
