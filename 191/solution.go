package _191

func hammingWeight(n int) int {
	resp := 0

	for n > 0 {
		n &= n - 1
		resp++
	}

	return resp
}
