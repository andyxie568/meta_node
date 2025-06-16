package main

import "sort"

/*
	给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，

其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 所有的数相异或，最终的结果就是值出现一次的元素
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}

	nums := make([]int, 0, 64)
	for x > 0 {
		nums = append(nums, x%10)
		x /= 10
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			} else {
				top := stack[len(stack)-1]
				if top == '(' && s[i] == ')' || top == '{' && s[i] == '}' || top == '[' && s[i] == ']' {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := make([]byte, 0)
	baseStr := strs[0]
	for i := 0; i < len(baseStr); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) < i+1 {
				return string(result)
			}
			if strs[j][i] != baseStr[i] {
				return string(result)
			}
		}
		result = append(result, baseStr[i])
	}
	return string(result)
}

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
*/

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。
*/
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}

	digits = reverse(digits)
	digits[0] += 1
	var carry int
	for i := 0; i < len(digits); i++ {
		digits[i] += carry
		if digits[i] >= 10 {
			digits[i] %= 10
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		digits = append(digits, carry)
	}
	return reverse(digits)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	for _, interval := range intervals {
		if len(interval) != 2 {
			return nil
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else if intervals[i][1] > end {
			end = intervals[i][1]
		}
	}
	result = append(result, []int{start, end})
	return result
}

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
*/
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if index, ok := numMap[target-nums[i]]; ok && index != i {
			return []int{index, i}
		}
	}
	return nil
}
