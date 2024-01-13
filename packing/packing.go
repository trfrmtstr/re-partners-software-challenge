package packing

import "math"

// Solution type for easier readability - each key represents pack size and value how many packs of that size are in this specific solution
type Solution map[int]int

// Creates a 0 item solution - such solution will obviously contain no packs since 0 items are required
func createZeroItemSolution(packSizes []int) Solution {
	sol := make(Solution)
	for _, size := range packSizes {
		sol[size] = 0
	}
	return sol
}

// Deep copies a solution
func deepCopySolution(src Solution) Solution {
	copy := make(Solution)
	for key, value := range src {
		copy[key] = value
	}
	return copy
}

// Compute the total number of items in a solution
func totalItems(sol Solution) int {
	total := 0
	for size, count := range sol {
		total += size * count
	}
	return total
}

// Compute the total number of packs in a solution
func totalPacks(sol Solution) int {
	total := 0
	for _, count := range sol {
		total += count
	}
	return total
}

// To find optimal packing solution for given rules we'll utilize dynamic programming.
// Let's say we have 3 pack sizes A, B and C and let's denote minimum number of items shipped for n desired items as S(n).
//
// If we assume that we know all solutions S(k) where k < n then we can assume one of 3 following scenarios:
// 1. S(n) contains A pack, which means S(n) = S(n - A) + A
// 2. S(n) contains B pack, which means S(n) = S(n - B) + A
// 3. S(n) contains C pack, which means S(n) = S(n - C) + A
//
// Note that S(k) where k <= 0 is 0 since we don't need to ship anything if desired number of items is 0 or less than 0 (which even doesn't make any sense here).
//
// From above assumptions we'll just pick one which minimizes S(n) so:
//
// S(n) = MIN[S(n - A) + A, S(n - B) + B, S(n - A) + B]
//
// Since 3rd rule of challenge prefers solutions with minimal number of packs if it happens that there exits multiple subsolutions that both have minimum number
// of items we'll just pick one with minimum number of packs.
//
// Above assumes only 3 pack sizes but can be easily generalized to any number of packs which is what below function does.
//
// Since our assumption that we know all previous solutions S(k) was wrong we'll build solutions from ground up - from 1 to 'numItems'.
// Once we get to `numItems` we have answer for original solution and we just return it.
func FindOptimalPacking(numItems int, packSizes []int) Solution {
	// memoization map keeping track of optimal solutions for each number from 1 to 'numItems'
	memo := make(map[int]Solution)

	for i := 1; i <= numItems; i++ {
		// 2nd rule - we want min num of items to be sent
		minItems := math.MaxInt32
		// 3rd rule - we want min num of packs to be sent
		minPacks := math.MaxInt32
		// we'll need to report back optimal solution
		var bestSolution Solution

		for _, size := range packSizes {
			var currentSolution Solution

			if i-size <= 0 {
				// trivial solution S(k) where k <= 0
				currentSolution = createZeroItemSolution(packSizes)
			} else {
				// we have to do deep copy since we'll mutate map
				currentSolution = deepCopySolution(memo[i-size])
			}
			// we assumed some pack size so we increment this solution accordingly
			currentSolution[size]++

			items := totalItems(currentSolution)
			packs := totalPacks(currentSolution)

			// enforce 2nd rule
			if items < minItems {
				minItems = items
				minPacks = packs
				bestSolution = currentSolution
			}

			// enforce 3rd rule
			if items == minItems && packs < minPacks {
				minItems = items
				minPacks = packs
				bestSolution = currentSolution
			}
		}

		memo[i] = bestSolution
	}

	return memo[numItems]
}
