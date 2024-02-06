package main

import "fmt"

func sum(lhs, rhs int) int {
	return lhs + rhs
}

//	func multiReturn() (int, int, int) {
//		return 1, 2, 3
//	}
func twoNum() (int, int) {
	return 33, 44
}

func double(x int) int {
	return x + x
}

func greet() {
	fmt.Println("Hello from greet function")
}
func main() {

	// a, b, _ := multiReturn()

	// result := sum(45, 77)
	// fmt.Println(result)
	// fmt.Println(sum(12, 44))

	greet()
	dozen := double(6)
	fmt.Println(dozen)

	bakersDozen := sum(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	fmt.Println(twoNum())
}
