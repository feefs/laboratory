package main

import (
	"fmt"
	"time"
)

func nums() <-chan int {
	c := make(chan int)
	n := 0
	go func() {
		for {
			c <- n
			n += 1
		}
	}()
	return c
}

func fill(freq time.Duration, delay time.Duration, src <-chan int, dest chan<- int) {
	go func() {
		time.Sleep(delay)
		for {
			dest <- <-src
			time.Sleep(freq)
		}
	}()
}

func tick(freq time.Duration, delay time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		time.Sleep(delay)
		for {
			c <- struct{}{}
			time.Sleep(freq)
		}
	}()
	return c
}

func main() {
	messages := make(chan int)
	fill(333*time.Millisecond, 1*time.Second, nums(), messages)
	tick := tick(1*time.Second, 3*time.Second)

	fmt.Println("Starting program!")
	batch := []int{}
	for i := 0; i < 25; i++ {
		select {
		case message := <-messages:
			batch = append(batch, message)
			fmt.Printf("Added %v to messages\n", message)
		case <-tick:
			fmt.Printf("Batch: %v\n", batch)
			batch = []int{}
		}
	}
}
