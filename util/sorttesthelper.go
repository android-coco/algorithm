package util

import (
	"math/rand"
	"time"
	"fmt"
	"strconv"
)

//面向接口编程  实现java泛型
type Sortable interface {
	Len() int           //长度
	Less(int, int) bool //最小值
	Swap(int, int)      //交换
}
type Student struct {
	Name  string //名称
	Score int    //分数
}

func (stu Student) toString() string {
	return "[ 姓名：" + stu.Name + " 分数：" + strconv.Itoa(stu.Score) + " ]"
}

type StuArray []Student

func (arrStu StuArray) Len() int {
	return len(arrStu)
}
func (arrStu StuArray) Less(i, j int) bool {
	//分数相同 按照名字先后先后顺序排
	if arrStu[i].Score != arrStu[j].Score {
		return arrStu[i].Score < arrStu[j].Score
	}
	return arrStu[i].Name < arrStu[j].Name
}
func (arrStu StuArray) Swap(i, j int) {
	arrStu[i], arrStu[j] = arrStu[j], arrStu[i]
}

type IntArray []int

func (arrInt IntArray) Len() int {
	return len(arrInt)
}

func (arrInt IntArray) Less(i, j int) bool {
	return arrInt[i] < arrInt[j]
}

func (arrInt IntArray) Swap(i, j int) {
	arrInt[i], arrInt[j] = arrInt[j], arrInt[i]
}

type Float64Array []float64

func (arrFloat64 Float64Array) Len() int {
	return len(arrFloat64)
}

func (arrFloat64 Float64Array) Less(i, j int) bool {
	return arrFloat64[i] < arrFloat64[j]
}

func (arrFloat64 Float64Array) Swap(i, j int) {
	arrFloat64[i], arrFloat64[j] = arrFloat64[j], arrFloat64[i]
}

type StringArray []string

func (arrStr StringArray) Len() int {
	return len(arrStr)
}

func (arrStr StringArray) Less(i, j int) bool {
	return len(arrStr[i]) < len(arrStr[j])
}

func (arrStr StringArray) Swap(i, j int) {
	arrStr[i], arrStr[j] = arrStr[j], arrStr[i]
}

//生成随机数 切片
func SortTestHelper(n, rangeL, rangeR int) IntArray {
	arrInt := make([]int, 0)
	for i := 0; i < n; i++ {
		arrInt = append(arrInt, randInt(rangeL, rangeR))
	}
	return IntArray(arrInt)
}

//随机数
func randInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(max-min) + min
}

//测试排序运行时间
func TestSort(sortName string, method func(array Sortable), arr Sortable) {
	startTime := time.Now()
	method(arr)
	if isSort(arr) {
		fmt.Println(sortName+" 运行时间：", time.Now().Sub(startTime))
	} else {
		fmt.Println(sortName+" 运行时间：", time.Now().Sub(startTime), "但是排序错误")
	}

}

//打印输出
func PrintArray(array Sortable) {
	switch array.(type) {
	case IntArray:
		intArray, ok := array.(IntArray)
		checkConvert(ok)
		for _, v := range intArray {
			fmt.Println(v)
		}
	case Float64Array:
		float64Array, ok := array.(Float64Array)
		checkConvert(ok)
		for _, v := range float64Array {
			fmt.Println(v)
		}
	case StringArray:
		stringArray, ok := array.(StringArray)
		checkConvert(ok)
		for _, v := range stringArray {
			fmt.Println(v)
		}
	case StuArray:
		stuArray, ok := array.(StuArray)
		checkConvert(ok)
		for _, v := range stuArray {
			fmt.Println(v.toString())
		}
	default:
		fmt.Println("未知类型")
	}
}

//判断是否有序
func isSort(array Sortable) bool {
	switch array.(type) {
	case IntArray:
		intArray, ok := array.(IntArray)
		checkConvert(ok)
		for i := 0; i < intArray.Len()-1; i++ {
			if intArray[i] > intArray[i+1] {
				return false
			}
		}
		return true
	case Float64Array:
	case StringArray:
	case StuArray:
	default:
		fmt.Println("未知类型")
	}
	return false
}

func checkConvert(ok bool) {
	if !ok {
		panic("type cast failed!")
	}
}


