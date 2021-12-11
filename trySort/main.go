package main

import "fmt"

func main() {
	/*
		不稳定排序
		选择— O(n²) = n次O(n) = (n,n-1,n-2)
		希尔— O(nlogn)
		堆— O(nlogn)
		快排— O(nlogn) 期望时间, O(n²) 最坏情况; 对于大的、乱数串行一般相信是最快的已知排序
	*/
	unStableSort()

}

func unStableSort() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(selectSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(headSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println(quickSort1(arr))
}
