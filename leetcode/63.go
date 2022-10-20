package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := range obstacleGrid {
		for j := range obstacleGrid[0] {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i-1 < 0 && j-1 < 0 {
				dp[i][j] = 1
			} else if i-1 < 0 {
				dp[i][j] = dp[i][j-1]
			} else if j-1 < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
