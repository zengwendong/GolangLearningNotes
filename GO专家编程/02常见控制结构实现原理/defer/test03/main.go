package main

import "fmt"

func deferFuncReturn() (result int) {
	i := 1

	defer func() {
		result++
	}()

	return i
}

func main() {
	result := deferFuncReturn()
	fmt.Printf("result:%d\n", result) // result:2
}

/*
关键字return不是一个原子操作，实际上return只代理汇编指令ret，即将跳转程序执行。
比如语句return i，实际上分两步进行，即将i值存入栈中作为返回值，然后执行跳转，
而defer的执行时机正是跳转前，所以说defer执行时还是有机会操作返回值的。
 */