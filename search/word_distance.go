package search

import "math"

func FindDistance(a, b string) int {
	aLen := len(a)
	bLen := len(b)

	dp := make([][]int, aLen+1)
	for i := 0; i < aLen+1; i++ {
		dp[i] = make([]int, bLen+1)
	}

	for i := 0; i <= aLen; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= bLen; i++ {
		dp[0][i] = i
	}

	for aIdx := 1; aIdx <= aLen; aIdx++ {
		for bIdx := 1; bIdx <= bLen; bIdx++ {
			if a[aIdx-1] != b[bIdx-1] {
				dp[aIdx][bIdx] = min(
					1+dp[aIdx-1][bIdx],
					1+dp[aIdx][bIdx-1],
					1+dp[aIdx-1][bIdx-1])
			} else {
				dp[aIdx][bIdx] = dp[aIdx-1][bIdx-1]
			}
		}
	}

	return dp[aLen][bLen]
}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}
