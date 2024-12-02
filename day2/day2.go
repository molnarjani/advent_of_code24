package day2

import (
	"math"
	"sort"
)

type Report struct {
	levels []int
}

func sortDesc(l []int) {
	sort.Slice(l, func(i, j int) bool {
		return l[i] < l[j]
	})
}

func sortAsc(l []int) {
	sort.Slice(l, func(i, j int) bool {
		return l[i] > l[j]
	})
}

// calculateValidityA returns 1 if report is valid, 0 if report is invalid
// validity is calulated by:
//
//	a, check for repeating numbers in list
//	b, checking if numbers in report is all increasing or all decreasing
//	c, checking if adjacent numbers in list are increasing or decreasing by at least 1 and at most 3
func calculateValidityA(r Report) int {
	// check all desc or all inc
	descCopy := make([]int, len(r.levels))
	copy(descCopy, r.levels)
	sortDesc(descCopy)

	ascCopy := make([]int, len(r.levels))
	copy(ascCopy, r.levels)
	sortAsc(ascCopy)

	ascMatch := true
	descMatch := true
	for i := range r.levels {
		if r.levels[i] != ascCopy[i] {
			ascMatch = false
			break
		}
	}
	for i := range r.levels {
		if r.levels[i] != descCopy[i] {
			descMatch = false
			break
		}
	}
	if !ascMatch && !descMatch {
		return 0
	}

	// check for diff atleast 1, atmost 3
	for i := 0; i < len(r.levels)-1; i++ {
		c := r.levels[i]
		n := r.levels[i+1]
		diff := math.Abs(float64(c) - float64(n))
		if diff < 1 || diff > 3 {
			return 0
		}
	}

	return 1
}

// solve gives the solution to day2 challenge by:
// 1. calculating if a report is valid by:
//  1. a, checking if numbers in report is all increasing or all decreasing
//  1. b, checking if adjacent numbers in list are increasing or decreasing by at least 1 and at most 3
//
// 2. adding the valid reports together
func solveA(reports []Report) int {
	validCount := 0
	for _, r := range reports {
		validCount += calculateValidityA(r)
	}
	return validCount
}

// solve gives the solution to day2 challenge by:
// calculates validity using calculateValidityA
// if the result is wrong, tries removing each element from the list
// and retry if removing one element solves the problem
func solveB(reports []Report) int {
	validCount := 0
	for _, r := range reports {
		valid := calculateValidityA(r)
		if valid == 1 {
			validCount += 1
		} else {
			// make a new report for each possible variation for removing an item from the list
			levelLen := len(r.levels)
			for i := 0; i <= levelLen-1; i++ {
				start := append([]int{}, r.levels[0:i]...)
				end := append([]int{}, r.levels[i:levelLen]...)
				concat := append(start, end[1:]...)
				valid := calculateValidityA(Report{levels: concat})
				if valid == 1 {
					validCount += 1
					break
				}

			}
		}
	}
	return validCount
}