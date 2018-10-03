// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //This construct can be a clean way to write long if-then-else chains.
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
