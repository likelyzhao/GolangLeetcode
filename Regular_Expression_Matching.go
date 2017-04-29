package leecode

//Regular Expression Matching

func search(s string, p string, idx int) bool {
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		}
		return false

	}
	if len(s) == 0 {
		return false
	}
	switch s[0] {
	case '*':
		if len(s) == 1 {
			return true
		}
		for i := 1; i < len(p); i++ {
			if p[i] == s[1] {
				return search(s[1:len(s)], p[i:len(p)], idx)
			}
		}
		return false
	case '.':
		return search(s[1:len(s)], p[1:len(p)], idx)
	case p[0]:
		return search(s[1:len(s)], p[1:len(p)], idx)
	default:
		return false

	}

}

func IsMatch(s string, p string) bool {
	if len(s) == 0 {
		return len(p) == 0
	}
	if len(p) < len(s) {
		return false
	}
	for i := 0; i < len(p); i++ {
		if p[i] != s[0] && p[i] != '.' {
			continue
		}
		if search(p[i:len(p)], s, i) {
			return true
		}
	}
	return false

}
