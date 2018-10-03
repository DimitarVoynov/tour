// +build OMIT

package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 { //the same as for-continued.go (check it out)
		sum += sum
	}
	fmt.Println(sum)
}
