package day2

import (
	"strconv"
)

func (r *Rule) TestLetters(against string) bool {
	// xor
	return (r.Letter == string(against[r.FirstPart-1])) != (r.Letter == string(against[r.LastPart-1]))
}

func StartPart2() string {
	input := parseInput()

	correct := 0

	for _, s := range input {
		rule, password := parseRuleAndInput(s)

		if rule.TestLetters(password) {
			correct++
		}
	}

	return strconv.Itoa(correct)
}
