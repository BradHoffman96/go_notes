package main

import "fmt"
import "math"

//Typical implementation of a struct
type Vertex struct {
	X int
	Y int
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//functions can be taken as parameters and used as return values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	//& references the address value for a variable, * dereferences a pointer and shows the actual value for that address
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	//Standard use of a struct
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//declares
	// var a [10]int

	//make() is how you create dynamically-sized arrays
	//make() will allocated a zeroed array and returns a slice that refers to that array
	//Specify a capacity by passing a third argument to make, e.g. make([]int, 0, 5)
	// b := make([]int, 5)

	//slices can contain other slices. This is an example of a tic-tac-toe board
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	//Changing the elements of a slice modifies the corresponding elements of its underlying array
	//A slice can have both a length and a capacity. The length of the slice is the number of elements it contains.
	//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	//Use len() and cap() to retrieve these values.
	//You can extend a slice's length by re-slicing it.

	var s []int
	printSlice(s)

	//append() is used to append values to a slice. The first param is the slice, and the remaining are the values that the user wants to add to the slice
	//A larger array will be allocated if the backing array of the slice is too small
	s = append(s, 0)
	printSlice(s)

	//range can be used in a for loop to iterate through the elements in a slice or map
	//the first variable is the index, and the second is the value
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	//Zero value of a map is nil. It has no keys and keys cannot be added
	//make() will return a map initialized and ready to use
	//delete(map, key) is used to delete a key, if it is present
	//elem, ok = m[key] to test if a key is present
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40, 74,
	}
	fmt.Println(m["Bell Labs"])

	//An example of a map literal
	// var h = map[string]Vertex{
	// 	"Bell Labs": {40, -74},
	// 	"Google":    {37, -122},
	// }

	//Functions can be created in-line and even passed around to other functions
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}
