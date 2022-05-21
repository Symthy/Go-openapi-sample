package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "sending 1"
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	} // timeout

	ch2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "sending 2"
	}()
	select {
	case res := <-ch2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	} // success
}
