package solution

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

// 双指针法，i为快指针，j为慢指针
func moveZeroes(nums []int) {
	j := 0
	for i, v := range nums {
		if v != 0 {
			if i != j {
				nums[j] = nums[i]
			}
			j++
		}
	}

	// j < len(nums)的位数补齐0
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}
