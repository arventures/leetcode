package _347

func topKFrequent(nums []int, k int) []int {

	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	var result []int

	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		for _, val := range buckets[i] {
			result = append(result, val)
		}
	}

	return result

}
