package _704

func Search(nums []int, target int) int {
	lo := 0
	hi := len(nums)

	for lo < hi {
		mid := lo + (hi - lo) / 2

		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return -1
}
