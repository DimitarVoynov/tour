// +build OMIT

package main

import "fmt"

//Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
