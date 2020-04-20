package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("This is first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	func(n string) {

        fmt.Println("A string argument is passed to the anonymous Welcome", n)

    }("Go")


	fmt.Println("This is end of main function")
}

// Use Cases:
// 1. Can be used as decorators i.e passing function as a argument to other function
// 2. It can be use like lambda function in python