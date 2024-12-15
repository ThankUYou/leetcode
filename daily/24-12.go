package daily

import (
	"math"
	"slices"
)

// 12.23
// 3264. K 次乘运算后的最终数组 I
func getFinalState(nums []int, k int, multiplier int) []int {
	for k > 0 {
		minx2(nums, multiplier)
		k--
	}
	return nums
}

func minx2(nums []int, multiplier int) {
	minVal := math.MaxInt
	index := 0
	for i, v := range nums {
		if v < minVal {
			minVal = v
			index = i
		}
	}
	nums[index] *= multiplier
}

// 1338. 数组大小减半
// 数组计数 + 排序 + 贪心
func minSetSize(arr []int) int {
	cnt := make([]int, slices.Max(arr)+1)
	for _, val := range arr {
		cnt[val]++
	}
	slices.SortFunc(cnt, func(a, b int) int { return b - a })
	s := 0
	for i, c := range cnt {
		s += c
		if s >= len(arr)/2 {
			return i + 1
		}
	}
	return 0
}
