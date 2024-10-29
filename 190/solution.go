package _190

func reverseBits(n uint32) uint32 {
	var result uint32 = 0

	for i := 0; i < 32; i++ {

		result <<= 1

		result |= n & 1

		n >>= 1
	}

	return result
}
