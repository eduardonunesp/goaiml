package goaiml

import (
	"encoding/xml"
	"errors"
	"regexp"
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

func (aimlPattern *AIMLPattern) Regexify() *regexp.Regexp {
	rString := aimlPattern.Content
	rString = stringMinifier(rString)
	rString = strings.Replace(rString, "*", "(.*)", -1)
	return regexp.MustCompile("(?i)" + rString)
}

func (aimlPattern *AIMLPattern) ProcessBot(aiml *AIML) error {
	botStruct := struct {
		XMLName xml.Name `xml:"bot"`
		Name    string   `xml:"name,attr"`
	}{}

	err := xml.Unmarshal([]byte(aimlPattern.Content), &botStruct)

	if err != nil {
		return err
	}

	content, ok := aiml.Bot[botStruct.Name]

	if !ok {
		return errors.New("Key not found in memory")
	}

	aimlPattern.Content = strings.Replace(aimlPattern.Content, `<bot name="`+botStruct.Name+`"/>`, content, -1)
	return nil
}
