package solution

func intersect(nums1 []int, nums2 []int) []int {

	var result []int
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		m1[v] += 1
	}

	for _, v := range nums2 {
		m2[v] += 1
	}

	// 使用短的map做基准，查询长map中是否有相同的值
	if len(nums1) < len(nums2) {
		for k1, v1 := range m1 {
			v2, ok := m2[k1]
			if ok {
				min := 0
				if v1 > v2 {
					min = v2
				} else {
					min = v1
				}
				for i := 0; i < min; i++ {
					result = append(result, k1)
				}
			}
		}
	} else {
		for k2, v2 := range m2 {
			v1, ok := m1[k2]
			if ok {
				min := 0
				if v1 > v2 {
					min = v2
				} else {
					min = v1
				}
				for i := 0; i < min; i++ {
					result = append(result, k2)
				}
			}
		}
	}

	return result

}

func intersectWinner(nums1 []int, nums2 []int) []int {
	var result []int
	var short, long []int
	m := make(map[int]int)

	if len(nums1) < len(nums2) {
		short, long = nums1, nums2
	} else {
		short, long = nums2, nums1
	}

	for _, v := range short {
		m[v]++
	}

	for _, v := range long {
		times, ok := m[v]

		if ok && times > 0 {
			result = append(result, v)
			m[v]--
		}
	}

	return result
}
