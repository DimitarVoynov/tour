// +build OMIT

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" //Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes!
	fmt.Println(a, b)
	fmt.Println(names)
}
