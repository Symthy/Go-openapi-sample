package main

import (
	"fmt"
	"strconv"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("worker " + strconv.Itoa(id) + " start")
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const jobNum = 5
	jobs := make(chan int)
	results := make(chan int, jobNum)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//go func() {
	for j := 1; j <= jobNum; j++ {
		jobs <- j
	}
	//}()
	close(jobs)
	fmt.Println("sleep...")
	time.Sleep(2 * time.Second)
	fmt.Println("sleep end")

	for a := 1; a <= jobNum; a++ {
		fmt.Println(<-results)
	}
}
