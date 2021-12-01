package main

import (
	"fmt"
	"github.com/Tommylans/Advent2020/day1"
	"github.com/Tommylans/Advent2020/day2"
	"time"
)

func main() {
	start := time.Now()
	day1Part1 := day1.StartPart1()
	fmt.Println("Day1:p1 Answer", day1Part1)
	fmt.Println("Day1:p1", time.Since(start))

	start = time.Now()
	day1Part2 := day1.StartPart2()
	fmt.Println("Day1:p2 Answer", day1Part2)
	fmt.Println("Day1:p2", time.Since(start))

	start = time.Now()
	day2Part1 := day2.StartPart1()
	fmt.Println("Day2:p1 Answer", day2Part1)
	fmt.Println("Day2:p1", time.Since(start))

	start = time.Now()
	day2Part2 := day2.StartPart2()
	fmt.Println("Day2:p2 Answer", day2Part2)
	fmt.Println("Day2:p2", time.Since(start))
}
