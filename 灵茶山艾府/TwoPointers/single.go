package twopointers

import (
	"math"
	"sort"
	"strings"
)

// 344. 反转字符串
// 要求原地修改
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 125. 验证回文串
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var check func(b byte) bool
	check = func(b byte) bool {
		return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
	}
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !check(s[left]) {
			left++
		}
		for left < right && !check(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 1750. 删除字符串两端相同字符后的最短长度
// 左右两侧要去除所有的重复元素
func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		tmp := s[left]
		for left <= right && s[left] == tmp {
			left++
		}
		for right >= left && s[right] == tmp {
			right--
		}
	}
	return right - left + 1
}

// 2105. 给植物浇水 II
func minimumRefill(plants []int, capacityA int, capacityB int) (res int) {
	left, right := 0, len(plants)-1
	a, b := capacityA, capacityB
	for left < right {
		if a < plants[left] { //  <= 会多一次
			a = capacityA
			res++
		}
		a -= plants[left]
		left++
		if b < plants[right] {
			b = capacityB
			res++
		}
		b -= plants[right]
		right--
	}
	if right == left && max(a, b) < plants[left] {
		res++
	}
	return
}

// 977. 有序数组的平方
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}
	return res
}

// 658. 找到 K 个最接近的元素
// 方法一：排序
func findClosestElements(arr []int, k, x int) []int {
	// 稳定排序，在绝对值相同的情况下，保证更小的数排在前面
	sort.SliceStable(arr, func(i, j int) bool { return abs(arr[i]-x) < abs(arr[j]-x) })
	arr = arr[:k]
	sort.Ints(arr)
	return arr
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 1471. 数组中的 k 个最强值
func getStrongest(arr []int, k int) (res []int) {
	// 排序
	sort.Ints(arr)
	//中位数
	mid := arr[(len(arr)-1)/2]
	left, right := 0, len(arr)-1
	for k > 0 {
		if arr[right]-mid > mid-arr[left] {
			res = append(res, arr[right])
			right--
		} else {
			res = append(res, arr[left])
			left++
		}
		k--
	}
	return
}

// 167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

// 633. 平方数之和
func judgeSquareSum(c int) bool {
	n := math.Sqrt(float64(c))
	left, right := 0, n
	// ==， 1+1 = 2
	for left <= int(right) {
		if int(left)*int(left)+int(right)*int(right) == c {
			return true
		} else if int(left)*int(left)+int(right)*int(right) > c {
			right--
		} else {
			left++
		}
	}
	return false
}

// 2824. 统计和小于目标的下标对数目
func countPairs(nums []int, target int) (res int) {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] >= target {
			right--
		} else {
			res += right - left
			left++
		}
	}
	return
}

// LCP 28. 采购方案
func purchasePlans(nums []int, target int) (res int) {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			res += right - left
			// 忘记模
			res = res % (1e9 + 7)
			left++
		}
	}
	return
}

// 15. 三数之和
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	for first := 0; first < len(nums); first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}

			if second == third {
				break
			}

			if nums[first]+nums[second]+nums[third] == 0 {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}

