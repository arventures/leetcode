package _1091

var directions = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func shortestPathBinaryMatrix(grid [][]int) int {

	ROWS := len(grid)
	COLS := len(grid[0])

	if grid[0][0] == 1 || grid[ROWS-1][COLS-1] == 1 {
		return -1
	}

	visited := make(map[[2]int]bool)
	queue := [][]int{{0, 0}}
	visited[[2]int{0, 0}] = true
	distance := make(map[[2]int]int)
	distance[[2]int{0, 0}] = 1

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		x, y := cell[0], cell[1]

		if x == ROWS-1 && y == COLS-1 {
			return distance[[2]int{x, y}]
		}

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]

			if newX >= 0 && newX < ROWS && newY >= 0 && newY < COLS && grid[newX][newY] == 0 {
				coord := [2]int{newX, newY} // Coordinate for the new cell
				if !visited[coord] {
					// Mark as visited and store the distance
					visited[coord] = true
					distance[coord] = distance[[2]int{x, y}] + 1
					queue = append(queue, []int{newX, newY}) // Add the new cell to the queue
				}
			}
		}
	}

	return -1

}
