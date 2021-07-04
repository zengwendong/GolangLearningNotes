package main

import "fmt"

func printArray(array *[3]int) {
	for i := range array {
		fmt.Println(array[i])
	}
}

func deferFuncParameter() {
	var aArray = [3]int{1, 2, 3}
	defer printArray(&aArray)
	aArray[0] = 10
	return
}

func main() {
	deferFuncParameter()
}

/*
10
2
3
 */