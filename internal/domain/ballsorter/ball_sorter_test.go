package ballsorter

import (
	"testing"
)

type testCase struct {
	name           string
	containers     [][]int
	expectedResult bool
}

func TestCanSortPositive(t *testing.T) {
	testCases := []testCase{
		{
			name: "positive_1",
			containers: [][]int{
				{1, 2},
				{2, 1},
			},
			expectedResult: true,
		},
		{
			name: "positive_2",
			containers: [][]int{
				{1},
			},
			expectedResult: true,
		},
		{
			name: "positive_3",
			containers: [][]int{
				{1, 2, 3},
				{2, 3, 1},
				{3, 1, 2},
			},
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := CanSort(tc.containers)
			if actualResult != tc.expectedResult {
				t.Errorf("got: %v, expected: %v", actualResult, tc.expectedResult)
			}
		})
	}
}

func TestCanSortNegative(t *testing.T) {
	testCases := []testCase{
		{
			name: "negative_1",
			containers: [][]int{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			},
			expectedResult: false,
		},
		{
			name: "negative_2",
			containers: [][]int{
				{10, 20, 30},
				{15, 25, 35},
				{5, 12, 19},
			},
			expectedResult: false,
		},
		{
			name: "negative_3",
			containers: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := CanSort(tc.containers)
			if actualResult != tc.expectedResult {
				t.Errorf("got: %v, expected: %v", actualResult, tc.expectedResult)
			}
		})
	}
}
