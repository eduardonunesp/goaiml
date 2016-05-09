package goaiml

import (
	"regexp"
	"strings"
)

func (aimlPattern *AIMLPattern) Regexify() *regexp.Regexp {
	rString := aimlPattern.Content
	rString = strings.Replace(rString, "*", "\\s?(.*)", -1)
	return regexp.MustCompile("(?i)" + rString)
}
