package _682

import "strconv"

func CalPoints(operations []string) int {
	var nums []int
	sum := 0

	for _, o := range operations {
		switch o {
		case "+":
			lastEl := nums[len(nums)-1]
			secondLastEl := nums[len(nums)-2]
			nums = append(nums, lastEl+secondLastEl)
		case "C":
			nums = nums[:len(nums)-1]
		case "D":
			lastEl := nums[len(nums)-1]
			nums = append(nums, 2*lastEl)
		default:
			x, _ := strconv.Atoi(o)
			nums = append(nums, x)
		}
	}
	for _, num := range nums {
		sum = sum + num
	}
	return sum
}
