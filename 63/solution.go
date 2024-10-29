package _63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows, cols := len(obstacleGrid), len(obstacleGrid[0])

	if obstacleGrid[rows-1][cols-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	cache := make(map[[2]int]int)

	var dfs func(i, j int) int

	dfs = func(i, j int) int {
		if i >= rows || j >= cols || obstacleGrid[i][j] == 1 {
			return 0
		}

		if i == rows-1 && j == cols-1 {
			return 1
		}

		if val, exists := cache[[2]int{i, j}]; exists {
			return val
		}

		cache[[2]int{i, j}] = dfs(i+1, j) + dfs(i, j+1)

		return cache[[2]int{i, j}]
	}

	return dfs(0, 0)

}
