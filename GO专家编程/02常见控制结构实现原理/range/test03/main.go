package main

import "fmt"

func main() {
	ch:=make(chan int,10)
	go func() {
		for i:=0;i<10;i++{
			ch <- i
		}
		close(ch)
	}()
	for v := range ch{ // 通道关闭后会退出for range循环
		fmt.Println(v)
	}
}
