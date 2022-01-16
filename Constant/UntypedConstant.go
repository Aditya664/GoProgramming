package main

import "fmt"

func main() {
	const a float64 = 5.1 //Typed constant
	const b = 77.32       //Untyped Constant
	const c float64 = a * b
	const str = "Hello" + "Go!"
	const d = 5 > 100
	fmt.Println(d)

	const x int = 5
	//const y float64 = 2.2 * x

	const x1 = 5
	const y = 2.2 * x1
	fmt.Println(y)
	//it is possible because both are untyped
}
