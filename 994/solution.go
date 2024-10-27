package _994

func orangesRotting(grid [][]int) int {

	var directions = [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1}
	}

	rows, cols := len(grid), len(grid[0])
	queue := [][]int{}
	freshOranges := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			} else if grid[r][c] == 1 {
				freshOranges++
			}
		}
	}

	minutes := 0

	for len(queue) > 0 {
		newQueue := [][]int{}

		for _, pos := range queue {
			r, c := pos[0], pos[1]

			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] == 2
					freshOranges--
					newQueue = append(newQueue, []int{nr, nc})
				}
			}
		}

		queue = newQueue

		if len(queue) > 0 {
			minutes++
		}
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes


}