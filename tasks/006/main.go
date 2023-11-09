package main

import "log"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	go func() {
		defer close(natural)

		for i := 0; i <= 10; i++ {
			natural <- i
		}
	}()

	go func() {
		defer close(squares)

		for n := range natural {
			squares <- n * n
		}
	}()

	for v := range squares {
		log.Println(v)
	}
}
