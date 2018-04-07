package main

import (
	"algorithm/util"
	"fmt"
)
//https://www.cnblogs.com/chengxiao/p/6104371.html
func main() {
	n := 100000
	arr := util.SortTestHelper(n, 0, n)
	util.TestSort("ShellSort1", ShellSort1, arr)
	util.TestSort("ShellSort2", ShellSort2, arr)
}

/**
希尔排序 针对有序序列在插入时采用交换法
 */
func ShellSort1(sortable util.Sortable) {
	//增量gap，并逐步缩小增量
	for gap := sortable.Len() / 2; gap > 0; gap /= 2 {
		//从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < sortable.Len(); i++ {
			j := i
			for j-gap >= 0 && sortable.Less(j, j-gap) {
				//插入排序采用交换法
				sortable.Swap(j, j-gap)
				j -= gap
			}
		}
	}
}

/**
希尔排序 针对有序序列在插入时采用移动法。
 */
func ShellSort2(sortable util.Sortable) {

	switch sortable.(type) {
	case util.IntArray:
		intArray, ok := sortable.(util.IntArray)
		util.CheckConvert(ok)
		len := sortable.Len()
		//增量gap，并逐步缩小增量
		for gap := len / 2; gap > 0; gap /= 2 {
			//从第gap个元素，逐个对其所在组进行直接插入排序操作
			for i := gap; i < len; i++ {
				j := i
				temp := intArray[j]
				for j-gap >= 0 && temp < intArray[j-gap] {
					//移动法
					intArray[j] = intArray[j-gap]
					j -= gap
				}
				intArray[j] = temp
			}
		}
	case util.Float64Array:
		float64Array, ok := sortable.(util.Float64Array)
		util.CheckConvert(ok)
		len := sortable.Len()
		//增量gap，并逐步缩小增量
		for gap := len / 2; gap > 0; gap /= 2 {
			//从第gap个元素，逐个对其所在组进行直接插入排序操作
			for i := gap; i < len; i++ {
				j := i
				temp := float64Array[j]
				for j-gap >= 0 && temp < float64Array[j-gap] {
					//移动法
					float64Array[j] = float64Array[j-gap]
					j -= gap
				}
				float64Array[j] = temp
			}
		}
	case util.StringArray:
		stringArray, ok := sortable.(util.StringArray)
		util.CheckConvert(ok)
		len := sortable.Len()
		//增量gap，并逐步缩小增量
		for gap := len / 2; gap > 0; gap /= 2 {
			//从第gap个元素，逐个对其所在组进行直接插入排序操作
			for i := gap; i < len; i++ {
				j := i
				temp := stringArray[j]
				for j-gap >= 0 && temp < stringArray[j-gap] {
					//移动法
					stringArray[j] = stringArray[j-gap]
					j -= gap
				}
				stringArray[j] = temp
			}
		}
	case util.StuArray:
		stuArray, ok := sortable.(util.StuArray)
		util.CheckConvert(ok)
		len := sortable.Len()

		//增量gap，并逐步缩小增量
		for gap := len / 2; gap > 0; gap /= 2 {
			//从第gap个元素，逐个对其所在组进行直接插入排序操作
			for i := gap; i < len; i++ {
				j := i
				temp := stuArray[j]
				for j-gap >= 0 && temp.Score < stuArray[j-gap].Score {
					//移动法
					stuArray[j] = stuArray[j-gap]
					j -= gap
				}
				stuArray[j] = temp
			}
		}


	default:
		fmt.Println("未知类型")
	}
}
