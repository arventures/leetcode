package _4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)

	halfLen := (m + n + 1) / 2

	left, right := 0, m

	for left <= right {
		i := (left + right) / 2

		j := halfLen - i

		if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else if i < m && nums[i] < nums2[j-1] {
			left = i + 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int

			minRight = min(nums1[i], nums2[j])

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
