package day2

import (
	"strconv"
	"strings"
)

func (r *Rule) TestMinMax(against string) bool {
	count := strings.Count(against, r.Letter)
	return count >= r.FirstPart && count <= r.LastPart
}

func StartPart1() string {
	input := parseInput()

	correct := 0

	for _, s := range input {
		rule, password := parseRuleAndInput(s)

		if rule.TestMinMax(password) {
			correct++
		}
	}

	return strconv.Itoa(correct)
}
