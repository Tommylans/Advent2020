package day1

import (
	"bufio"
	"os"
	"strconv"
)

func parseInput() (out []int) {
	file, _ := os.Open("./day1/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		out = append(out, number)
	}
	return
}
