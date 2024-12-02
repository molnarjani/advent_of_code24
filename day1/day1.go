package day1

import (
	"math"
	"sort"
)

func sortDesc(l []int) {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
}

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

func PrintSolution() {
	print("solved")
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
