package goaiml

import (
	"regexp"
	"strings"
)

func (aimlPattern *AIMLPattern) Regexify() *regexp.Regexp {
	rString := aimlPattern.Content
	rString = strings.Replace(rString, "*", "\\s?(.*)", -1)
	rString = strings.TrimSpace(rString)
	rString = strings.Replace(rString, "\n", "", -1)
	rString = stringMinifier(rString)
	return regexp.MustCompile("(?i)" + rString)
}
