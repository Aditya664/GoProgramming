package main

//vriables is name for memeory location where values is stored
//A decleared variable must be used or we get error
//_ is blank identifier and mutes the compile time error returned by unused variables.

import "fmt"

func main() {
	//using var keyword
	var x1 int = 7
	var y1 string = "Kp"
	fmt.Println(x1)
	fmt.Println(y1)
	//Using the short Declaration Operator(:=)
	x2 := 6
	y2 := "kp"
	fmt.Println(x2)
	fmt.Println(y2)
}
