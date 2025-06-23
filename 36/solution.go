package _36

func isValidSudoku(board [][]byte) bool {
	L := 9
	rows := make([]map[byte]bool, L)
	cols := make([]map[byte]bool, L)
	squares := make([]map[byte]bool, L)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		squares[i] = make(map[byte]bool)
	}

	for r := 0; r < L; r++ {
		for c := 0; c < L; c++ {
			value := board[r][c]

			if string(value) == "." {
				continue
			}

			squareIdx := (r/3)*3 + c/3

			if rows[r][value] || cols[c][value] || squares[squareIdx][value] {
				return false
			}

			rows[r][value] = true
			cols[c][value] = true
			squares[squareIdx][value] = true

		}
	}

	return true

}
