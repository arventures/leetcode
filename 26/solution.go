package _26

func RemoveDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	index := 1
	uniqueVal := nums[0]
	k := 1

	for _, num := range nums[1:] {
		if uniqueVal == num {
			continue
		}

		uniqueVal = num
		nums[index] = uniqueVal
		k++
		index++
	}

	return k
}
