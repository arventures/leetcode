package _1929

func GetConcatenation(nums []int) []int {

	lenNums := len(nums)
	ans := make([]int, 2*lenNums)

	for i, num := range nums {
		ans[i] = num
		ans[i+lenNums] = num
	}

	return ans
}
