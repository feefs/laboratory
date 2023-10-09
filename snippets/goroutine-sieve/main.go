package main

import "fmt"

func gen() (src <-chan int) {
	c := make(chan int)
	go func() {
		n := 2
		for {
			c <- n
			n += 1
		}
	}()
	return (<-chan int)(c) // Cast can be excluded if desired
}

func sieve(src <-chan int, prime int) (new_src <-chan int) {
	output := make(chan int)
	go func() {
		for {
			n := <-src
			if n%prime != 0 {
				output <- n
			}
		}
	}()
	return (<-chan int)(output)
}

func main() {
	src := gen()
	for i := 0; i < 50; i++ {
		prime := <-src
		fmt.Printf("Prime sieved: %v\n", prime)
		src = sieve(src, prime)
	}
}
