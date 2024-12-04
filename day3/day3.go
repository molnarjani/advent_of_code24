package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func solveA(input string) int {
	count := 0
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)

	numPattern := `\d+`
	reNumPtrn := regexp.MustCompile(numPattern)

	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		numbers := reNumPtrn.FindAllString(match, -1)
		if len(numbers) != 2 {
			fmt.Println("invalid, too many or too few numbers: ", match)
			continue
		}
		int1, err1 := strconv.Atoi(numbers[0])
		int2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("invalid, not a number: ", match)
		}
		count += int1 * int2
	}
	return count
}

func solveB(input string) int {
	mulEnabled := true
	count := 0
	pattern := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(pattern)

	numPattern := `\d+`
	reNumPtrn := regexp.MustCompile(numPattern)

	matches := re.FindAllString(input, -1)
	for _, match := range matches {
		if match == `do()` {
			mulEnabled = true
			continue
		}

		if match == `don't()` {
			mulEnabled = false
			continue
		}

		numbers := reNumPtrn.FindAllString(match, -1)
		if len(numbers) != 2 {
			fmt.Println("invalid, too many or too few numbers: ", match)
			continue
		}
		int1, err1 := strconv.Atoi(numbers[0])
		int2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			fmt.Println("invalid, not a number: ", match)
		}
		if mulEnabled {
			count += int1 * int2
		}
	}
	return count
}
