package main

import "fmt"

func main() {
	// showing the different ways to declare a basic variable
	var example = 3

	var example2 int = 3

	var example3 int // <-- uninitialized variable
	example3 = 3     // <-- now it is intitialized with a value

	var a, b, c = 1, 2, "sample"

	var (
		d int = 3
		e int = 4
		f     = "sample"
	)

	// create and assign " := "
	test := 3 // creates the variable and assigns the value, a shorthand for var test int = 3
	y, z := 10, "sampleTest"

	const MaxSpeed = 30
	const MinPurchasePrice = 7.50
	const AppAuthor = "Bob"

	fmt.Println(MaxSpeed, MinPurchasePrice, AppAuthor)
	fmt.Println(y, z)
	fmt.Println(test)

	fmt.Println(example, example2, example3)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)

}
