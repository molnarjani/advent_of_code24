package day1

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// solve gives the solution to day1 challenge by:
// 1. sorting list1 in descending order
// 2. sorting list2 in descending order
// 3. calculating difference of each number in the lists
// 4. adding the numbers together
func solveA(l1 []int, l2 []int) int {
	if len(l1) != len(l2) {
		panic("length of lists mismatching")
	}

	sortDesc(l1)
	sortDesc(l2)

	totalDiff := 0
	for i := range l1 {
		diff := math.Abs(float64(l1[i]) - float64(l2[i]))
		totalDiff += int(diff)
	}

	return totalDiff
}

// solve gives the solution to day1 second challenge by:
// for each number in list1
// 1. calculate the nubmer of occurences in list2
// 2. multiply the number by the number of occurences
// 3. add together all the result values
func solveB(l1 []int, l2 []int) int {
	if len(l1) != len(l2) {
		panic("length of lists mismatching")
	}
	totalCount := 0
	for i := range l1 {
		elemCount := 0
		for j := range l2 {
			if l1[i] == l2[j] {
				elemCount += 1
			}
		}
		totalCount += l1[i] * elemCount
	}

	return totalCount
}

// UTILS
// utility function to sort a slice in descending order
func sortDesc(l []int) {
	sort.Slice(l, func(i, j int) bool {
		return l[i] > l[j]
	})
}

// ReadInput reads the input file and returns the list of numbers
func ReadInput(filepath string) ([]int, []int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, err
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return list1, list2, nil
}
