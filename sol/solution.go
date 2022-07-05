package sol

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for row := range dp {
		dp[row] = make([]int, n+1)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// start from last Idx reduce the repetition
	// dp[i][j] = dp[i+1][j+1] if text1[i] == text2[j]
	// dp[i][j] = max(dp[i+1][j], dp[i][j+1]) if text1[i] != text2[j]
	for text1_start := m - 1; text1_start >= 0; text1_start-- {
		for text2_start := n - 1; text2_start >= 0; text2_start-- {
			if text1[text1_start] == text2[text2_start] {
				dp[text1_start][text2_start] = dp[text1_start+1][text2_start+1] + 1
			} else {
				dp[text1_start][text2_start] = max(
					dp[text1_start+1][text2_start],
					dp[text1_start][text2_start+1],
				)
			}
		}
	}
	return dp[0][0]
}
