package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Go does not have classes. However, you can define methods on types.
//A method is a function with a special receiver argument between "func" and the method name

//Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
