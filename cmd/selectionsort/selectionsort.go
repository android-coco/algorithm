package main

import (
	"fmt"
	"algorithm/util"
	"time"
)

func main() {
	n := 10000

	//arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	arr := util.SortTestHelper(n, 0, n)
	//SelectionSort(arr, len(arr))
	//for _, v := range arr {
	//	fmt.Println(v)
	//}
	//SelectionSortByInterface(arr)
	util.TestSort("SelectionSortByInterface", SelectionSortByInterface, arr)
	//util.TestSort("bubbleSort", bubbleSort, util.IntArray(arr))
	fmt.Println("===========================int")
	//实现泛型
	arrint := util.IntArray{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	SelectionSortByInterface(arrint) //排序
	util.PrintArray(arrint)          //打印

	fmt.Println("===========================string")
	arrstr := util.StringArray{"hello", "i", "am", "go", "lang"}
	SelectionSortByInterface(arrstr)
	util.PrintArray(arrstr) //打印

	fmt.Println("===========================float64")
	//实现泛型
	arrFloat64 := util.Float64Array{10.0, 9.0, 8.0, 7.0, 6.0, 5.5, 4.3, 3.9, 2.0, 1.9, 0.8}
	SelectionSortByInterface(arrFloat64)
	util.PrintArray(arrFloat64) //打印

	fmt.Println("===========================struct")
	//实现泛型
	arrStu := util.StuArray{util.Student{Name: "yh", Score: 99}, util.Student{Name: "lyl", Score: 88}, util.Student{Name: "zj", Score: 60}, util.Student{Name: "B", Score: 77}, util.Student{Name: "A", Score: 77}}
	SelectionSortByInterface(arrStu)
	util.PrintArray(arrStu) //打印

}

//选择排序
func SelectionSort(arr []int, n int) {
	starttime := time.Now()
	for i := 0; i < n; i++ {
		//寻找[i,n]区间的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	fmt.Println("运行时间：", time.Now().Sub(starttime))
}

//选择排序
func SelectionSortByInterface(array util.Sortable) {
	for i := 0; i < array.Len(); i++ {
		//寻找[i,n]区间的最小值
		minIndex := i
		for j := i + 1; j < array.Len(); j++ {
			if array.Less(j, minIndex) {
				minIndex = j
			}
		}
		//交换
		array.Swap(i, minIndex)
	}
}

