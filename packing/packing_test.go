package packing

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindOptimalPacking(t *testing.T) {
	// setup test cases
	pack_sizes_1 := []int{250, 500, 1000, 2000, 5000}
	pack_sizes_2 := []int{23, 31, 53}

	test_cases := []struct {
		numItems  int
		packSizes []int
		expected  Solution
	}{
		{1, pack_sizes_1, Solution{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}},
		{250, pack_sizes_1, Solution{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}},
		{251, pack_sizes_1, Solution{250: 0, 500: 1, 1000: 0, 2000: 0, 5000: 0}},
		{501, pack_sizes_1, Solution{250: 1, 500: 1, 1000: 0, 2000: 0, 5000: 0}},
		{12001, pack_sizes_1, Solution{250: 1, 500: 0, 1000: 0, 2000: 1, 5000: 2}},
		{263, pack_sizes_2, Solution{23: 2, 31: 7, 53: 0}},
	}

	// for each test case assert expected solution
	for _, test_case := range test_cases {
		t.Run(fmt.Sprintf("NumItems%d", test_case.numItems), func(t *testing.T) {
			result := FindOptimalPacking(test_case.numItems, test_case.packSizes)
			if !reflect.DeepEqual(result, test_case.expected) {
				t.Errorf("findOptimalPacking(%d, %v) = %v; want %v", test_case.numItems, test_case.packSizes, result, test_case.expected)
			}
		})
	}
}
