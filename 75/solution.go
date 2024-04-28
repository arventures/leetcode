package _75

func SortColors(nums []int) []int {

	count := [3]int{0, 0, 0}

	for _, num := range nums {
		count[num]++
	}

	index := 0

	for i := 0; i < len(count); i++ {
		for j := 0; j < count[i]; j++ {
			nums[index] = i
			index++
		}
	}

	return nums
}
