package _338

func countBits(n int) []int {
	dp := make([]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}

	return dp
}
