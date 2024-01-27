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

	var myName = "Craig"
	var name string = "Timmy"
	userName := "admin"

	var sum int
	part1, other := 1, 5
	part2, other := 2, 0

	sum = part1 + part2

	var ( // --> Block assignment
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println("LessonName=", lessonName)
	fmt.Println("LessonType=", lessonType)

	word1, word2, _ := "hello", "world", "!" // --> compound assignment
	fmt.Println(word1, word2)
	fmt.Println("part 2 is", part2, "other is", other)
	fmt.Println("part 1 is", part1, "other is", other)
	fmt.Println("The sum is", sum)

	fmt.Println("userName = ", userName)
	fmt.Println("name = ", name)
	fmt.Println("my name is ", myName)

	fmt.Println(MaxSpeed, MinPurchasePrice, AppAuthor)
	fmt.Println(y, z)
	fmt.Println(test)

	fmt.Println(example, example2, example3)
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)

}
