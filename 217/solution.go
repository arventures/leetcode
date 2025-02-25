package _217

func containsDuplicate(nums []int) bool {
	data := make(map[int]bool)

	for _, v := range nums {
		if _, found := data[v]; found {
			return true
		} else {
			data[v] = true
		}
	}
	return false
}
