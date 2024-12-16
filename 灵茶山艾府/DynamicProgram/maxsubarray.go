package maxsubarray

// 53. 最大子数组和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		maxVal = max(nums[i], maxVal)
	}
	return maxVal
}

// 2606. 找到最大开销的子字符串
func maximumCostSubstring(s string, chars string, vals []int) (res int) {
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = i + 1
	}
	for i, c := range chars {
		cnt[c-'a'] = vals[i]
	}
	dp := 0
	for _, c := range s {
		dp = max(dp, 0) + cnt[c-'a']
		res = max(res, dp)
	}
	return
}

// 1749. 任意子数组和的绝对值的最大值
// 负值最小 / 正值最大
func maxAbsoluteSum(nums []int) (res int) {
	maxSum, minSum := 0, 0
	curMax, curMin := 0, 0

	for _, num := range nums {
		// 更新正向最大子数组和
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)

		// 更新负向最大子数组和
		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)
	}

	// 返回最大绝对值
	return max(maxSum, -minSum)
}

