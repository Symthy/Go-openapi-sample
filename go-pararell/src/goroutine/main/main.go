package main

import (
	"fmt"
	"time"
)

func output(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {

	output("main thread")

	go output("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("immediate execution")

	time.Sleep(time.Second)
	fmt.Println("done")
}
