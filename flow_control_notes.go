package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {

	//Standard declaration of an "if" statement
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	//Include a short initialization statement before the conditional statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	//Defers the execution of a function until the surrounding function returns
	//If there are multiple, the defers are pushed onto a stack and executed in LIFO order
	defer fmt.Println("world")

	sum := 0

	//Typical declaration of a for loop in go
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	//The init and post statements are optional
	//Essentially this is a "while" loop in C
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	//Standard implementation of a switch statement in golang
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
