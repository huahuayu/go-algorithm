package problem001

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

// 解法一：暴力法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return nil
}

// 解法二：两次遍历，第一次遍历将slice转换为map，第二次遍历slice同时查询map中是否有匹配项
func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		want := target - num
		j, ok := m[want]
		if ok {
			if i != j {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 解法三：一次遍历，遍历的时候同时检查是否能和map中的key匹配
func twoSumHashImproved(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		want := target - num
		j, ok := m[want]
		if ok {
			return []int{j, i}
		}

		m[num] = i
	}

	return nil
}
