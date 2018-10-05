// +build OMIT

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
//Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) //this is fine but the below is also fine (as opposed to functions! see method-pointers-explained.go line 26)
	// p := &v
	// p.Scale(10)
	fmt.Println(v.Abs())
}
