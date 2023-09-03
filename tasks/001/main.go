package main

import (
	"fmt"
	"sync"
)

func run(s []int, ch chan<- int, wg *sync.WaitGroup) {
	defer close(ch)

	for _, n := range s {
		wg.Wait()
		wg.Add(1)
		ch <- n
	}
}

func handler(ch chan<- int, ch1, ch2 <-chan int, wg1, wg2 *sync.WaitGroup) {
	defer close(ch)

	n1, ok1 := <-ch1
	n2, ok2 := <-ch2

	for ok1 || ok2 {
		if !ok1 {
			ch <- n2
			wg2.Done()
			n2, ok2 = <-ch2
		} else if !ok2 {
			ch <- n1
			wg1.Done()
			n1, ok1 = <-ch1
		} else if n1 < n2 {
			ch <- n1
			wg1.Done()
			n1, ok1 = <-ch1
		} else {
			ch <- n2
			wg2.Done()
			n2, ok2 = <-ch2
		}
	}
}

func main() {
	var (
		wg1, wg2 sync.WaitGroup
		result   []int
	)

	s1 := []int{1, 3, 5, 7, 9}
	s2 := []int{2, 4, 6, 8, 10}

	ch := make(chan int)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go run(s1, ch1, &wg1)
	go run(s2, ch2, &wg2)

	go handler(ch, ch1, ch2, &wg1, &wg2)

	for n := range ch {
		result = append(result, n)
	}

	fmt.Println(result)
}
