// +build OMIT

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens in the new goroutine.

	// Goroutines run in the same address space, so access to shared memory must be synchronized
	go say("world")
	say("hello")
}
