package main

import (
	"fmt"
)

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

	const (
		min1 = -222
		min2 = 3333
		min3 = 4544
	)
	fmt.Println(min1, min2, min3)
	//we get -222 3333 4544 this output
	const (
		min4 = -222
		min5
		min6
	)
	fmt.Println(min4, min5, min6)
	// in this case we get -222 -222 -222 but How
	//min5 and min6 get type value from priveous value

	//Rules in constant
	//1.we cannot change the value of constant
	const temp int = 100
	//temp = 55
	//2.we cannot initiate a constant at Runtime
	//const power  = math.Pow(2,3)
	//it gives error because math funcyion runs at runtime.
	//we know we cannot initiate constant at runtime
	//3.we cannot use variable to initialize constant.
	//t := 55
	//const tc = t

}
