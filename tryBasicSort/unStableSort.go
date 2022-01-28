package main

// selectSort 选择--------------------------------------------------
// 手里一把扑克牌，把最小的放在最左边，然后从下一张开始，接着找最小的。
func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}

// 希尔排序---------------------------------------------
//  加强版分组插入，按3x+1进行递增排序
func shellSort(arr []int) []int {
	h := 1
	for h < len(arr)/3 {
		h = 3*h + 1 // 3x+1为一个合适的分组状态
	}

	for h >= 1 {
		// 插入排序
		for i := h; i < len(arr); i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h /= 3
	}
	return arr
}

// quickSort 快排---------------------------------------------------
// 以第一个为基准，大于第一个放右边，小于第一个放左边，分别对左右递归，然后集合
func quickSort(arr []int) []int {
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

	left = quickSort(left)
	right = quickSort(right)
	return append(left, append([]int{pivot}, right...)...)
}

// 堆排序---------------------------------------------
// 基于堆的排序其实是一种选择排序
func headSort(arr []int) []int {
	// 构建二叉堆
	len := len(arr)
	last_node := len - 1
	last_node_parent := last_node / 2
	for i := last_node_parent; i >= 0; i-- {
		heapify(arr, len, i)
	}
	// 排序
	// 把最大的放最后，然后重新建堆
	for i := len - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}

	return arr
}

// 大顶堆
func heapify(tree []int, num, index int) []int {
	if index >= num {
		return tree
	}
	lchild := 2*index + 1
	rchild := 2*index + 2
	max := index

	if lchild < num && tree[lchild] > tree[max] {
		max = lchild
	}

	if rchild < num && tree[rchild] > tree[max] {
		max = rchild
	}

	if max != index {
		tree[index], tree[max] = tree[max], tree[index]
		heapify(tree, num, max)
	}
	return tree

}
