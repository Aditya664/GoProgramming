package main

import "fmt"

func main() {
	//we can declear variable without mension datatype
	var x = 8
	var y = "sdf"
	fmt.Println(x)
	fmt.Println(y)
	//we cannot assign one type of variable as it firstname
	//we need to convert it into that type of datatype
	var a = 12
	var b = 22.3
	//now  assign a to b it gives me error error
	//Cannot use 'b' (type float64) as the type int
	//in this case we need to cast float to int
	a = int(b)
	b = float64(a)
	fmt.Println(a)
	fmt.Println(b)

	//Literal means the value itself
	//Zero values
	//Zero value means default values of datatype
	//Go zero values
	//-numeric types: 0
	//--bool type: false
	//-string type: ""(Empty string)
	//-Pointer type: nil
	var value int
	var prince float64
	var name string
	var done bool
	fmt.Println(value, prince, name, done)
	//It Gives := 0 0  false ... as output that means
	//value of int and float is 0
	//and value of string is  ""
	//and value of bool is false

}
