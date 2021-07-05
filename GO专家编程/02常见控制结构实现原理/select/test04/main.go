package main

func main() {

	//for i:=0;i<10;i++{
	//	fmt.Println(i)
	//}

	//go func() {
	//	for i:=0;i<10;i++{
	//		fmt.Println(i)
	//	}
	//}()

	// fatal error: all goroutines are asleep - deadlock!
	select {

	}
}

/*
对于空的select语句，程序会被阻塞，准确的说是当前协程被阻塞，
同时Golang自带死锁检测机制，当发现当前协程再也没有机会被唤醒时，则会panic
 */