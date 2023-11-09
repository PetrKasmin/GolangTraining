package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	r := randomToNumber(5)
	log.Println(r)

	r = randomToNumber(10)
	log.Println(r)
}

func randomToNumber(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < n; i++ {
			ch <- r.Intn(n)
		}
	}()

	return <-ch
}
