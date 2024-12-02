package day1

import (
	"testing"
)

func TestSortDesc(t *testing.T) {
	l := []int{3, 4, 2, 1, 3, 3}
	sortDesc(l)
	expected := []int{4, 3, 3, 3, 2, 1}
	for i, elem := range l {
		if elem != expected[i] {
			t.Fatalf("sortDesc failed, expected=%v, got=%v", expected, l)
		}
	}
}

func TestDay1ASolution(t *testing.T) {
	challengeL1, challengeL2, err := ReadInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		list1    []int
		list2    []int
		expected int
	}{
		{
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			11,
		},
		{
			challengeL1,
			challengeL2,
			1580061,
		},
	}

	for i, tt := range tests {
		solution := solveA(tt.list1, tt.list2)
		if solution != tt.expected {
			t.Fatalf("tests[%d] solution is wrong, expected=%d, got=%d", i, tt.expected, solution)
		}

	}
}

func TestDay1BSolution(t *testing.T) {
	challengeL1, challengeL2, err := ReadInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		list1    []int
		list2    []int
		expected int
	}{
		{
			[]int{3, 4, 2, 1, 3, 3},
			[]int{4, 3, 5, 3, 9, 3},
			31,
		},
		{
			challengeL1,
			challengeL2,
			23046913,
		},
	}

	for i, tt := range tests {
		solution := solveB(tt.list1, tt.list2)
		if solution != tt.expected {
			t.Fatalf("tests[%d] solution is wrong, expected=%d, got=%d", i, tt.expected, solution)
		}

	}
}
