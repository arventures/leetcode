package _875

func MinEatingSpeed(piles []int, h int) int {

	lo := 1
	hi := 1

	for _, num := range piles {
		if num > hi {
			hi = num
		}
	}

	for lo <= hi {
		mid := lo + (hi-lo)/2

		totalHours := 0

		for _, num := range piles {
			totalHours += int(math.Ceil(float64(num) / float64(mid)))
		}

		if totalHours == h {
			return mid
		} else if totalHours < h {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}
