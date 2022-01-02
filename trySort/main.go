package main

import "fmt"

func main() {
	/*
		不稳定排序
		选择	— O(n²) = n次O(n) = (n,n-1,n-2)
		堆  — O(nlogn); 						作用：取集合中前k大元素
		希尔	— O(nlogn)
		快排	— O(nlogn) 期望时间, O(n²) 最坏情况; 	作用：对于大的、乱数串行一般相信是最快的已知排序
	*/
	unStableSort()

	/*
		冒泡— O(n²)
		插入— O(n²)

	*/

}

func unStableSort() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(selectSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(headSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(quickSort(arr))
}
