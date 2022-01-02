package main

import "fmt"

func main() {
	/*
		不稳定排序
		选择	— O(n²) = n次O(n) = (n,n-1,n-2)
		堆  — O(nlogn); 						作用：取集合中前k大元素
		希尔	— O(nlogn);
		快排	— O(nlogn) 期望时间, O(n²) 最坏情况; 	作用：对于大的、乱数串行一般相信是最快的已知排序
	*/
	unStableSort()

	/*
		稳定排序
		冒泡— O(n²)
		插入— O(n²)
		归并— O(nlogn);
		计数— O(n + m) 其中 m 为原始数组的范围
		基数— O(n²)
		桶— O(n²)

	*/
	stableSort()
}

func unStableSort() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("selectSort", selectSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("headSort", headSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("shellSort", shellSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("quickSort", quickSort(arr))
}

func stableSort() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("insertSort", insertSort(arr))
}
