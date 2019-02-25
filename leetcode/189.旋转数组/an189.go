package problem189

// 解法一，map暂存
func rotate(nums []int, k int) {
	if k == 0 || nums == nil {
		return
	}
	n := len(nums)
	if k > n {
		k = k % n
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[i] = nums[i]

		if k > i {
			nums[i] = nums[n-k+i]
		} else {
			nums[i] = m[i-k]
		}
	}
}

// 解法二，数组暂存, 易于理解
func rotateSlice(nums []int, k int) {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[(i+k)%len(nums)] = nums[i] // i+k没超过len的下标加k，超过len的溢出
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}

// 解法三，翻转
func rotateWinner(nums []int, k int) {
	if k == 0 || nums == nil {
		return
	}

	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	return
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
