package main

import "fmt"

func main() {
	//In golang we use the term of constant to represent fixed values.
	//constant created in compile time
	//How to declear constant?
	const days int = 7
	//we can declear constant without using it
	var i int
	fmt.Println(i)

	const pi float64 = 3.14

	//x, y := 5, 0
	//fmt.Println(x / y)
	// we cannot get error at this time because variable are created at Runtime
	//const a int= 5
	//const b int = 0
	//fmt.Println(a / b)
}
