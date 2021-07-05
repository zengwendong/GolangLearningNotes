// 遍历性能上的优化

package main

// range遍历切片下标和值
func RangeSlice(slice []int) {
	for index, value := range slice {
		_, _ = index, value
	}
}

// 遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，
// 可以在for-range中忽略value值，使用slice[index]引用value值
func RangeSliceOptimize(slice []int) {
	for index, _ := range slice {
		_, _ = index, slice[index]
	}
}

// 函数中for-range语句中只获取key值，然后根据key值获取value值，虽然看似减少了一次赋值，
// 但通过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行
func RangeMap(myMap map[int]string) {
	for key, _ := range myMap {
		_, _ = key, myMap[key]
	}
}


func main() {
	
}
