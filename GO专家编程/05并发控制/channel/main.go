package main

import (
	"fmt"
	"time"
)

func Process(ch chan int) {
	//Do some work...
	time.Sleep(time.Second)

	ch <- 1 //管道中写入一个元素表示当前协程已结束
}

func main() {
	channels := make([]chan int, 10) //创建一个10个元素的切片，元素类型为channel

	for i:= 0; i < 10; i++ {
		channels[i] = make(chan int) //切片中放入一个channel
		go Process(channels[i])      //启动协程，传一个管道用于通信
	}

	for i, ch := range channels {  //遍历切片，等待子协程结束
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}

/*
Routine  0  quit!
Routine  1  quit!
Routine  2  quit!
Routine  3  quit!
Routine  4  quit!
Routine  5  quit!
Routine  6  quit!
Routine  7  quit!
Routine  8  quit!
Routine  9  quit!
 */

// 使用channel来控制子协程的优点是实现简单，缺点是当需要大量创建协程时就需要有相同数量的channel，
// 而且对于子协程继续派生出来的协程不方便控制