package main

import "fmt"

func main() {
	/*A type determines a set of values together with operations and methods specific to those values
	there are three types:
		1.predeclared types
		2.introduced types
		3.comosite types
	Predeclared/Build in types
	* Numeric types
		-- int8,int16,int32,int64:Used to represent signed integer
		-- uint8,uint16,uint32,uint64:Used to represent unsigned (positive) integer.
		-- float32,float64: zero before the decimal point seperator can be omitted
		-- byte(alias for uint8)
	* Bool Type
		-- Pre-defined constants true and false.
	* String Type
		-- Unicode chars written enclosed by double-quotes
		-- A String value is a sequence of bytes.
	Comosite types
	*Array and Slice Type
		-- Array is collection of elements of same type.
		-- Array has fixed length but slice is dynamuic length*/

	//code
	//int8
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)
	//int16
	//var i2 int16 = 65539
	//fmt.Printf("%T\n", i2)

	//float64
	var f1, f2, f3 float64 = 1.1, -20.0, 5.0
	fmt.Printf("%T %T %T", f1, f2, f3)

}
