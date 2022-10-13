package awesome

func longestCommonPrefix(strs []string) string {


	if len(strs) < 1 {
		return ""
	}

	checkSame := func(v byte, index int) bool {
		for _, s := range strs {
			if index >= len(s) {
				return false
			}
			if s[index] != v {
				return false
			}
		}
		return true
	}

	index := 0
	for ;; {
		if index >= len(strs[0]) {
			break
		}

		if !checkSame(strs[0][index], index) {
			break
		}
		index ++
	}

	if index == 0 {
		return ""
	}
	return strs[0][:index]

}

