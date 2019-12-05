package main

import (
	"fmt"
	"strconv"
)

var (
	min = 359282
	max = 820401
)

func isEqualToOrHigher(one, two int) bool {
	if one <= two {
		return true
	}
	return false
}

func hasConsecutive(in []int) bool {
	i := 0
	for i < len(in) {
		fmt.Printf("i: %+v\n", i)
		if i+1 != len(in) {
			if in[i] == in[i+1] {
				return true
			}
		}
		i++
	}
	return false
}

func iterator(in []int) bool {
	i := 0
	for i < len(in) {
		fmt.Printf("i: %+v\n", i)
		if i+1 != len(in) {
			if !isEqualToOrHigher(in[i], in[i+1]) {
				return false
			}
		}
		i++
	}
	return true
}

func main() {
	matches := []int{}
	for min <= max {
		fmt.Printf("%+v\n", min)
		s := strconv.Itoa(min)
		digits := []int{}
		for _, r := range s {
			digits = append(digits, int(r-'0'))
		}
		if iterator(digits) {
			fmt.Println("%i matches!\n", min)

			if hasConsecutive(digits) {
				matches = append(matches, min)
			}
		}
		min++
	}

	fmt.Printf("total matches: %+v", len(matches))

}
