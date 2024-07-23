package _78

func subsets(nums []int) [][]int {

	var res [][]int
	var subset []int

	var dfs func(int)
	dfs = func(idx int) {
		if idx >= len(nums) {
			res = append(res, append([]int(nil), subset...))
			return
		}

		subset = append(subset, nums[idx])
		dfs(idx + 1)

		subset = subset[:len(subset)-1]
		dfs(idx + 1)
	}

	dfs(0)
	return res
}
