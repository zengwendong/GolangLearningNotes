package main

import "fmt"

// 规则一：延迟函数的参数在defer语句出现时就已经确定下来了
// 注意：对于指针类型参数，规则仍然适用，只不过延迟函数的参数是一个地址值，这种情况下，
// defer后面的语句对变量的修改可能会影响延迟函数

func deferFuncParameter() {
	var aInt = 1

	// 延迟函数fmt.Println(aInt)的参数在defer语句出现时就已经确定了，是值拷贝
	// 所以无论后面如何修改aInt变量都不会影响延迟函数
	defer fmt.Println(aInt) // 1

	aInt = 2
	return
}

func main()  {
	deferFuncParameter()
}