package main

import (
	"log"
)

func main() {
	words := []string{"abcde", "abcda"}

	for _, word := range words {
		count := make(map[rune]string)

		for _, char := range word {
			_, ok := count[char]
			if ok {
				continue
			}

			count[char] = string(char)
		}

		log.Println(len(count) == len([]rune(word)))
	}
}
