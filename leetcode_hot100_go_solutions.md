# LeetCode Hot 100 - Go语言解题思路与详解

## 目录
- [1-20题](#1-20题)
- [21-40题](#21-40题)
- [41-60题](#41-60题)
- [61-80题](#61-80题)
- [81-100题](#81-100题)

---

## 1-20题

### 1. 两数之和 (Two Sum)
**难度：** Easy  
**题目描述：** 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。

**解题思路：**
使用哈希表存储已遍历的元素及其索引，对于每个元素，检查 target - nums[i] 是否在哈希表中。

```go
func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)
    for i, num := range nums {
        if index, exists := hashMap[target-num]; exists {
            return []int{index, i}
        }
        hashMap[num] = i
    }
    return nil
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 2. 两数相加 (Add Two Numbers)
**难度：** Medium  
**题目描述：** 给你两个非空的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。

**解题思路：**
模拟加法过程，从个位开始相加，处理进位。

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    carry := 0
    
    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        
        carry = sum / 10
        current.Next = &ListNode{Val: sum % 10}
        current = current.Next
    }
    
    return dummy.Next
}
```

**时间复杂度：** O(max(m,n))  
**空间复杂度：** O(max(m,n))

---

### 3. 无重复字符的最长子串 (Longest Substring Without Repeating Characters)
**难度：** Medium  
**题目描述：** 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

**解题思路：**
使用滑动窗口，维护一个字符到索引的映射，当遇到重复字符时，移动左指针。

```go
func lengthOfLongestSubstring(s string) int {
    charMap := make(map[byte]int)
    left, maxLen := 0, 0
    
    for right := 0; right < len(s); right++ {
        if index, exists := charMap[s[right]]; exists && index >= left {
            left = index + 1
        }
        charMap[s[right]] = right
        maxLen = max(maxLen, right-left+1)
    }
    
    return maxLen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(min(m,n))

---

### 4. 寻找两个正序数组的中位数 (Median of Two Sorted Arrays)
**难度：** Hard  
**题目描述：** 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

**解题思路：**
使用二分查找，将问题转化为寻找第k小的元素。

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)
    if total%2 == 1 {
        return float64(findKth(nums1, nums2, total/2+1))
    }
    return float64(findKth(nums1, nums2, total/2)+findKth(nums1, nums2, total/2+1)) / 2.0
}

func findKth(nums1, nums2 []int, k int) int {
    if len(nums1) > len(nums2) {
        return findKth(nums2, nums1, k)
    }
    
    if len(nums1) == 0 {
        return nums2[k-1]
    }
    
    if k == 1 {
        return min(nums1[0], nums2[0])
    }
    
    i := min(len(nums1), k/2)
    j := k - i
    
    if nums1[i-1] < nums2[j-1] {
        return findKth(nums1[i:], nums2, k-i)
    }
    return findKth(nums1, nums2[j:], k-j)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

**时间复杂度：** O(log(min(m,n)))  
**空间复杂度：** O(1)

---

### 5. 最长回文子串 (Longest Palindromic Substring)
**难度：** Medium  
**题目描述：** 给你一个字符串 s，找到 s 中最长的回文子串。

**解题思路：**
使用中心扩展法，以每个字符和每两个字符之间为中心向两边扩展。

```go
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    
    start, maxLen := 0, 1
    
    for i := 0; i < len(s); i++ {
        // 奇数长度回文串
        len1 := expandAroundCenter(s, i, i)
        // 偶数长度回文串
        len2 := expandAroundCenter(s, i, i+1)
        
        currentMax := max(len1, len2)
        if currentMax > maxLen {
            maxLen = currentMax
            start = i - (currentMax-1)/2
        }
    }
    
    return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(1)

---

### 6. 正则表达式匹配 (Regular Expression Matching)
**难度：** Hard  
**题目描述：** 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

**解题思路：**
使用动态规划，dp[i][j] 表示 s 的前 i 个字符和 p 的前 j 个字符是否匹配。

```go
func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    
    dp[0][0] = true
    
    // 处理 p 中 * 的情况
    for j := 2; j <= n; j += 2 {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-2]
        }
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if p[j-1] == '*' {
                // * 匹配 0 个字符
                dp[i][j] = dp[i][j-2]
                // * 匹配 1 个或多个字符
                if match(s[i-1], p[j-2]) {
                    dp[i][j] = dp[i][j] || dp[i-1][j]
                }
            } else {
                if match(s[i-1], p[j-1]) {
                    dp[i][j] = dp[i-1][j-1]
                }
            }
        }
    }
    
    return dp[m][n]
}

func match(s, p byte) bool {
    return p == '.' || s == p
}
```

**时间复杂度：** O(mn)  
**空间复杂度：** O(mn)

---

### 7. 盛最多水的容器 (Container With Most Water)
**难度：** Medium  
**题目描述：** 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai)。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, 0) 和 (i, ai)。

**解题思路：**
使用双指针，从两端开始向中间移动，每次移动较短的指针。

```go
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxWater := 0
    
    for left < right {
        width := right - left
        h := min(height[left], height[right])
        maxWater = max(maxWater, width*h)
        
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    
    return maxWater
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 8. 三数之和 (3Sum)
**难度：** Medium  
**题目描述：** 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c，使得 a + b + c = 0？请你找出所有满足条件且不重复的三元组。

**解题思路：**
先排序，然后固定一个数，使用双指针在剩余部分寻找两个数。

```go
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    var result [][]int
    
    for i := 0; i < len(nums)-2; i++ {
        // 跳过重复元素
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        left, right := i+1, len(nums)-1
        target := -nums[i]
        
        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                
                // 跳过重复元素
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                left++
                right--
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }
    
    return result
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(1)

---

### 9. 电话号码的字母组合 (Letter Combinations of a Phone Number)
**难度：** Medium  
**题目描述：** 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按任意顺序返回。

**解题思路：**
使用回溯算法，递归生成所有可能的组合。

```go
func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    phoneMap := map[byte]string{
        '2': "abc", '3': "def", '4': "ghi", '5': "jkl",
        '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
    }
    
    var result []string
    var backtrack func(index int, current string)
    
    backtrack = func(index int, current string) {
        if index == len(digits) {
            result = append(result, current)
            return
        }
        
        letters := phoneMap[digits[index]]
        for _, letter := range letters {
            backtrack(index+1, current+string(letter))
        }
    }
    
    backtrack(0, "")
    return result
}
```

**时间复杂度：** O(3^m × 4^n)  
**空间复杂度：** O(m+n)

---

### 10. 删除链表的倒数第N个结点 (Remove Nth Node From End of List)
**难度：** Medium  
**题目描述：** 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

**解题思路：**
使用双指针，先让快指针走n步，然后快慢指针同时移动。

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    first, second := dummy, dummy
    
    // 先让first指针走n+1步
    for i := 0; i <= n; i++ {
        first = first.Next
    }
    
    // 然后两个指针同时移动
    for first != nil {
        first = first.Next
        second = second.Next
    }
    
    // 删除节点
    second.Next = second.Next.Next
    
    return dummy.Next
}
```

**时间复杂度：** O(L)  
**空间复杂度：** O(1)

---

### 11. 有效的括号 (Valid Parentheses)
**难度：** Easy  
**题目描述：** 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s，判断字符串是否有效。

**解题思路：**
使用栈，遇到左括号入栈，遇到右括号检查栈顶是否匹配。

```go
func isValid(s string) bool {
    stack := []rune{}
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
    
    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 12. 合并两个有序链表 (Merge Two Sorted Lists)
**难度：** Easy  
**题目描述：** 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

**解题思路：**
使用双指针，比较两个链表的节点值，将较小的节点加入结果链表。

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    current := dummy
    
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            current.Next = l1
            l1 = l1.Next
        } else {
            current.Next = l2
            l2 = l2.Next
        }
        current = current.Next
    }
    
    // 连接剩余节点
    if l1 != nil {
        current.Next = l1
    } else {
        current.Next = l2
    }
    
    return dummy.Next
}
```

**时间复杂度：** O(m+n)  
**空间复杂度：** O(1)

---

### 13. 括号生成 (Generate Parentheses)
**难度：** Medium  
**题目描述：** 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且有效的括号组合。

**解题思路：**
使用回溯算法，在每一步选择添加左括号或右括号，但要保证有效性。

```go
func generateParenthesis(n int) []string {
    var result []string
    var backtrack func(current string, open, close int)
    
    backtrack = func(current string, open, close int) {
        if len(current) == 2*n {
            result = append(result, current)
            return
        }
        
        if open < n {
            backtrack(current+"(", open+1, close)
        }
        if close < open {
            backtrack(current+")", open, close+1)
        }
    }
    
    backtrack("", 0, 0)
    return result
}
```

**时间复杂度：** O(4^n / √n)  
**空间复杂度：** O(n)

---

### 14. 合并K个升序链表 (Merge k Sorted Lists)
**难度：** Hard  
**题目描述：** 给你一个链表数组，每个链表都已经按升序排列。请你将所有链表合并到一个升序链表中，返回合并后的链表。

**解题思路：**
使用分治法，将k个链表两两合并，直到只剩一个链表。

```go
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    
    for len(lists) > 1 {
        var merged []*ListNode
        for i := 0; i < len(lists); i += 2 {
            if i+1 < len(lists) {
                merged = append(merged, mergeTwoLists(lists[i], lists[i+1]))
            } else {
                merged = append(merged, lists[i])
            }
        }
        lists = merged
    }
    
    return lists[0]
}
```

**时间复杂度：** O(n log k)  
**空间复杂度：** O(log k)

---

### 15. 下一个排列 (Next Permutation)
**难度：** Medium  
**题目描述：** 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

**解题思路：**
1. 从右向左找到第一个升序对
2. 从右向左找到第一个大于该元素的值
3. 交换这两个元素
4. 反转后面的部分

```go
func nextPermutation(nums []int) {
    i := len(nums) - 2
    // 找到第一个升序对
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    
    if i >= 0 {
        j := len(nums) - 1
        // 找到第一个大于nums[i]的元素
        for j >= 0 && nums[i] >= nums[j] {
            j--
        }
        // 交换
        nums[i], nums[j] = nums[j], nums[i]
    }
    
    // 反转i+1到末尾的部分
    reverse(nums, i+1)
}

func reverse(nums []int, start int) {
    i, j := start, len(nums)-1
    for i < j {
        nums[i], nums[j] = nums[j], nums[i]
        i++
        j--
    }
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 16. 最长有效括号 (Longest Valid Parentheses)
**难度：** Hard  
**题目描述：** 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

**解题思路：**
使用栈，栈中存储索引，遇到左括号入栈，遇到右括号出栈并计算长度。

```go
func longestValidParentheses(s string) int {
    stack := []int{-1} // 栈底放-1作为边界
    maxLen := 0
    
    for i, char := range s {
        if char == '(' {
            stack = append(stack, i)
        } else {
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                stack = append(stack, i)
            } else {
                maxLen = max(maxLen, i-stack[len(stack)-1])
            }
        }
    }
    
    return maxLen
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 17. 搜索旋转排序数组 (Search in Rotated Sorted Array)
**难度：** Medium  
**题目描述：** 整数数组 nums 按升序排列，数组中的值互不相同。在传递给函数之前，nums 在预先未知的某个下标 k 上进行了旋转。

**解题思路：**
使用二分查找，判断哪一半是有序的，然后决定搜索方向。

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left <= right {
        mid := left + (right-left)/2
        
        if nums[mid] == target {
            return mid
        }
        
        // 判断左半部分是否有序
        if nums[left] <= nums[mid] {
            if target >= nums[left] && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // 右半部分有序
            if target > nums[mid] && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    
    return -1
}
```

**时间复杂度：** O(log n)  
**空间复杂度：** O(1)

---

### 18. 在排序数组中查找元素的第一个和最后一个位置 (Find First and Last Position of Element in Sorted Array)
**难度：** Medium  
**题目描述：** 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

**解题思路：**
使用二分查找分别找到第一个和最后一个位置。

```go
func searchRange(nums []int, target int) []int {
    left := findFirst(nums, target)
    if left == -1 {
        return []int{-1, -1}
    }
    right := findLast(nums, target)
    return []int{left, right}
}

func findFirst(nums []int, target int) int {
    left, right := 0, len(nums)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            result = mid
            right = mid - 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}

func findLast(nums []int, target int) int {
    left, right := 0, len(nums)-1
    result := -1
    
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            result = mid
            left = mid + 1
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return result
}
```

**时间复杂度：** O(log n)  
**空间复杂度：** O(1)

---

### 19. 组合总和 (Combination Sum)
**难度：** Medium  
**题目描述：** 给定一个无重复元素的数组 candidates 和一个目标数 target，找出 candidates 中所有可以使数字和为 target 的组合。

**解题思路：**
使用回溯算法，递归搜索所有可能的组合。

```go
func combinationSum(candidates []int, target int) [][]int {
    var result [][]int
    var backtrack func(start int, current []int, sum int)
    
    backtrack = func(start int, current []int, sum int) {
        if sum == target {
            temp := make([]int, len(current))
            copy(temp, current)
            result = append(result, temp)
            return
        }
        
        if sum > target {
            return
        }
        
        for i := start; i < len(candidates); i++ {
            current = append(current, candidates[i])
            backtrack(i, current, sum+candidates[i])
            current = current[:len(current)-1]
        }
    }
    
    backtrack(0, []int{}, 0)
    return result
}
```

**时间复杂度：** O(N^(T/M))  
**空间复杂度：** O(T/M)

---

### 20. 接雨水 (Trapping Rain Water)
**难度：** Hard  
**题目描述：** 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

**解题思路：**
使用双指针，维护左右最大值，每次移动较小的一边。

```go
func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }
    
    left, right := 0, len(height)-1
    leftMax, rightMax := 0, 0
    water := 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                water += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                water += rightMax - height[right]
            }
            right--
        }
    }
    
    return water
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

## 21-40题

### 21. 全排列 (Permutations)
**难度：** Medium  
**题目描述：** 给定一个不含重复数字的数组 nums，返回其所有可能的全排列。你可以按任意顺序返回答案。

**解题思路：**
使用回溯算法，递归生成所有可能的排列。

```go
func permute(nums []int) [][]int {
    var result [][]int
    var backtrack func(current []int, used []bool)
    
    backtrack = func(current []int, used []bool) {
        if len(current) == len(nums) {
            temp := make([]int, len(current))
            copy(temp, current)
            result = append(result, temp)
            return
        }
        
        for i := 0; i < len(nums); i++ {
            if !used[i] {
                used[i] = true
                current = append(current, nums[i])
                backtrack(current, used)
                current = current[:len(current)-1]
                used[i] = false
            }
        }
    }
    
    backtrack([]int{}, make([]bool, len(nums)))
    return result
}
```

**时间复杂度：** O(n! × n)  
**空间复杂度：** O(n)

---

### 22. 旋转图像 (Rotate Image)
**难度：** Medium  
**题目描述：** 给定一个 n × n 的二维矩阵表示一个图像。将图像顺时针旋转 90 度。

**解题思路：**
先转置矩阵，然后翻转每一行。

```go
func rotate(matrix [][]int) {
    n := len(matrix)
    
    // 转置矩阵
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    // 翻转每一行
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(1)

---

### 23. 字母异位词分组 (Group Anagrams)
**难度：** Medium  
**题目描述：** 给你一个字符串数组，请你将字母异位词组合在一起。可以按任意顺序返回结果列表。

**解题思路：**
使用哈希表，将排序后的字符串作为键，原字符串作为值。

```go
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    
    for _, str := range strs {
        key := sortString(str)
        groups[key] = append(groups[key], str)
    }
    
    var result [][]string
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}

func sortString(s string) string {
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool {
        return runes[i] < runes[j]
    })
    return string(runes)
}
```

**时间复杂度：** O(nk log k)  
**空间复杂度：** O(nk)

---

### 24. 最大子数组和 (Maximum Subarray)
**难度：** Easy  
**题目描述：** 给你一个整数数组 nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**解题思路：**
使用Kadane算法，维护当前最大和和全局最大和。

```go
func maxSubArray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxSum := nums[0]
    currentSum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum+nums[i])
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 25. 跳跃游戏 (Jump Game)
**难度：** Medium  
**题目描述：** 给定一个非负整数数组 nums，你最初位于数组的第一个下标。数组中的每个元素代表你在该位置可以跳跃的最大长度。

**解题思路：**
贪心算法，维护能到达的最远位置。

```go
func canJump(nums []int) bool {
    maxReach := 0
    
    for i := 0; i < len(nums); i++ {
        if i > maxReach {
            return false
        }
        maxReach = max(maxReach, i+nums[i])
        if maxReach >= len(nums)-1 {
            return true
        }
    }
    
    return true
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 26. 合并区间 (Merge Intervals)
**难度：** Medium  
**题目描述：** 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi]。请你合并所有重叠的区间，并返回一个不重叠的区间数组。

**解题思路：**
先排序，然后遍历合并重叠的区间。

```go
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // 按起始位置排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    var result [][]int
    current := intervals[0]
    
    for i := 1; i < len(intervals); i++ {
        if current[1] >= intervals[i][0] {
            // 重叠，合并
            current[1] = max(current[1], intervals[i][1])
        } else {
            // 不重叠，添加当前区间
            result = append(result, current)
            current = intervals[i]
        }
    }
    
    result = append(result, current)
    return result
}
```

**时间复杂度：** O(n log n)  
**空间复杂度：** O(1)

---

### 27. 不同路径 (Unique Paths)
**难度：** Medium  
**题目描述：** 一个机器人位于一个 m x n 网格的左上角。机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角。

**解题思路：**
动态规划，dp[i][j] = dp[i-1][j] + dp[i][j-1]。

```go
func uniquePaths(m int, n int) int {
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    
    // 初始化第一行和第一列
    for i := 0; i < m; i++ {
        dp[i][0] = 1
    }
    for j := 0; j < n; j++ {
        dp[0][j] = 1
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = dp[i-1][j] + dp[i][j-1]
        }
    }
    
    return dp[m-1][n-1]
}
```

**时间复杂度：** O(mn)  
**空间复杂度：** O(mn)

---

### 28. 最小路径和 (Minimum Path Sum)
**难度：** Medium  
**题目描述：** 给定一个包含非负整数的 m x n 网格 grid，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

**解题思路：**
动态规划，dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]。

```go
func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    
    // 初始化第一行
    for j := 1; j < n; j++ {
        grid[0][j] += grid[0][j-1]
    }
    
    // 初始化第一列
    for i := 1; i < m; i++ {
        grid[i][0] += grid[i-1][0]
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    
    return grid[m-1][n-1]
}
```

**时间复杂度：** O(mn)  
**空间复杂度：** O(1)

---

### 29. 爬楼梯 (Climbing Stairs)
**难度：** Easy  
**题目描述：** 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

**解题思路：**
动态规划，dp[i] = dp[i-1] + dp[i-2]。

```go
func climbStairs(n int) int {
    if n <= 2 {
        return n
    }
    
    prev2, prev1 := 1, 2
    
    for i := 3; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 30. 编辑距离 (Edit Distance)
**难度：** Hard  
**题目描述：** 给你两个单词 word1 和 word2，请返回将 word1 转换成 word2 所使用的最少操作数。

**解题思路：**
动态规划，dp[i][j] 表示 word1 前 i 个字符转换为 word2 前 j 个字符的最少操作数。

```go
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    // 初始化
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
            }
        }
    }
    
    return dp[m][n]
}
```

**时间复杂度：** O(mn)  
**空间复杂度：** O(mn)

---

### 31. 颜色分类 (Sort Colors)
**难度：** Medium  
**题目描述：** 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻。

**解题思路：**
三指针法，分别维护红色、白色、蓝色的边界。

```go
func sortColors(nums []int) {
    red, white, blue := 0, 0, len(nums)-1
    
    for white <= blue {
        switch nums[white] {
        case 0: // 红色
            nums[red], nums[white] = nums[white], nums[red]
            red++
            white++
        case 1: // 白色
            white++
        case 2: // 蓝色
            nums[white], nums[blue] = nums[blue], nums[white]
            blue--
        }
    }
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 32. 最小覆盖子串 (Minimum Window Substring)
**难度：** Hard  
**题目描述：** 给你一个字符串 s、一个字符串 t。返回 s 中涵盖 t 所有字符的最小子串。

**解题思路：**
滑动窗口，使用双指针维护窗口，用哈希表记录字符频率。

```go
func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }
    
    need := make(map[byte]int)
    for i := range t {
        need[t[i]]++
    }
    
    left, right := 0, 0
    valid := 0
    window := make(map[byte]int)
    start, length := 0, math.MaxInt32
    
    for right < len(s) {
        c := s[right]
        right++
        
        if need[c] > 0 {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        
        for valid == len(need) {
            if right-left < length {
                start = left
                length = right - left
            }
            
            d := s[left]
            left++
            
            if need[d] > 0 {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    
    if length == math.MaxInt32 {
        return ""
    }
    return s[start : start+length]
}
```

**时间复杂度：** O(|s| + |t|)  
**空间复杂度：** O(|s| + |t|)

---

### 33. 子集 (Subsets)
**难度：** Medium  
**题目描述：** 给你一个整数数组 nums，数组中的元素互不相同。返回该数组所有可能的子集（幂集）。

**解题思路：**
回溯算法，每个元素都有选择和不选择两种状态。

```go
func subsets(nums []int) [][]int {
    var result [][]int
    var backtrack func(start int, current []int)
    
    backtrack = func(start int, current []int) {
        temp := make([]int, len(current))
        copy(temp, current)
        result = append(result, temp)
        
        for i := start; i < len(nums); i++ {
            current = append(current, nums[i])
            backtrack(i+1, current)
            current = current[:len(current)-1]
        }
    }
    
    backtrack(0, []int{})
    return result
}
```

**时间复杂度：** O(2^n)  
**空间复杂度：** O(n)

---

### 34. 单词搜索 (Word Search)
**难度：** Medium  
**题目描述：** 给定一个 m x n 二维字符网格 board 和一个字符串单词 word。如果 word 存在于网格中，返回 true；否则，返回 false。

**解题思路：**
深度优先搜索，从每个位置开始搜索单词。

```go
func exist(board [][]byte, word string) bool {
    m, n := len(board), len(board[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(board, word, i, j, 0, visited) {
                return true
            }
        }
    }
    
    return false
}

func dfs(board [][]byte, word string, i, j, index int, visited [][]bool) bool {
    if index == len(word) {
        return true
    }
    
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
       visited[i][j] || board[i][j] != word[index] {
        return false
    }
    
    visited[i][j] = true
    
    directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
    for _, dir := range directions {
        if dfs(board, word, i+dir[0], j+dir[1], index+1, visited) {
            return true
        }
    }
    
    visited[i][j] = false
    return false
}
```

**时间复杂度：** O(mn × 4^L)  
**空间复杂度：** O(L)

---

### 35. 柱状图中最大的矩形 (Largest Rectangle in Histogram)
**难度：** Hard  
**题目描述：** 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1。

**解题思路：**
使用单调栈，维护一个递增的栈。

```go
func largestRectangleArea(heights []int) int {
    stack := []int{}
    maxArea := 0
    
    for i := 0; i <= len(heights); i++ {
        var h int
        if i < len(heights) {
            h = heights[i]
        } else {
            h = 0
        }
        
        for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            var width int
            if len(stack) == 0 {
                width = i
            } else {
                width = i - stack[len(stack)-1] - 1
            }
            
            maxArea = max(maxArea, height*width)
        }
        
        stack = append(stack, i)
    }
    
    return maxArea
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 36. 最大矩形 (Maximal Rectangle)
**难度：** Hard  
**题目描述：** 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

**解题思路：**
将问题转化为多个柱状图问题，使用单调栈求解。

```go
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
    
    m, n := len(matrix), len(matrix[0])
    heights := make([]int, n)
    maxArea := 0
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }
        maxArea = max(maxArea, largestRectangleArea(heights))
    }
    
    return maxArea
}
```

**时间复杂度：** O(mn)  
**空间复杂度：** O(n)

---

### 37. 二叉树中的最大路径和 (Binary Tree Maximum Path Sum)
**难度：** Hard  
**题目描述：** 路径被定义为一条从树中任意节点开始，沿父节点-子节点连接，达到任意节点的序列。

**解题思路：**
递归计算每个节点的最大贡献值，同时更新全局最大值。

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    maxSum := math.MinInt32
    
    var maxGain func(node *TreeNode) int
    maxGain = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftGain := max(maxGain(node.Left), 0)
        rightGain := max(maxGain(node.Right), 0)
        
        priceNewPath := node.Val + leftGain + rightGain
        maxSum = max(maxSum, priceNewPath)
        
        return node.Val + max(leftGain, rightGain)
    }
    
    maxGain(root)
    return maxSum
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 38. 最长连续序列 (Longest Consecutive Sequence)
**难度：** Medium  
**题目描述：** 给定一个未排序的整数数组 nums，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

**解题思路：**
使用哈希集合，对每个数字检查其前驱是否存在。

```go
func longestConsecutive(nums []int) int {
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }
    
    longestStreak := 0
    
    for num := range numSet {
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1
            
            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }
            
            longestStreak = max(longestStreak, currentStreak)
        }
    }
    
    return longestStreak
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 39. 只出现一次的数字 (Single Number)
**难度：** Easy  
**题目描述：** 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

**解题思路：**
使用异或运算，相同数字异或结果为0。

```go
func singleNumber(nums []int) int {
    result := 0
    for _, num := range nums {
        result ^= num
    }
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 40. 单词拆分 (Word Break)
**难度：** Medium  
**题目描述：** 给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

**解题思路：**
动态规划，dp[i] 表示前 i 个字符是否可以拆分。

```go
func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }
    
    dp := make([]bool, len(s)+1)
    dp[0] = true
    
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[len(s)]
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n)

---

## 41-60题

### 41. 环形链表 (Linked List Cycle)
**难度：** Easy  
**题目描述：** 给定一个链表，判断链表中是否有环。

**解题思路：**
使用快慢指针，如果存在环，快指针最终会追上慢指针。

```go
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    
    slow, fast := head, head.Next
    
    for fast != nil && fast.Next != nil {
        if slow == fast {
            return true
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    return false
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 42. 环形链表 II (Linked List Cycle II)
**难度：** Medium  
**题目描述：** 给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回 null。

**解题思路：**
使用快慢指针找到相遇点，然后一个指针从头开始，一个从相遇点开始，同时移动直到相遇。

```go
func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    
    slow, fast := head, head
    
    // 找到相遇点
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    
    if fast == nil || fast.Next == nil {
        return nil
    }
    
    // 找到环的入口
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }
    
    return slow
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 43. 重排链表 (Reorder List)
**难度：** Medium  
**题目描述：** 给定一个单链表 L 的头节点 head，单链表 L 表示为：L0 → L1 → … → Ln-1 → Ln。请将其重新排列为：L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …

**解题思路：**
1. 找到链表中点
2. 反转后半部分
3. 合并两个链表

```go
func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }
    
    // 找到中点
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    // 反转后半部分
    second := reverseList(slow.Next)
    slow.Next = nil
    
    // 合并两个链表
    first := head
    for second != nil {
        temp1 := first.Next
        temp2 := second.Next
        first.Next = second
        second.Next = temp1
        first = temp1
        second = temp2
    }
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

### 44. 二叉树的前序遍历 (Binary Tree Preorder Traversal)
**难度：** Easy  
**题目描述：** 给你二叉树的根节点 root，返回它节点值的前序遍历。

**解题思路：**
使用递归或迭代方法进行前序遍历。

```go
func preorderTraversal(root *TreeNode) []int {
    var result []int
    var preorder func(node *TreeNode)
    
    preorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        result = append(result, node.Val)
        preorder(node.Left)
        preorder(node.Right)
    }
    
    preorder(root)
    return result
}

// 迭代版本
func preorderTraversalIterative(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    var result []int
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, node.Val)
        
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 45. 二叉树的中序遍历 (Binary Tree Inorder Traversal)
**难度：** Easy  
**题目描述：** 给定一个二叉树的根节点 root，返回它的中序遍历。

**解题思路：**
使用递归或迭代方法进行中序遍历。

```go
func inorderTraversal(root *TreeNode) []int {
    var result []int
    var inorder func(node *TreeNode)
    
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        result = append(result, node.Val)
        inorder(node.Right)
    }
    
    inorder(root)
    return result
}

// 迭代版本
func inorderTraversalIterative(root *TreeNode) []int {
    var result []int
    stack := []*TreeNode{}
    curr := root
    
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, curr.Val)
        curr = curr.Right
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 46. 二叉树的后序遍历 (Binary Tree Postorder Traversal)
**难度：** Easy  
**题目描述：** 给你一棵二叉树的根节点 root，返回其节点值的后序遍历。

**解题思路：**
使用递归或迭代方法进行后序遍历。

```go
func postorderTraversal(root *TreeNode) []int {
    var result []int
    var postorder func(node *TreeNode)
    
    postorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        postorder(node.Left)
        postorder(node.Right)
        result = append(result, node.Val)
    }
    
    postorder(root)
    return result
}

// 迭代版本
func postorderTraversalIterative(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    
    var result []int
    stack := []*TreeNode{root}
    
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append([]int{node.Val}, result...)
        
        if node.Left != nil {
            stack = append(stack, node.Left)
        }
        if node.Right != nil {
            stack = append(stack, node.Right)
        }
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 47. 二叉树的层序遍历 (Binary Tree Level Order Traversal)
**难度：** Medium  
**题目描述：** 给你二叉树的根节点 root，返回其节点值的层序遍历。

**解题思路：**
使用队列进行广度优先搜索。

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        levelSize := len(queue)
        level := make([]int, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            level[i] = node.Val
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, level)
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(w)

---

### 48. 二叉树的锯齿形层序遍历 (Binary Tree Zigzag Level Order Traversal)
**难度：** Medium  
**题目描述：** 给你二叉树的根节点 root，返回其节点值的锯齿形层序遍历。

**解题思路：**
使用队列进行层序遍历，根据层数决定是否反转当前层的结果。

```go
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    var result [][]int
    queue := []*TreeNode{root}
    leftToRight := true
    
    for len(queue) > 0 {
        levelSize := len(queue)
        level := make([]int, levelSize)
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if leftToRight {
                level[i] = node.Val
            } else {
                level[levelSize-1-i] = node.Val
            }
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        
        result = append(result, level)
        leftToRight = !leftToRight
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(w)

---

### 49. 二叉树的最大深度 (Maximum Depth of Binary Tree)
**难度：** Easy  
**题目描述：** 给定一个二叉树，找出其最大深度。

**解题思路：**
使用递归或迭代方法计算最大深度。

```go
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    
    return max(leftDepth, rightDepth) + 1
}

// 迭代版本
func maxDepthIterative(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    queue := []*TreeNode{root}
    depth := 0
    
    for len(queue) > 0 {
        levelSize := len(queue)
        depth++
        
        for i := 0; i < levelSize; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
    }
    
    return depth
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 50. 从前序与中序遍历序列构造二叉树 (Construct Binary Tree from Preorder and Inorder Traversal)
**难度：** Medium  
**题目描述：** 给定两个整数数组 preorder 和 inorder，其中 preorder 是二叉树的先序遍历，inorder 是同一棵树的中序遍历，请构造二叉树并返回其根节点。

**解题思路：**
使用递归，根据前序遍历确定根节点，根据中序遍历确定左右子树的范围。

```go
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    
    // 创建中序遍历的索引映射
    inorderMap := make(map[int]int)
    for i, val := range inorder {
        inorderMap[val] = i
    }
    
    var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
    build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
        if preStart > preEnd {
            return nil
        }
        
        rootVal := preorder[preStart]
        root := &TreeNode{Val: rootVal}
        
        rootIndex := inorderMap[rootVal]
        leftSize := rootIndex - inStart
        
        root.Left = build(preStart+1, preStart+leftSize, inStart, rootIndex-1)
        root.Right = build(preStart+leftSize+1, preEnd, rootIndex+1, inEnd)
        
        return root
    }
    
    return build(0, len(preorder)-1, 0, len(inorder)-1)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 51. 从中序与后序遍历序列构造二叉树 (Construct Binary Tree from Inorder and Postorder Traversal)
**难度：** Medium  
**题目描述：** 给定两个整数数组 inorder 和 postorder，其中 inorder 是二叉树的中序遍历，postorder 是同一棵树的后序遍历，请你构造并返回这颗二叉树。

**解题思路：**
使用递归，根据后序遍历确定根节点，根据中序遍历确定左右子树的范围。

```go
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    
    // 创建中序遍历的索引映射
    inorderMap := make(map[int]int)
    for i, val := range inorder {
        inorderMap[val] = i
    }
    
    var build func(inStart, inEnd, postStart, postEnd int) *TreeNode
    build = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
        if inStart > inEnd {
            return nil
        }
        
        rootVal := postorder[postEnd]
        root := &TreeNode{Val: rootVal}
        
        rootIndex := inorderMap[rootVal]
        leftSize := rootIndex - inStart
        
        root.Left = build(inStart, rootIndex-1, postStart, postStart+leftSize-1)
        root.Right = build(rootIndex+1, inEnd, postStart+leftSize, postEnd-1)
        
        return root
    }
    
    return build(0, len(inorder)-1, 0, len(postorder)-1)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 52. 二叉树的最近公共祖先 (Lowest Common Ancestor of a Binary Tree)
**难度：** Medium  
**题目描述：** 给定一个二叉树，找到该树中两个指定节点的最近公共祖先。

**解题思路：**
使用递归，如果当前节点是p或q，或者左右子树分别包含p和q，则当前节点是LCA。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {
        return root
    }
    
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    
    if left != nil && right != nil {
        return root
    }
    
    if left != nil {
        return left
    }
    
    return right
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 53. 二叉树的序列化与反序列化 (Serialize and Deserialize Binary Tree)
**难度：** Hard  
**题目描述：** 序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，而且也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

**解题思路：**
使用前序遍历进行序列化和反序列化。

```go
type Codec struct{}

func Constructor() Codec {
    return Codec{}
}

// 序列化
func (this *Codec) serialize(root *TreeNode) string {
    var result []string
    var preorder func(node *TreeNode)
    
    preorder = func(node *TreeNode) {
        if node == nil {
            result = append(result, "null")
            return
        }
        result = append(result, strconv.Itoa(node.Val))
        preorder(node.Left)
        preorder(node.Right)
    }
    
    preorder(root)
    return strings.Join(result, ",")
}

// 反序列化
func (this *Codec) deserialize(data string) *TreeNode {
    values := strings.Split(data, ",")
    index := 0
    
    var build func() *TreeNode
    build = func() *TreeNode {
        if index >= len(values) || values[index] == "null" {
            index++
            return nil
        }
        
        val, _ := strconv.Atoi(values[index])
        index++
        
        node := &TreeNode{Val: val}
        node.Left = build()
        node.Right = build()
        
        return node
    }
    
    return build()
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 54. 路径总和 (Path Sum)
**难度：** Easy  
**题目描述：** 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum。判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和 targetSum。

**解题思路：**
使用递归，从根节点开始，每次减去当前节点的值，直到叶子节点。

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }
    
    return hasPathSum(root.Left, targetSum-root.Val) || 
           hasPathSum(root.Right, targetSum-root.Val)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 55. 路径总和 II (Path Sum II)
**难度：** Medium  
**题目描述：** 给你二叉树的根节点 root 和一个整数目标和 targetSum，找出所有从根节点到叶子节点路径总和等于给定目标和的路径。

**解题思路：**
使用回溯算法，记录当前路径，当到达叶子节点且路径和等于目标时，将路径加入结果。

```go
func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    var path []int
    
    var dfs func(node *TreeNode, sum int)
    dfs = func(node *TreeNode, sum int) {
        if node == nil {
            return
        }
        
        path = append(path, node.Val)
        sum += node.Val
        
        if node.Left == nil && node.Right == nil && sum == targetSum {
            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)
        }
        
        dfs(node.Left, sum)
        dfs(node.Right, sum)
        
        path = path[:len(path)-1]
    }
    
    dfs(root, 0)
    return result
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(h)

---

### 56. 路径总和 III (Path Sum III)
**难度：** Medium  
**题目描述：** 给定一个二叉树的根节点 root，和一个整数 targetSum，求该二叉树里节点值之和等于 targetSum 的路径的数目。

**解题思路：**
使用前缀和，记录从根节点到当前节点的路径和，查找是否存在前缀和等于当前和减去目标值。

```go
func pathSum(root *TreeNode, targetSum int) int {
    prefixSum := make(map[int]int)
    prefixSum[0] = 1
    
    var dfs func(node *TreeNode, currSum int) int
    dfs = func(node *TreeNode, currSum int) int {
        if node == nil {
            return 0
        }
        
        currSum += node.Val
        count := prefixSum[currSum-targetSum]
        
        prefixSum[currSum]++
        count += dfs(node.Left, currSum)
        count += dfs(node.Right, currSum)
        prefixSum[currSum]--
        
        return count
    }
    
    return dfs(root, 0)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 57. 二叉搜索树中的搜索 (Search in a Binary Search Tree)
**难度：** Easy  
**题目描述：** 给定二叉搜索树（BST）的根节点和一个值。你需要在BST中找到节点值等于给定值的节点。返回以该节点为根的子树。

**解题思路：**
利用二叉搜索树的性质，如果目标值小于当前节点值，搜索左子树；如果大于，搜索右子树。

```go
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    
    if val < root.Val {
        return searchBST(root.Left, val)
    }
    
    return searchBST(root.Right, val)
}

// 迭代版本
func searchBSTIterative(root *TreeNode, val int) *TreeNode {
    curr := root
    
    for curr != nil {
        if curr.Val == val {
            return curr
        } else if val < curr.Val {
            curr = curr.Left
        } else {
            curr = curr.Right
        }
    }
    
    return nil
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 58. 二叉搜索树中的插入操作 (Insert into a Binary Search Tree)
**难度：** Medium  
**题目描述：** 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。返回插入后二叉搜索树的根节点。

**解题思路：**
根据二叉搜索树的性质，找到合适的插入位置。

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    
    return root
}

// 迭代版本
func insertIntoBSTIterative(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    
    curr := root
    for true {
        if val < curr.Val {
            if curr.Left == nil {
                curr.Left = &TreeNode{Val: val}
                break
            }
            curr = curr.Left
        } else {
            if curr.Right == nil {
                curr.Right = &TreeNode{Val: val}
                break
            }
            curr = curr.Right
        }
    }
    
    return root
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 59. 删除二叉搜索树中的节点 (Delete Node in a BST)
**难度：** Medium  
**题目描述：** 给定一个二叉搜索树的根节点 root 和一个值 key，删除BST中对应的 key 的节点，并保证BST的性质不变。

**解题思路：**
根据要删除节点的子节点情况，采用不同的删除策略。

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // 要删除的节点
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        
        // 有两个子节点，找到右子树的最小值
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }
    
    return root
}

func findMin(node *TreeNode) *TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 60. 验证二叉搜索树 (Validate Binary Search Tree)
**难度：** Medium  
**题目描述：** 给你一个二叉树的根节点 root，判断其是否是一个有效的二叉搜索树。

**解题思路：**
使用中序遍历，检查序列是否严格递增。

```go
func isValidBST(root *TreeNode) bool {
    var prev *int
    
    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }
        
        if !inorder(node.Left) {
            return false
        }
        
        if prev != nil && node.Val <= *prev {
            return false
        }
        prev = &node.Val
        
        return inorder(node.Right)
    }
    
    return inorder(root)
}

// 递归版本
func isValidBSTRecursive(root *TreeNode) bool {
    return validate(root, nil, nil)
}

func validate(node *TreeNode, min, max *int) bool {
    if node == nil {
        return true
    }
    
    if (min != nil && node.Val <= *min) || (max != nil && node.Val >= *max) {
        return false
    }
    
    return validate(node.Left, min, &node.Val) && 
           validate(node.Right, &node.Val, max)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

## 61-80题

### 61. 对称二叉树 (Symmetric Tree)
**难度：** Easy  
**题目描述：** 给你一个二叉树的根节点 root，检查它是否轴对称。

**解题思路：**
使用递归，比较左右子树是否镜像对称。

```go
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isMirror(root.Left, root.Right)
}

func isMirror(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    return left.Val == right.Val && 
           isMirror(left.Left, right.Right) && 
           isMirror(left.Right, right.Left)
}

// 迭代版本
func isSymmetricIterative(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    queue := []*TreeNode{root.Left, root.Right}
    
    for len(queue) > 0 {
        left := queue[0]
        right := queue[1]
        queue = queue[2:]
        
        if left == nil && right == nil {
            continue
        }
        if left == nil || right == nil || left.Val != right.Val {
            return false
        }
        
        queue = append(queue, left.Left, right.Right, left.Right, right.Left)
    }
    
    return true
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 62. 翻转二叉树 (Invert Binary Tree)
**难度：** Easy  
**题目描述：** 给你一棵二叉树的根节点 root，翻转这棵二叉树，并返回其根节点。

**解题思路：**
使用递归，交换每个节点的左右子树。

```go
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    // 交换左右子树
    root.Left, root.Right = root.Right, root.Left
    
    // 递归翻转左右子树
    invertTree(root.Left)
    invertTree(root.Right)
    
    return root
}

// 迭代版本
func invertTreeIterative(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        // 交换左右子树
        node.Left, node.Right = node.Right, node.Left
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
    
    return root
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 63. 二叉树的直径 (Diameter of Binary Tree)
**难度：** Easy  
**题目描述：** 给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。

**解题思路：**
使用递归，计算每个节点的左右子树高度之和，同时更新最大直径。

```go
func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    
    var depth func(node *TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftDepth := depth(node.Left)
        rightDepth := depth(node.Right)
        
        // 更新最大直径
        maxDiameter = max(maxDiameter, leftDepth+rightDepth)
        
        return max(leftDepth, rightDepth) + 1
    }
    
    depth(root)
    return maxDiameter
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 64. 合并二叉树 (Merge Two Binary Trees)
**难度：** Easy  
**题目描述：** 给你两棵二叉树：root1 和 root2。想象一下，当你将其中一棵覆盖到另一棵上时，两棵树上的一些节点会重叠（而另一些不会）。

**解题思路：**
使用递归，如果两个节点都存在则相加，否则返回非空的节点。

```go
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }
    
    root1.Val += root2.Val
    root1.Left = mergeTrees(root1.Left, root2.Left)
    root1.Right = mergeTrees(root1.Right, root2.Right)
    
    return root1
}
```

**时间复杂度：** O(min(m,n))  
**空间复杂度：** O(min(m,n))

---

### 65. 二叉树的坡度 (Binary Tree Tilt)
**难度：** Easy  
**题目描述：** 给你一个二叉树的根节点 root，计算整个树的坡度。

**解题思路：**
使用递归，计算每个节点的坡度（左右子树和的差的绝对值），并累加所有节点的坡度。

```go
func findTilt(root *TreeNode) int {
    totalTilt := 0
    
    var sum func(node *TreeNode) int
    sum = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftSum := sum(node.Left)
        rightSum := sum(node.Right)
        
        tilt := abs(leftSum - rightSum)
        totalTilt += tilt
        
        return node.Val + leftSum + rightSum
    }
    
    sum(root)
    return totalTilt
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 66. 二叉搜索树的最小绝对差 (Minimum Absolute Difference in BST)
**难度：** Easy  
**题目描述：** 给你一个二叉搜索树的根节点 root，返回树中任意两不同节点值之间的最小差值。

**解题思路：**
使用中序遍历，因为二叉搜索树的中序遍历是有序的，相邻节点的差值最小。

```go
func getMinimumDifference(root *TreeNode) int {
    minDiff := math.MaxInt32
    var prev *int
    
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        inorder(node.Left)
        
        if prev != nil {
            diff := node.Val - *prev
            if diff < minDiff {
                minDiff = diff
            }
        }
        prev = &node.Val
        
        inorder(node.Right)
    }
    
    inorder(root)
    return minDiff
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 67. 二叉搜索树中的众数 (Find Mode in Binary Search Tree)
**难度：** Easy  
**题目描述：** 给你一个含重复值的二叉搜索树（BST）的根节点 root，找出并返回 BST 中的所有众数（即，出现频率最高的元素）。

**解题思路：**
使用中序遍历，统计每个值的出现次数，找出出现次数最多的值。

```go
func findMode(root *TreeNode) []int {
    var result []int
    var prev *int
    count, maxCount := 0, 0
    
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        inorder(node.Left)
        
        if prev != nil && *prev == node.Val {
            count++
        } else {
            count = 1
        }
        
        if count > maxCount {
            maxCount = count
            result = []int{node.Val}
        } else if count == maxCount {
            result = append(result, node.Val)
        }
        
        prev = &node.Val
        inorder(node.Right)
    }
    
    inorder(root)
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 68. 二叉搜索树中的第K小元素 (Kth Smallest Element in a BST)
**难度：** Medium  
**题目描述：** 给定一个二叉搜索树的根节点 root，和一个整数 k，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。

**解题思路：**
使用中序遍历，因为二叉搜索树的中序遍历是有序的，第k个访问的节点就是第k小的元素。

```go
func kthSmallest(root *TreeNode, k int) int {
    var result int
    count := 0
    
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        inorder(node.Left)
        
        count++
        if count == k {
            result = node.Val
            return
        }
        
        inorder(node.Right)
    }
    
    inorder(root)
    return result
}

// 迭代版本
func kthSmallestIterative(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    curr := root
    
    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        k--
        if k == 0 {
            return curr.Val
        }
        
        curr = curr.Right
    }
    
    return -1
}
```

**时间复杂度：** O(h + k)  
**空间复杂度：** O(h)

---

### 69. 二叉搜索树的最近公共祖先 (Lowest Common Ancestor of a Binary Search Tree)
**难度：** Easy  
**题目描述：** 给定一个二叉搜索树，找到该树中两个指定节点的最近公共祖先。

**解题思路：**
利用二叉搜索树的性质，如果两个节点都小于当前节点，则在左子树中查找；如果都大于，则在右子树中查找；否则当前节点就是LCA。

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    for root != nil {
        if p.Val < root.Val && q.Val < root.Val {
            root = root.Left
        } else if p.Val > root.Val && q.Val > root.Val {
            root = root.Right
        } else {
            return root
        }
    }
    return nil
}

// 递归版本
func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestorRecursive(root.Left, p, q)
    }
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestorRecursive(root.Right, p, q)
    }
    return root
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 70. 将有序数组转换为二叉搜索树 (Convert Sorted Array to Binary Search Tree)
**难度：** Easy  
**题目描述：** 给你一个整数数组 nums，其中元素已经按升序排列，请你将其转换为一棵高度平衡二叉搜索树。

**解题思路：**
使用递归，每次选择数组中间的元素作为根节点，然后递归构造左右子树。

```go
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    mid := len(nums) / 2
    root := &TreeNode{Val: nums[mid]}
    
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid+1:])
    
    return root
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(log n)

---

### 71. 平衡二叉树 (Balanced Binary Tree)
**难度：** Easy  
**题目描述：** 给定一个二叉树，判断它是否是高度平衡的二叉树。

**解题思路：**
使用递归，计算每个节点的左右子树高度差，如果超过1则不平衡。

```go
func isBalanced(root *TreeNode) bool {
    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        leftHeight := height(node.Left)
        rightHeight := height(node.Right)
        
        if leftHeight == -1 || rightHeight == -1 || 
           abs(leftHeight-rightHeight) > 1 {
            return -1
        }
        
        return max(leftHeight, rightHeight) + 1
    }
    
    return height(root) != -1
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 72. 路径总和 (Path Sum)
**难度：** Easy  
**题目描述：** 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum。判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和 targetSum。

**解题思路：**
使用递归，从根节点开始，每次减去当前节点的值，直到叶子节点。

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if root.Left == nil && root.Right == nil {
        return root.Val == targetSum
    }
    
    return hasPathSum(root.Left, targetSum-root.Val) || 
           hasPathSum(root.Right, targetSum-root.Val)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 73. 路径总和 II (Path Sum II)
**难度：** Medium  
**题目描述：** 给你二叉树的根节点 root 和一个整数目标和 targetSum，找出所有从根节点到叶子节点路径总和等于给定目标和的路径。

**解题思路：**
使用回溯算法，记录当前路径，当到达叶子节点且路径和等于目标时，将路径加入结果。

```go
func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    var path []int
    
    var dfs func(node *TreeNode, sum int)
    dfs = func(node *TreeNode, sum int) {
        if node == nil {
            return
        }
        
        path = append(path, node.Val)
        sum += node.Val
        
        if node.Left == nil && node.Right == nil && sum == targetSum {
            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)
        }
        
        dfs(node.Left, sum)
        dfs(node.Right, sum)
        
        path = path[:len(path)-1]
    }
    
    dfs(root, 0)
    return result
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(h)

---

### 74. 路径总和 III (Path Sum III)
**难度：** Medium  
**题目描述：** 给定一个二叉树的根节点 root，和一个整数 targetSum，求该二叉树里节点值之和等于 targetSum 的路径的数目。

**解题思路：**
使用前缀和，记录从根节点到当前节点的路径和，查找是否存在前缀和等于当前和减去目标值。

```go
func pathSum(root *TreeNode, targetSum int) int {
    prefixSum := make(map[int]int)
    prefixSum[0] = 1
    
    var dfs func(node *TreeNode, currSum int) int
    dfs = func(node *TreeNode, currSum int) int {
        if node == nil {
            return 0
        }
        
        currSum += node.Val
        count := prefixSum[currSum-targetSum]
        
        prefixSum[currSum]++
        count += dfs(node.Left, currSum)
        count += dfs(node.Right, currSum)
        prefixSum[currSum]--
        
        return count
    }
    
    return dfs(root, 0)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 75. 二叉搜索树中的搜索 (Search in a Binary Search Tree)
**难度：** Easy  
**题目描述：** 给定二叉搜索树（BST）的根节点和一个值。你需要在BST中找到节点值等于给定值的节点。返回以该节点为根的子树。

**解题思路：**
利用二叉搜索树的性质，如果目标值小于当前节点值，搜索左子树；如果大于，搜索右子树。

```go
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    
    if val < root.Val {
        return searchBST(root.Left, val)
    }
    
    return searchBST(root.Right, val)
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 76. 二叉搜索树中的插入操作 (Insert into a Binary Search Tree)
**难度：** Medium  
**题目描述：** 给定二叉搜索树（BST）的根节点和要插入树中的值，将值插入二叉搜索树。返回插入后二叉搜索树的根节点。

**解题思路：**
根据二叉搜索树的性质，找到合适的插入位置。

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    
    return root
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 77. 删除二叉搜索树中的节点 (Delete Node in a BST)
**难度：** Medium  
**题目描述：** 给定一个二叉搜索树的根节点 root 和一个值 key，删除BST中对应的 key 的节点，并保证BST的性质不变。

**解题思路：**
根据要删除节点的子节点情况，采用不同的删除策略。

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // 要删除的节点
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        
        // 有两个子节点，找到右子树的最小值
        minNode := findMin(root.Right)
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }
    
    return root
}

func findMin(node *TreeNode) *TreeNode {
    for node.Left != nil {
        node = node.Left
    }
    return node
}
```

**时间复杂度：** O(h)  
**空间复杂度：** O(h)

---

### 78. 验证二叉搜索树 (Validate Binary Search Tree)
**难度：** Medium  
**题目描述：** 给你一个二叉树的根节点 root，判断其是否是一个有效的二叉搜索树。

**解题思路：**
使用中序遍历，检查序列是否严格递增。

```go
func isValidBST(root *TreeNode) bool {
    var prev *int
    
    var inorder func(node *TreeNode) bool
    inorder = func(node *TreeNode) bool {
        if node == nil {
            return true
        }
        
        if !inorder(node.Left) {
            return false
        }
        
        if prev != nil && node.Val <= *prev {
            return false
        }
        prev = &node.Val
        
        return inorder(node.Right)
    }
    
    return inorder(root)
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 79. 恢复二叉搜索树 (Recover Binary Search Tree)
**难度：** Medium  
**题目描述：** 给你二叉搜索树的根节点 root，该树中的恰好两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树。

**解题思路：**
使用中序遍历找到两个错误的节点，然后交换它们的值。

```go
func recoverTree(root *TreeNode) {
    var first, second, prev *TreeNode
    
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        inorder(node.Left)
        
        if prev != nil && prev.Val > node.Val {
            if first == nil {
                first = prev
            }
            second = node
        }
        prev = node
        
        inorder(node.Right)
    }
    
    inorder(root)
    
    if first != nil && second != nil {
        first.Val, second.Val = second.Val, first.Val
    }
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(h)

---

### 80. 不同的二叉搜索树 (Unique Binary Search Trees)
**难度：** Medium  
**题目描述：** 给你一个整数 n，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的二叉搜索树有多少种？

**解题思路：**
使用动态规划，dp[i]表示i个节点能组成的不同BST数量。

```go
func numTrees(n int) int {
    dp := make([]int, n+1)
    dp[0] = 1
    dp[1] = 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[j-1] * dp[i-j]
        }
    }
    
    return dp[n]
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n)

---

## 81-100题

### 81. 不同的二叉搜索树 II (Unique Binary Search Trees II)
**难度：** Medium  
**题目描述：** 给你一个整数 n，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同二叉搜索树。

**解题思路：**
使用递归，对于每个可能的根节点，递归构造左右子树，然后组合所有可能的树。

```go
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    return generate(1, n)
}

func generate(start, end int) []*TreeNode {
    if start > end {
        return []*TreeNode{nil}
    }
    
    var result []*TreeNode
    
    for i := start; i <= end; i++ {
        leftTrees := generate(start, i-1)
        rightTrees := generate(i+1, end)
        
        for _, left := range leftTrees {
            for _, right := range rightTrees {
                root := &TreeNode{Val: i}
                root.Left = left
                root.Right = right
                result = append(result, root)
            }
        }
    }
    
    return result
}
```

**时间复杂度：** O(4^n / n^(3/2))  
**空间复杂度：** O(4^n / n^(3/2))

---

### 82. 不同的子序列 (Distinct Subsequences)
**难度：** Hard  
**题目描述：** 给定一个字符串 s 和一个字符串 t，计算在 s 的子序列中 t 出现的个数。

**解题思路：**
使用动态规划，dp[i][j]表示s的前i个字符中t的前j个字符出现的次数。

```go
func numDistinct(s string, t string) int {
    m, n := len(s), len(t)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    // 空字符串是任何字符串的子序列
    for i := 0; i <= m; i++ {
        dp[i][0] = 1
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == t[j-1] {
                dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    
    return dp[m][n]
}
```

**时间复杂度：** O(m*n)  
**空间复杂度：** O(m*n)

---

### 83. 分割回文串 (Palindrome Partitioning)
**难度：** Medium  
**题目描述：** 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。返回 s 所有可能的分割方案。

**解题思路：**
使用回溯算法，尝试所有可能的分割方式，并使用动态规划预处理回文串。

```go
func partition(s string) [][]string {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    
    // 预处理回文串
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
                dp[j][i] = true
            }
        }
    }
    
    var result [][]string
    var path []string
    
    var backtrack func(start int)
    backtrack = func(start int) {
        if start == n {
            temp := make([]string, len(path))
            copy(temp, path)
            result = append(result, temp)
            return
        }
        
        for i := start; i < n; i++ {
            if dp[start][i] {
                path = append(path, s[start:i+1])
                backtrack(i + 1)
                path = path[:len(path)-1]
            }
        }
    }
    
    backtrack(0)
    return result
}
```

**时间复杂度：** O(2^n)  
**空间复杂度：** O(n)

---

### 84. 分割回文串 II (Palindrome Partitioning II)
**难度：** Hard  
**题目描述：** 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文串。返回符合要求的最少分割次数。

**解题思路：**
使用动态规划，先预处理回文串，然后计算最少分割次数。

```go
func minCut(s string) int {
    n := len(s)
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }
    
    // 预处理回文串
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
                dp[j][i] = true
            }
        }
    }
    
    // 计算最少分割次数
    cuts := make([]int, n)
    for i := 0; i < n; i++ {
        if dp[0][i] {
            cuts[i] = 0
        } else {
            cuts[i] = i
            for j := 0; j < i; j++ {
                if dp[j+1][i] {
                    cuts[i] = min(cuts[i], cuts[j]+1)
                }
            }
        }
    }
    
    return cuts[n-1]
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n²)

---

### 85. 单词拆分 (Word Break)
**难度：** Medium  
**题目描述：** 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s。

**解题思路：**
使用动态规划，dp[i]表示s的前i个字符是否可以被字典中的单词拼接。

```go
func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }
    
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[n]
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n)

---

### 86. 单词拆分 II (Word Break II)
**难度：** Hard  
**题目描述：** 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

**解题思路：**
使用回溯算法，结合动态规划优化，找出所有可能的拆分方案。

```go
func wordBreak(s string, wordDict []string) []string {
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }
    
    n := len(s)
    dp := make([]bool, n+1)
    dp[0] = true
    
    // 先判断是否可以拆分
    for i := 1; i <= n; i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    
    if !dp[n] {
        return []string{}
    }
    
    var result []string
    var path []string
    
    var backtrack func(start int)
    backtrack = func(start int) {
        if start == n {
            result = append(result, strings.Join(path, " "))
            return
        }
        
        for i := start + 1; i <= n; i++ {
            if wordSet[s[start:i]] {
                path = append(path, s[start:i])
                backtrack(i)
                path = path[:len(path)-1]
            }
        }
    }
    
    backtrack(0)
    return result
}
```

**时间复杂度：** O(2^n)  
**空间复杂度：** O(2^n)

---

### 87. 最长递增子序列 (Longest Increasing Subsequence)
**难度：** Medium  
**题目描述：** 给你一个整数数组 nums，找到其中最长严格递增子序列的长度。

**解题思路：**
使用动态规划，dp[i]表示以nums[i]结尾的最长递增子序列的长度。

```go
func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    n := len(nums)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    
    maxLen := 1
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        maxLen = max(maxLen, dp[i])
    }
    
    return maxLen
}

// 优化版本：使用二分查找
func lengthOfLISOptimized(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    tails := []int{nums[0]}
    
    for i := 1; i < len(nums); i++ {
        if nums[i] > tails[len(tails)-1] {
            tails = append(tails, nums[i])
        } else {
            // 二分查找第一个大于等于nums[i]的位置
            left, right := 0, len(tails)-1
            for left < right {
                mid := (left + right) / 2
                if tails[mid] < nums[i] {
                    left = mid + 1
                } else {
                    right = mid
                }
            }
            tails[left] = nums[i]
        }
    }
    
    return len(tails)
}
```

**时间复杂度：** O(n²) / O(n log n)  
**空间复杂度：** O(n)

---

### 88. 最长递增子序列的个数 (Number of Longest Increasing Subsequence)
**难度：** Medium  
**题目描述：** 给定一个未排序的整数数组，找到最长递增子序列的个数。

**解题思路：**
使用动态规划，同时记录长度和个数。

```go
func findNumberOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    n := len(nums)
    lengths := make([]int, n)
    counts := make([]int, n)
    
    for i := range lengths {
        lengths[i] = 1
        counts[i] = 1
    }
    
    maxLen := 1
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[j] < nums[i] {
                if lengths[j]+1 > lengths[i] {
                    lengths[i] = lengths[j] + 1
                    counts[i] = counts[j]
                } else if lengths[j]+1 == lengths[i] {
                    counts[i] += counts[j]
                }
            }
        }
        maxLen = max(maxLen, lengths[i])
    }
    
    result := 0
    for i := 0; i < n; i++ {
        if lengths[i] == maxLen {
            result += counts[i]
        }
    }
    
    return result
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n)

---

### 89. 俄罗斯套娃信封问题 (Russian Doll Envelopes)
**难度：** Hard  
**题目描述：** 给你一个二维整数数组 envelopes，其中 envelopes[i] = [wi, hi]，表示第 i 个信封的宽度和高度。

**解题思路：**
先按宽度排序，然后对高度使用最长递增子序列算法。

```go
func maxEnvelopes(envelopes [][]int) int {
    if len(envelopes) == 0 {
        return 0
    }
    
    // 按宽度升序，高度降序排序
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    
    // 对高度使用最长递增子序列
    heights := make([]int, len(envelopes))
    for i, env := range envelopes {
        heights[i] = env[1]
    }
    
    return lengthOfLIS(heights)
}

func lengthOfLIS(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    tails := []int{nums[0]}
    
    for i := 1; i < len(nums); i++ {
        if nums[i] > tails[len(tails)-1] {
            tails = append(tails, nums[i])
        } else {
            left, right := 0, len(tails)-1
            for left < right {
                mid := (left + right) / 2
                if tails[mid] < nums[i] {
                    left = mid + 1
                } else {
                    right = mid
                }
            }
            tails[left] = nums[i]
        }
    }
    
    return len(tails)
}
```

**时间复杂度：** O(n log n)  
**空间复杂度：** O(n)

---

### 90. 最长公共子序列 (Longest Common Subsequence)
**难度：** Medium  
**题目描述：** 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

**解题思路：**
使用动态规划，dp[i][j]表示text1前i个字符和text2前j个字符的最长公共子序列长度。

```go
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[m][n]
}
```

**时间复杂度：** O(m*n)  
**空间复杂度：** O(m*n)

---

### 91. 最长公共子串 (Longest Common Substring)
**难度：** Medium  
**题目描述：** 给定两个字符串，求它们的最长公共子串的长度。

**解题思路：**
使用动态规划，dp[i][j]表示以text1[i-1]和text2[j-1]结尾的最长公共子串长度。

```go
func longestCommonSubstring(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    maxLen := 0
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = dp[i-1][j-1] + 1
                maxLen = max(maxLen, dp[i][j])
            } else {
                dp[i][j] = 0
            }
        }
    }
    
    return maxLen
}
```

**时间复杂度：** O(m*n)  
**空间复杂度：** O(m*n)

---

### 92. 编辑距离 (Edit Distance)
**难度：** Hard  
**题目描述：** 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数。

**解题思路：**
使用动态规划，dp[i][j]表示word1前i个字符转换成word2前j个字符的最少操作数。

```go
func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    // 初始化
    for i := 0; i <= m; i++ {
        dp[i][0] = i
    }
    for j := 0; j <= n; j++ {
        dp[0][j] = j
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
            }
        }
    }
    
    return dp[m][n]
}
```

**时间复杂度：** O(m*n)  
**空间复杂度：** O(m*n)

---

### 93. 最长回文子序列 (Longest Palindromic Subsequence)
**难度：** Medium  
**题目描述：** 给你一个字符串 s，找出其中最长的回文子序列，并返回该序列的长度。

**解题思路：**
使用动态规划，dp[i][j]表示s[i:j+1]的最长回文子序列长度。

```go
func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    
    // 单个字符都是回文
    for i := 0; i < n; i++ {
        dp[i][i] = 1
    }
    
    // 从长度为2开始
    for length := 2; length <= n; length++ {
        for i := 0; i <= n-length; i++ {
            j := i + length - 1
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[0][n-1]
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(n²)

---

### 94. 最长回文子串 (Longest Palindromic Substring)
**难度：** Medium  
**题目描述：** 给你一个字符串 s，找到 s 中最长的回文子串。

**解题思路：**
使用中心扩展法，从每个可能的中心向两边扩展。

```go
func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }
    
    start, maxLen := 0, 1
    
    for i := 0; i < len(s); i++ {
        // 奇数长度回文
        len1 := expandAroundCenter(s, i, i)
        // 偶数长度回文
        len2 := expandAroundCenter(s, i, i+1)
        
        maxLenCurr := max(len1, len2)
        if maxLenCurr > maxLen {
            maxLen = maxLenCurr
            start = i - (maxLenCurr-1)/2
        }
    }
    
    return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}
```

**时间复杂度：** O(n²)  
**空间复杂度：** O(1)

---

### 95. 最长有效括号 (Longest Valid Parentheses)
**难度：** Hard  
**题目描述：** 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

**解题思路：**
使用动态规划，dp[i]表示以s[i]结尾的最长有效括号长度。

```go
func longestValidParentheses(s string) int {
    n := len(s)
    if n < 2 {
        return 0
    }
    
    dp := make([]int, n)
    maxLen := 0
    
    for i := 1; i < n; i++ {
        if s[i] == ')' {
            if s[i-1] == '(' {
                dp[i] = 2
                if i >= 2 {
                    dp[i] += dp[i-2]
                }
            } else if dp[i-1] > 0 {
                // 找到对应的左括号
                left := i - dp[i-1] - 1
                if left >= 0 && s[left] == '(' {
                    dp[i] = dp[i-1] + 2
                    if left > 0 {
                        dp[i] += dp[left-1]
                    }
                }
            }
            maxLen = max(maxLen, dp[i])
        }
    }
    
    return maxLen
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 96. 最小覆盖子串 (Minimum Window Substring)
**难度：** Hard  
**题目描述：** 给你一个字符串 s、一个字符串 t。返回 s 中涵盖 t 所有字符的最小子串。

**解题思路：**
使用滑动窗口，维护一个包含t中所有字符的窗口。

```go
func minWindow(s string, t string) string {
    if len(s) < len(t) {
        return ""
    }
    
    need := make(map[byte]int)
    for i := range t {
        need[t[i]]++
    }
    
    left, right := 0, 0
    valid := 0
    window := make(map[byte]int)
    
    start, length := 0, math.MaxInt32
    
    for right < len(s) {
        c := s[right]
        right++
        
        if need[c] > 0 {
            window[c]++
            if window[c] == need[c] {
                valid++
            }
        }
        
        for valid == len(need) {
            if right-left < length {
                start = left
                length = right - left
            }
            
            d := s[left]
            left++
            
            if need[d] > 0 {
                if window[d] == need[d] {
                    valid--
                }
                window[d]--
            }
        }
    }
    
    if length == math.MaxInt32 {
        return ""
    }
    return s[start : start+length]
}
```

**时间复杂度：** O(m+n)  
**空间复杂度：** O(m+n)

---

### 97. 滑动窗口最大值 (Sliding Window Maximum)
**难度：** Hard  
**题目描述：** 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只能看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

**解题思路：**
使用单调队列，维护一个递减的队列，队首元素就是当前窗口的最大值。

```go
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return []int{}
    }
    
    var result []int
    deque := []int{}
    
    for i := 0; i < len(nums); i++ {
        // 移除窗口外的元素
        for len(deque) > 0 && deque[0] <= i-k {
            deque = deque[1:]
        }
        
        // 移除比当前元素小的元素
        for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
            deque = deque[:len(deque)-1]
        }
        
        deque = append(deque, i)
        
        // 当窗口大小达到k时，记录结果
        if i >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }
    
    return result
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(k)

---

### 98. 最大矩形 (Maximal Rectangle)
**难度：** Hard  
**题目描述：** 给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

**解题思路：**
将问题转化为柱状图中最大矩形问题，对每一行计算高度数组。

```go
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return 0
    }
    
    m, n := len(matrix), len(matrix[0])
    heights := make([]int, n)
    maxArea := 0
    
    for i := 0; i < m; i++ {
        // 更新高度数组
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }
        
        // 计算当前行的最大矩形面积
        maxArea = max(maxArea, largestRectangleArea(heights))
    }
    
    return maxArea
}

func largestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }
    
    stack := []int{}
    maxArea := 0
    
    for i := 0; i <= n; i++ {
        h := 0
        if i < n {
            h = heights[i]
        }
        
        for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            
            maxArea = max(maxArea, height*width)
        }
        
        stack = append(stack, i)
    }
    
    return maxArea
}
```

**时间复杂度：** O(m*n)  
**空间复杂度：** O(n)

---

### 99. 柱状图中最大的矩形 (Largest Rectangle in Histogram)
**难度：** Hard  
**题目描述：** 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1。求在该柱状图中，能够勾勒出来的矩形的最大面积。

**解题思路：**
使用单调栈，维护一个递增的栈，当遇到比栈顶元素小的元素时，计算以栈顶元素为高的最大矩形面积。

```go
func largestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {
        return 0
    }
    
    stack := []int{}
    maxArea := 0
    
    for i := 0; i <= n; i++ {
        h := 0
        if i < n {
            h = heights[i]
        }
        
        for len(stack) > 0 && heights[stack[len(stack)-1]] > h {
            height := heights[stack[len(stack)-1]]
            stack = stack[:len(stack)-1]
            
            width := i
            if len(stack) > 0 {
                width = i - stack[len(stack)-1] - 1
            }
            
            maxArea = max(maxArea, height*width)
        }
        
        stack = append(stack, i)
    }
    
    return maxArea
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(n)

---

### 100. 接雨水 (Trapping Rain Water)
**难度：** Hard  
**题目描述：** 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

**解题思路：**
使用双指针，从两边向中间移动，维护左右两边的最大高度。

```go
func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }
    
    left, right := 0, len(height)-1
    leftMax, rightMax := 0, 0
    water := 0
    
    for left < right {
        if height[left] < height[right] {
            if height[left] >= leftMax {
                leftMax = height[left]
            } else {
                water += leftMax - height[left]
            }
            left++
        } else {
            if height[right] >= rightMax {
                rightMax = height[right]
            } else {
                water += rightMax - height[right]
            }
            right--
        }
    }
    
    return water
}
```

**时间复杂度：** O(n)  
**空间复杂度：** O(1)

---

