package _695

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)

	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	var maxArea int

	visited := make(map[[2]int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[[2]int{i, j}] {
				area := dfs(i, j, grid, visited)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea

}

func dfs(i, j int, grid [][]int, visited map[[2]int]bool) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 || visited[[2]int{i, j}] {
		return 0
	}

	visited[[2]int{i, j}] = true

	area := 1
	area += dfs(i-1, j, grid, visited) // up
	area += dfs(i+1, j, grid, visited) // down
	area += dfs(i, j-1, grid, visited) // left
	area += dfs(i, j+1, grid, visited) // right

	return area
}
