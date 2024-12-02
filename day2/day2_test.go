package day2

import (
	"testing"
)

func TestDay2ASolution(t *testing.T) {
	challengeReports, err := ReadInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		reports       []Report
		expectedCount int
	}{
		{
			reports: []Report{
				{levels: []int{1, 2, 3}},
				{levels: []int{4, 3, 2}},
			},
			expectedCount: 2,
		},
		{
			reports: []Report{
				{levels: []int{7, 6, 4, 2, 1}},
				{levels: []int{1, 2, 7, 8, 9}},
				{levels: []int{9, 7, 6, 2, 1}},
				{levels: []int{1, 3, 2, 4, 5}},
				{levels: []int{8, 6, 4, 4, 1}},
				{levels: []int{1, 3, 6, 7, 9}},
			},
			expectedCount: 2,
		},
		{
			reports:       challengeReports,
			expectedCount: 282,
		},
	}

	for i, tt := range tests {
		solution := solveA(tt.reports)
		if solution != tt.expectedCount {
			t.Fatalf("tests[%d] solution wrong, expected=%d, got=%d", i, tt.expectedCount, solution)
		}
	}
}

func TestDay2BSolution(t *testing.T) {
	challengeReports, err := ReadInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		reports       []Report
		expectedCount int
	}{
		{
			reports: []Report{
				{levels: []int{1, 3, 3}},
				{levels: []int{7, 3, 2}},
			},
			expectedCount: 2,
		},
		{
			reports: []Report{
				{levels: []int{7, 6, 4, 2, 1}},
				{levels: []int{1, 2, 7, 8, 9}},
				{levels: []int{9, 7, 6, 2, 1}},
				{levels: []int{1, 3, 2, 4, 5}},
				{levels: []int{8, 6, 4, 4, 1}},
				{levels: []int{1, 3, 6, 7, 9}},
			},
			expectedCount: 4,
		},
		{
			reports:       challengeReports,
			expectedCount: 349,
		},
	}

	for i, tt := range tests {
		solution := solveB(tt.reports)
		if solution != tt.expectedCount {
			t.Fatalf("tests[%d] solution wrong, expected=%d, got=%d", i, tt.expectedCount, solution)
		}
	}
}
