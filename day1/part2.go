package day1

import (
	"errors"
	"fmt"
	"strconv"
)

func StartPart2() string {
	input := parseInput()

	a, b, c, err := find3SumPair(input)
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	return strconv.Itoa(a * b * c)
}

func find3SumPair(input []int) (int, int, int, error) {
	for _, a := range input {
		for _, b := range input {
			for _, c := range input {
				if (a + b + c) == 2020 {
					return a, b, c, nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("not found")
}
