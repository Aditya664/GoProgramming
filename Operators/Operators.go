package main

import "fmt"

func main() {
	//What is an Operator?
	/*An Operator is a symboll of the programming language whichis able to
	operate on values.in GO languange Operators can
	be categorized based upon their different functionallity
	is these category:
		* Arithmatic operator and bitwise operator: +,-,*,/,&,|,^,<,>.
		* Assignment Operator: +=,-=,*=,/=,%=.
		* Increament and decrement: ++,--.
		* Comparison operator: ==,!=,<=,>=,<,>.
		* Logical Operator: &&,||.
		* Operators for Pointer(&) and channels(<-).
	# Arithmatic operator
	it is used when we need to perform some common maths operations.
	there are following arithmatic operator:
		+(addition)
		-(substration)
		*(multiplication)
		/(Division)
		%(Modulus)
	*/
	//code
	var a int = 55
	var b int = 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
