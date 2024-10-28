package _62

func uniquePaths(r int, c int) int {
	row := make([]int, c)

	for i := range row {
		row[i] = 1
	}

	for i := r - 1; i >=0 0; i-- {
		newRow := make([]int, c)
		newRow[c-1] = 1

		for j := c - 2; j >= 0; j-- {
			newRow[j] = newRow[j+1] + row[j]
		}

		row = newRow
	}

	return row[0]
}