package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated string
	for range times {
		repeated += character
	}
	return repeated
}

func ToLower(char string) string {
	return strings.ToLower(char)
}