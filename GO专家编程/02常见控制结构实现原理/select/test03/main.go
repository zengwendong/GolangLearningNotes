package main

import (
	"fmt"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		close(chan1)
	}()

	go func() {
		close(chan2)
	}()

	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	}

	fmt.Println("main exit.")
}

/*
select会按照随机的顺序检测各case语句中channel是否ready，考虑到已关闭的channel也是可读的，默认为数据类型的零值
所以上述程序中select不会阻塞，具体执行哪个case语句具是随机的

对于读channel的case来说，如case elem, ok := <-chan1:,
如果channel有可能被其他协程关闭的情况下，一定要检测读取是否成功，因为close的channel也有可能返回
 */