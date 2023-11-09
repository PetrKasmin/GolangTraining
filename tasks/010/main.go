package main

import "log"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}

	isEqual := Equal(a, b)

	log.Println(isEqual)
}
