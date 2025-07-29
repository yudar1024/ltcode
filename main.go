package main

import (
	"fmt"
	"sort"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	// 无固定容量滑动窗口
	window := []string{}

	maxLength := 0
	for i := 0; i < len(s); i++ {
		if strings.Contains(strings.Join(window, ""), string(s[i])) {
			maxLength = max(maxLength, len(window))
			window = window[1:]
		} else {
			window = append(window, string(s[i]))
		}
	}

	return maxLength
}

func minWindow(s string, t string) string {
	// 记录目标字符串的所有字符出现次数，在此场景中，目标为t
	targetHash := make(map[rune]int)
	// 将目标字符串的所有字符出现次数放到哈希表中
	for _, char := range t {
		targetHash[char]++
	}
	windowLef := 0
	windowRight := 0

	validHash := make(map[rune]int)

	// 当目标字符串的所有字符都在验证哈西表的时候，说明包含了所有的目标字符串
	check := func() bool {
		for key, value := range targetHash {
			if validHash[key] < value {
				return false
			}
		}
		return true
	}

	for windowRight < len(s) {
		// 扩展窗口
		validHash[rune(s[windowRight])]++
		windowRight++
		// 当窗口包含了所有的目标字符串
		for check() {
			// 更新最小窗口
			if windowRight-windowLef < len(s) {
				s = s[windowLef:windowRight]
			}
			// 收缩窗口
			validHash[rune(s[windowLef])]--
			windowLef++
		}
	}

	return ""
}

func merge(intervals [][]int) [][]int {
	// 先按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := [][]int{}
	for _, interval := range intervals {
		// 如果merged为空，或者当前区间与merged中最后一个区间不重叠，直接添加
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			// 否则，合并当前区间与merged中最后一个区间
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	reverse := func(nums []int, start int, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	// 反转整个数组
	reverse(nums, 0, len(nums)-1)
	// 反转前k个元素
	reverse(nums, 0, k-1)
	// 反转剩余的元素
	reverse(nums, k, len(nums)-1)
}

func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 1, 1

	for i := 0; i < n; i++ {
		result[i] = left
		left *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}

func main() {
	// 测试
	nums := []int{1, 2, 3, 4}
	k := 2
	rotate(nums, k)
	fmt.Println(nums) // 输出: [3 4 1 2]
}
