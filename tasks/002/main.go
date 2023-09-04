package main

import (
	"fmt"
)

func run(s []int, ch chan int) {
	defer func() {
		close(ch)
	}()

	for _, n := range s {
		ch <- n
	}
}

func main() {
	var r []int
	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6, 8, 10}

	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go run(s1, ch1)
	go run(s2, ch2)
	go handler(ch, ch1, ch2)
	for n := range ch {
		r = append(r, n)
	}

	fmt.Println(r)
}

func handler(ch, ch1, ch2 chan int) {
	defer close(ch)

	for {
		n1, ok1 := <-ch1
		if ok1 {
			ch <- n1
		}

		n2, ok2 := <-ch2
		if ok2 {
			ch <- n2
		}

		if !ok1 && !ok2 {
			break
		}
	}
}
