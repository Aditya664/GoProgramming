package IOTA

func main() {
	//What is IOTA?
	//iota is identifier which is used with constant and which can simplify constant defination
	//that use auto increament numbers.the iota keyword represent integer constant starting from zero.

	//Auto increament constant without IOTA
	const (
		a = 0
		b = 1
		c = 3
	)
	//Auto increament constant with IOTA
	const (
		d = iota
		e
		f
	)
	//both will set
	//a = 0
	//b = 1
	//c = 2
	//and
	//d = 0
	//e = 1
	//f = 2
}
