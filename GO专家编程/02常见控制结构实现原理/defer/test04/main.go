package main

import "fmt"

func foo() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func foo2() *int {
	var i int
	defer func() {
		i++
	}()
	return &i
}

func foo3() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

func main() {
	fmt.Println(foo())   // 0
	fmt.Println(*foo2()) // 1
	fmt.Println(foo3())  // 1
}
