package main

import (
	"fmt"
)

//The return type comes after the function parameters
//func add(x, y int) int, since both variables are integers
func add(x int, y int) int {
	return x + y
}

//A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

//return values can be named as seen in the case of x and y
// When a return statement has no arguments, it is known as "naked". This should only be used in short functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//var is used to declare variables, and can be used at the package level or the function level.
//The type of the variable always goes last
//If a variable is declared without an explicit initial value, then they are assigned their "zero" value
//var c, python, java bool

//If you want to initialize, then you can do one value per variable
var c, python, java = true, false, "no!"

func main() {
	//Exported names always begin with a capital letter
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))

	//var i int
	var i, j int = 1, 2
	fmt.Println(i, j, c, python, java)

	//:= is known as a "short assignment" and can be used inside a function in place of a var declaration
	k := 3
	fmt.Println(k)

	//Typicall conversion types
	l := 42
	f := float64(l)
	u := uint(f)
	fmt.Println(l, f, u)

	//Constants are declared just like variables, except the short assignment cannot be used
	//They can also be declared on a package or function level
	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")
}
