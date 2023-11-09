package main

import "log"

func main() {
	s1 := []int{5, 3, 6, 7, 9}
	s2 := []int{3, 2, 5, 8, 7}

	r := intersection(s1, s2)
	log.Println(r)

	s1 = []int{1, 1, 1}
	s2 = []int{1, 1, 1, 1, 1}

	r = intersection(s1, s2)
	log.Println(r)
}

func intersection(s1, s2 []int) []int {
	counter := make(map[int]int)

	for _, v := range s1 {
		_, exist := counter[v]
		if exist {
			counter[v] += 1
			continue
		}

		counter[v] = 1
	}

	var result []int
	for _, v := range s2 {
		_, exist := counter[v]
		if exist && counter[v] > 0 {
			result = append(result, v)
			counter[v] -= 1
			continue
		}
	}

	return result
}
