package day1

import (
	"errors"
	"fmt"
	"strconv"
)

func StartPart1() string {
	input := parseInput()

	a, b, err := find2SumPair(input)
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	return strconv.Itoa(a * b)
}

func find2SumPair(input []int) (int, int, error) {
	for _, a := range input {
		for _, b := range input {
			if (a + b) == 2020 {
				return a, b, nil
			}
		}
	}
	return 0, 0, errors.New("not found")
}
