package prose

import "strings"

func JoinWithCommas(phrases []string) string {
	countElem := len(phrases)
	if countElem == 0 {
		return ""
	} else if countElem == 1 {
		return phrases[0]
	} else if countElem == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:countElem-1], ", ")
		result += ", and "
		result += phrases[countElem-1]
		return result
	}
}
