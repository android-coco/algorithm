package main

import (
	"algorithm/util"
	"fmt"
)

func main() {
	n := 100000
	arr := util.SortTestHelper(n, 0, n)
	arr1 := util.SortTestHelper(n, 0, n)
	util.TestSort("insertionSort2", insertionSort2, arr1)
	util.TestSort("insertionSort", insertionSort, arr)

}

//插入排序(近乎有序的数组排序  很快)
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

//插入排序升级版
func insertionSort2(sortable util.Sortable) {
	switch sortable.(type) {
	case util.IntArray:
		intArray, ok := sortable.(util.IntArray)
		util.CheckConvert(ok)
		len := sortable.Len()
		for i := 1; i < len; i++ {
			//寻找元素arr[i]合适的插入位置
			e := intArray[i]
			var j = 0 //保存元素e应该插入的位置
			for j = i; j > 0 && intArray[j-1] > e; j-- {
				//比前面的小  向后挪一位
				intArray[j] = intArray[j-1]
			}
			intArray[j] = e //到了 赋值
		}
	case util.Float64Array:
		float64Array, ok := sortable.(util.Float64Array)
		util.CheckConvert(ok)
		len := sortable.Len()
		for i := 1; i < len; i++ {
			//寻找元素arr[i]合适的插入位置
			e := float64Array[i]
			var j = 0 //保存元素e应该插入的位置
			for j = i; j > 0 && float64Array[j-1] > e; j-- {
				//比前面的小  向后挪一位
				float64Array[j] = float64Array[j-1]
			}
			float64Array[j] = e //到了 赋值
		}
	case util.StringArray:
		stringArray, ok := sortable.(util.StringArray)
		util.CheckConvert(ok)
		len := sortable.Len()
		for i := 1; i < len; i++ {
			//寻找元素arr[i]合适的插入位置
			e := stringArray[i]
			var j = 0 //保存元素e应该插入的位置
			for j = i; j > 0 && stringArray[j-1] > e; j-- {
				//比前面的小  向后挪一位
				stringArray[j] = stringArray[j-1]
			}
			stringArray[j] = e //到了 赋值
		}
	case util.StuArray:
		stuArray, ok := sortable.(util.StuArray)
		util.CheckConvert(ok)
		len := sortable.Len()
		for i := 1; i < len; i++ {
			//寻找元素arr[i]合适的插入位置
			e := stuArray[i]
			var j = 0 //保存元素e应该插入的位置
			for j = i; j > 0 && stuArray[j-1].Score > e.Score; j-- {
				//比前面的小  向后挪一位
				stuArray[j] = stuArray[j-1]
			}
			stuArray[j] = e //到了 赋值
		}
	default:
		fmt.Println("未知类型")
	}
}
