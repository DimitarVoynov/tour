// +build OMIT

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 //short var declaration is only available inside of a function
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
