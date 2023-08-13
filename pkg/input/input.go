package input

import (
	"regexp"
	"strings"
)

func RemoveWithPattern(value, pattern string) string {
	tmpvalue := strings.ReplaceAll(value, " ", "")
	if len(tmpvalue) == 0 {
		return tmpvalue
	}
	return regexp.MustCompile(pattern).ReplaceAllString(tmpvalue, "")
}
