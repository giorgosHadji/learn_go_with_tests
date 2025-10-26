package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	t := ""
	for i := 0; i < repeatCount; i++ {
		t += character
	}
	return t
}

func RepeatFaster(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
