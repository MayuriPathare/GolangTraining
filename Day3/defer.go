package main

import "fmt"

func main() {
	func1()
}

func func1() {
	var a = "Hello"
	for i := 0; i < 10; i++ {
		defer func() {
			if a == "Hello"{
				fmt.Println("Hi")
			}
		}()
		fmt.Println(i)
		if i == 5 {
			continue
		}
	}
}