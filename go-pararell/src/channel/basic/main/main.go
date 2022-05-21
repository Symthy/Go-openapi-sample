package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	msgs := make(chan string, 3)
	msgs <- "main"
	go func() {
		msgs <- "func1"
	}()
	go func() {
		msgs <- "func2"
	}()
	msg1 := <-msgs
	msg2 := <-msgs
	msg3 := <-msgs
	fmt.Println(msg1, msg2, msg3)

	// send slow
	ch1 := make(chan string, 2)
	go func() {
		for i := 0; i < 6; i++ {
			send := "send" + strconv.Itoa(i)
			ch1 <- send
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for j := 0; j < 3; j++ {
			recv := <-ch1
			fmt.Println("sub:", j, recv)
		}
	}()
	for j := 0; j < 3; j++ {
		recv := <-ch1
		fmt.Println("main:", j, recv)
	}
	time.Sleep(3 * time.Second)
}
