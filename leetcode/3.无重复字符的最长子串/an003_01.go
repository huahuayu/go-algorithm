package problem003

/*
java solution:
public class Solution {
    public int lengthOfLongestSubstring(String s) {
        int n = s.length(), ans = 0;
        Map<Character, Integer> map = new HashMap<>(); // current index of character
        // try to extend the range [i, j]
        for (int j = 0, i = 0; j < n; j++) {
            if (map.containsKey(s.charAt(j))) {
                i = Math.max(map.get(s.charAt(j)), i);
            }
            ans = Math.max(ans, j - i + 1);
            map.put(s.charAt(j), j + 1);
        }
        return ans;
    }
}
*/

// 滑动窗口
func lengthOfLongestSubstringWindows(s string) int {
	ans := 0
	m := make(map[rune]int)
	rs := []rune(s)

	for i, j := 0, 0; j < len(s); j++ {
		k := rs[j] // map key
		if _, ok := m[k]; ok {
			if i < m[k]+1 {
				i = m[k] + 1
			}
		}

		if ans < j-i+1 {
			ans = j - i + 1
		}
		m[k] = j
	}

	return ans
}

func lengthOfLongestSubstringWinner(s string) int {
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i-left+1 > maxLen {
			maxLen = i - left + 1
		}
		location[s[i]] = i
	}
	return maxLen
}
