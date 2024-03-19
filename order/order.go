package order

import "math"

// Declare pack sizes
var packSizes = []int{5, 7, 9}

func FullFillOrder(amount int) (int, map[int]int) {
	// Initialize DP arrays for the minimum packs and for tracking the packs used
	dp := make([]int, amount+1)
	lastPack := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
		lastPack[i] = -1
	}

	// Base case
	dp[0] = 0

	// Fill the DP and lastPack arrays
	for a := 1; a <= amount; a++ {
		for _, d := range packSizes {
			if a-d >= 0 && dp[a-d]+1 < dp[a] {
				dp[a] = dp[a-d] + 1
				lastPack[a] = d
			}
		}
	}

	// Reconstruct the combination of packs
	packs := make(map[int]int)
	if dp[amount] != math.MaxInt32 {
		for a := amount; a > 0; a -= lastPack[a] {
			packs[lastPack[a]] += 1
		}
		return amount, packs
	}

	// Try to pack it with 1 more item
	return FullFillOrder(amount + 1)
}
