package main

import (
	"algorithm/util"
	"fmt"
)

func main() {
	n := 10000
	arr := util.SortTestHelper(n, 0, n)
	util.PrintArray(arr)
	fmt.Println("=============", arr.Len())
	util.TestSort("BubbleSort", BubbleSort, arr)
	//util.PrintArray(arr)
}

//冒泡排序
func BubbleSort(values util.Sortable) {
	flag := true
	vLen := values.Len() //长度
	for i := 0; i < vLen-1; i++ {
		flag = true
		for j := 0; j < vLen-i-1; j++ {
			if values.Less(j+1, j) {
				values.Swap(j, j+1)
				flag = false
				continue
			}
		}
		if flag {
			break
		}
	}
}
