package _912

func SortArray(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		j := i - 1

		for j >= 0 && nums[j] > nums[j+1] {
			temp := nums[j+1]
			nums[j+1] = nums[j]
			nums[j] = temp
			j--
		}

	}
	return nums
}
