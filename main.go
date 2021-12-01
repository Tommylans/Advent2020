package main

import (
	"fmt"
	"github.com/Tommylans/Advent2020/day1"
	"time"
)

func main() {
	start := time.Now()
	part1 := day1.StartPart1()
	fmt.Println("Day1:p1 Answer", part1)
	fmt.Println("Day1:p1", time.Since(start))
	
	start = time.Now()
	part2 := day1.StartPart2()
	fmt.Println("Day1:p2 Answer", part2)
	fmt.Println("Day1:p2", time.Since(start))
}
