package goaiml

import (
	"strings"
	"unicode"
)

func stringMinifier(in string) (out string) {
	white := false
	for _, c := range in {
		if unicode.IsSpace(c) {
			if !white {
				out = out + " "
			}
			white = true
		} else {
			out = out + string(c)
			white = false
		}
	}
	return
}

func postFormatInput(input string) string {
	return strings.TrimSpace(stringMinifier(strings.Replace(input, "\n", "", -1)))
}

func preFormatInput(input string) string {
	return " " + strings.TrimSpace(strings.Replace(input, "\n", "", -1)) + " "
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func reverseStrings(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseStrings(input[1:]), input[0])
}
