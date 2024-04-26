package _215

func FindKthLargest(nums []int, k int) int {

	k = len(nums) - k

	return quickSelect(0, len(nums)-1, nums, k)

}

func quickSelect(left, right int, nums []int, k int) int {
	pivot := nums[right]
	p := left

	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[right] = pivot, nums[p]

	if p > k {
		return quickSelect(left, p-1, nums, k)
	} else if p < k {
		return quickSelect(p+1, right, nums, k)
	} else {
		return nums[p]
	}

}
