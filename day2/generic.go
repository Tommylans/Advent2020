package day2

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Rule struct {
	FirstPart int
	LastPart  int
	Letter    string
}

var (
	lineRegex = regexp.MustCompile("([0-9]+)-([0-9]+) (.*): (.*)")
)

func parseInput() (out []string) {
	file, _ := os.Open("./day2/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return
}

func parseRuleAndInput(line string) (Rule, string) {
	matches := lineRegex.FindStringSubmatch(line)

	min, _ := strconv.Atoi(matches[1])
	max, _ := strconv.Atoi(matches[2])

	return Rule{
		FirstPart: min,
		LastPart:  max,
		Letter:    matches[3],
	}, matches[4]
}
