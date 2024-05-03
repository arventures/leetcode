package _74

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	lRow, rRow := 0, len(matrix)-1

	for lRow <= rRow {
		midRow := lRow + (rRow-lRow)/2
		row := matrix[midRow]

		if target >= row[0] && target <= row[len(row)-1] {
			return searchRow(row, target)
		} else if target < row[0] {
			rRow = midRow - 1
		} else {
			lRow = midRow + 1
		}
	}

	return false
}

func searchRow(row []int, target int) bool {
	l, r := 0, len(row)-1

	for l <= r {
		mid := l + (r-l)/2
		if row[mid] == target {
			return true
		} else if target < row[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
