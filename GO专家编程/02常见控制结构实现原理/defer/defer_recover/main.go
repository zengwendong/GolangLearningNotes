package main

import "fmt"

func NoPanic() {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
	}
}

func Dived(n int) {
	defer NoPanic()

	fmt.Println(1/n)
}

func main() {
	Dived(0)
}
