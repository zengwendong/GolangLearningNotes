package main

import "fmt"

func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}

func UpdateTable(n int) {
	// defer中决定提交还是回滚
	defer func() {
		if IsPanic() {
			fmt.Println("Rollback transaction")
		} else {
			fmt.Println("Commit transaction")
		}
	}()

	// 模拟 Database update operation...
	fmt.Println(100 / n)
}


func main() {
	UpdateTable(0) // 发生panic: runtime error: integer divide by zero 但依旧打印"Commit transaction"
}

/*
recover没有被defer方法直接调用时会失效
例子中recover() 调用栈为“defer （匿名）函数” –> IsPanic() –> recover()
 */