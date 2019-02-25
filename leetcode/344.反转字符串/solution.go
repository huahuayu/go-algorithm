package solution

func reverseString(s []byte) {

	for i, j := 0, len(s)-1; ; {
		if j > i {
			s[i], s[j] = s[j], s[i]
		} else {
			break
		}
		i++
		j--
	}

}
