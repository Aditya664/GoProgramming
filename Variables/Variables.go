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
	//Multiple variable decleration
	car, cost := "car", 6666
	fmt.Println(car, cost)
	car1, cost1 := "carff", 7777
	//we need to user variables once it will decleared
	//we can ignore by using _ (blank identifier)
	_, _ = car1, cost1

}
