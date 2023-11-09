package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, n := range []int{1, 2, 3} {
			a <- n
		}
		close(a)
	}()

	go func() {
		for _, n := range []int{20, 10, 30} {
			b <- n
		}
		close(b)
	}()

	go func() {
		for _, n := range []int{300, 200, 100} {
			c <- n
		}
		close(c)
	}()

	for n := range joinChannels(a, b, c) {
		fmt.Println(n)
	}
}

func joinChannels(chs ...<-chan int) <-chan int {
	mergeCh := make(chan int)

	go func() {
		wg := new(sync.WaitGroup)

		defer func() {
			wg.Wait()
			close(mergeCh)
		}()

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for n := range ch {
					mergeCh <- n
				}
			}(ch, wg)
		}
	}()

	return mergeCh
}
