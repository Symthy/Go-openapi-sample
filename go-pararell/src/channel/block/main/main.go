package main

import "fmt"

func main() {
	// var writeCh chan<- interface{}
	var readCh <-chan interface{}
	ch := make(chan interface{})
	// writeCh = ch
	readCh = ch

	go func() {
		// writeCh <- "Writing..."
	}()

	fmt.Println(<-readCh)
}
