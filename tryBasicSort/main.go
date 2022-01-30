package main

import "fmt"

/*
	一些总结：
	第三梯队：冒泡排序、选择排序、插入排序
			1、冒泡和插入 的元素比较和交换次数取决于原始数组的有序程度
			2、选择为不稳定，其他稳定
			3、插入的性能略高于冒泡，冒泡两个元素之间的交换是独立的，而插入是连续的
			4、选择排序元素的交换次数和原始数组有序程度无关
	第二梯队：归并排序、希尔排序、堆排序、快排
			1、快排复杂度O(nlogn)，但极端情况下为O(n²)、堆和归并比较稳定
			2、归并排序为稳定排序，其他为不稳定
			3、快排和堆不需要建立额外的存储空间，而归并排序需要建立merge数组
			4、平均复杂度来说，堆逊色于快排和归并，因为：“二叉堆的父子节点在内存中并不连续”
	第一梯队：计数数序、桶排序、基数排序
			1、第一梯队的排序算法都是“线性复杂度的排序算法”
			2、都为稳定排序
*/

func main() {
	/*
		不稳定排序
		选择	— O(n²) = n次O(n) = (n,n-1,n-2)
		希尔	— O(nlogn); 最快可为O(n1.3)
		快排	— O(nlogn) 期望时间, O(n²) 最坏情况; 	作用：对于大的、乱数串行一般相信是最快的已知排序
		堆  — O(nlogn); 						作用：取集合中前k大元素
	*/
	unStableSort()

	/*
		稳定排序
		冒泡— O(n²)
		插入— O(n²)
		归并— O(nlogn);
		计数— O(n + m) 其中 m 为原始数组的范围
		基数— O(k(n+m)) k为元素最大位数，m为每一位的取值范围
		桶— O(n) 其中n为桶的数量

	*/
	stableSort()
}

func unStableSort() {
	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("selectSort", selectSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("shellSort", shellSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	//arr = []int{9, 45, 23, 45, 23, 34, 34, 2, 21, 5, 3, 2, 9, -5, 0, 3}
	fmt.Println("quickSort", quickSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("headSort", headSort(arr))
}

func stableSort() {

	arr := []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("bubbleSort", bubbleSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("insertSort", insertSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("mergeSort", mergeSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	//arr = []int{33, 11, 55, 7, 44, 1, 44}
	fmt.Println("countingSort", countingSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	fmt.Println("radixSort", radixSort(arr))

	arr = []int{5, 7, 1, 8, 3, 2, 6, 4, 9}
	//arr = []int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 18, 28, 12, 9}
	fmt.Println("bucketSort", bucketSort(arr))
}
