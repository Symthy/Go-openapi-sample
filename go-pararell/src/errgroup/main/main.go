package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(id int) {
	fmt.Printf("Worker %d start\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d end\n", id)
}

func main() {
	var eg errgroup.Group
	for i := 1; i <= 5; i++ {
		id := i
		eg.Go(func() error {
			worker(id)
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {
		fmt.Println("error: ", err)
	}
}
