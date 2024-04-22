package _912

// Insertion Sort
//func SortArray(nums []int) []int {
//	for i := 1; i < len(nums); i++ {
//		j := i - 1
//
//		for j >= 0 && nums[j] > nums[j+1] {
//			temp := nums[j+1]
//			nums[j+1] = nums[j]
//			nums[j] = temp
//			j--
//		}
//
//	}
//	return nums
//}

// Merge sort
func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

func quickSort(arr []int, s, e int) {
	if e <= s {
		return
	}

	left := s
	pivot := arr[e]

	for i := s; i < e; i++ {
		if arr[i] < pivot {
			temp := arr[left]
			arr[left] = arr[i]
			arr[i] = temp
			left++
		}
	}

	arr[left], arr[e] = arr[e], arr[left]

	quickSort(arr, s, left-1)
	quickSort(arr, left+1, e)
}
