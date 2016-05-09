package goaiml

import (
	"regexp"
	"strings"
)

func (aimlPattern *AIMLPattern) Regexify() error {
	rString := aimlPattern.Content
	rString = strings.Replace(rString, "*", "\\s?(.*)", -1)
	rString = strings.Replace(rString, "_", "\\s?(.*)", -1)
	rString = strings.TrimSpace(rString)
	rString = strings.Replace(rString, "\n", "", -1)
	rString = stringMinifier(rString)

	var err error = nil
	aimlPattern.Re, err = regexp.Compile("(?i)" + rString)

	if err != nil {
		return err
	}

	return nil
}
