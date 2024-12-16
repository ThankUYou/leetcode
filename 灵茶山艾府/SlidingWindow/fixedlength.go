package slidingwindow

// 1456. 定长子串中元音的最大数目, medium
func maxVowels(s string, k int) (res int) {
	left, right := 0, 0
	num := 0
	for ; right < len(s); right++ {

		if s[right] == 'a' || s[right] == 'e' || s[right] == 'i' || s[right] == 'o' || s[right] == 'u' {
			num++
		}
		// 如果窗口大小超出 k，则从左侧移除
		if right-left+1 > k {
			if s[left] == 'a' || s[left] == 'e' || s[left] == 'i' || s[left] == 'o' || s[left] == 'u' {
				num--
			}
			left++
		}
		// 更新最大值
		res = max(res, num)
	}
	return
}

// 643. 子数组最大平均数 I, easy
func findMaxAverage(nums []int, k int) float64 {
	left := 0
	sum := 0
	maxSum := -1e9
	for right := 0; right < len(nums); right++ {
		// 右值添加
		sum += nums[right]
		if right-left+1 >= k {
			maxSum = max(maxSum, float64(sum))
			// 左值删除
			sum -= nums[left]
			left++
		}

	}
	return float64(maxSum) / float64(k)
}

// 1343. 大小为 K 且平均值大于等于阈值的子数组数目
func numOfSubarrays(arr []int, k int, threshold int) int {
	left := 0
	res := 0
	sum := 0

	for right := 0; right < len(arr); right++ {
		// 右值新增
		sum += arr[right]
		if right-left+1 >= k {
			// 左值删除
			if sum >= k*threshold {
				res++
			}
			sum -= arr[left]
			left++
		}
	}
	return res
}

// 2090. 半径为 k 的子数组平均值
func getAverages(nums []int, k int) []int {
	avg := make([]int, len(nums))
	sum := 0
	for right := 0; right < len(nums); right++ {
		if right < k || right+k > len(nums) {
			avg[right] = -1
		}
		sum += avg[right]
		if right >= 2*k {
			// avg[right-k] = sum / (2*k + 1)
			// sum -= nums[right-2*k]
		}
	}
	return avg
}

// 1297. 子串的最大出现次数
// 不同字母 -> map
// 次数 -> map
func maxFreq(s string, maxLetters int, minSize int, maxSize int) (res int) {
	strcnt := make(map[string]int)
	charcnt := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		charcnt[s[right]]++
		if right-left+1 > minSize {
			charcnt[s[left]]--
			if charcnt[s[left]] == 0 {
				delete(charcnt, s[left])
			}
			left++
		}

		if right-left+1 == minSize && len(charcnt) <= maxLetters {
			str := s[left : right+1]
			strcnt[str]++
			res = max(res, strcnt[str])
		}
	}
	return
}

