package util

import "strings"

// functions to assist in parsing strings
func FirstWord(s string) string {
	i := strings.IndexByte(s, " "[0])
	if i == -1 {
		return s
	} else {
		return s[:i]
	}
}


func ConsumeSpaces(p *string) {
	i := 0
	for i <len(*p) {
		if (*p)[i] != ' '{
			break
		}
		i++
	}
	*p = (*p)[i:]
}