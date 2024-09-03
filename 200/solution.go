package _200

func numIslands(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])
	numIslands := 0

	visited := make(map[[2]int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[[2]int{i, j}] {
				numIslands++
				dfs(grid, i, j, rows, cols, visited)
			}
		}
	}
	return numIslands
}

func dfs(grid [][]byte, i, j, rows, cols int, visited map[[2]int]bool) {

	if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == '0' || visited[[2]int{i, j}] {
		return
	}

	visited[[2]int{i, j}] = true

	dfs(grid, i-1, j, rows, cols, visited)
	dfs(grid, i+1, j, rows, cols, visited)
	dfs(grid, i, j-1, rows, cols, visited)
	dfs(grid, i, j+1, rows, cols, visited)
}
