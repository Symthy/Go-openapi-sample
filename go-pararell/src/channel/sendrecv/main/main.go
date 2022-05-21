package main

import "fmt"

func send(recvCh chan<- string, msg string) {
	recvCh <- msg
}

func receive(sendCh <-chan string, recvCh chan<- string) {
	msg := <-sendCh
	recvCh <- msg
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// send(ch1, "sending")  deadlock..
	// receive(ch1, ch2)
	go send(ch1, "sending")
	go receive(ch1, ch2)
	fmt.Println(<-ch2)
}
