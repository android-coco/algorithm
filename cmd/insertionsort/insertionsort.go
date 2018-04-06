package main

import (
	"algorithm/util"
)

func main() {
	n := 10000
	arr := util.SortTestHelper(n, 0, n)
	util.TestSort("insertionSort", insertionSort, arr)
}

//插入排序
func insertionSort(sortable util.Sortable) {
	len := sortable.Len()
	for i := 1; i < len; i++ {
		//寻找元素arr[i]合适的插入位置
		// 从当前位置向前比较 当前位置为1时并且前面的一个元素比当前元素小时停止
		for j := i; j > 0 && sortable.Less(j, j-1); j-- {
			sortable.Swap(j, j-1) //交换位置
		}
	}
}
