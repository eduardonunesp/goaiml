package goaiml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func (aiml *AIMLInterpreter) ProcessSetTag(content string) (string, error) {
	ret := ""
	setStruct := struct {
		XMLName xml.Name `xml:"set"`
		Name    string   `xml:"name,attr"`
		Content string   `xml:",innerxml"`
	}{}

	err := xml.Unmarshal([]byte(content), &setStruct)

	if err != nil {
		return ret, err
	}

	ret = content
	ret = strings.Replace(ret, `<set name="`+setStruct.Name+`">`, "", -1)
	ret = strings.Replace(ret, `</set>`, "", -1)

	aiml.Memory[setStruct.Name] = setStruct.Content

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessGetTag(content string) (string, error) {
	ret := ""
	getStruct := struct {
		XMLName xml.Name `xml:"get"`
		Name    string   `xml:"name,attr"`
	}{}

	err := xml.Unmarshal([]byte(content), &getStruct)

	if err != nil {
		return ret, err
	}

	value, ok := aiml.Memory[getStruct.Name]

	if !ok {
		return ret, errors.New("Key not found in memory")
	}

	ret = content
	ret = strings.Replace(ret, `<get name="`+getStruct.Name+`" />`, value, -1)
	ret = strings.Replace(ret, `<get name="`+getStruct.Name+`"/>`, value, -1)

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessBotTag(content string) (string, error) {
	ret := ""
	botStruct := struct {
		XMLName xml.Name `xml:"bot"`
		Name    string   `xml:"name,attr"`
	}{}

	err := xml.Unmarshal([]byte(content), &botStruct)

	if err != nil {
		return ret, err
	}

	value, ok := aiml.Bot[botStruct.Name]

	if !ok {
		return ret, errors.New("Key not found in bot")
	}

	ret = content
	ret = strings.Replace(ret, `<bot name="`+botStruct.Name+`" />`, value, -1)
	ret = strings.Replace(ret, `<bot name="`+botStruct.Name+`"/>`, value, -1)

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessStarTag(content string, starContent []string) string {
	ret := content
	for idx, sContent := range starContent {
		if idx > 0 {
			ret = strings.Replace(ret, "<star />", strings.TrimSpace(sContent), 1)
			ret = strings.Replace(ret, "<star/>", strings.TrimSpace(sContent), 1)
		}
	}

	ret = strings.Replace(ret, "<star />", "", -1)
	ret = strings.Replace(ret, "<star/>", "", -1)
	return ret
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (aiml *AIMLInterpreter) ProcessRandomTag(content string, starContent []string) (string, error) {
	ret := ""
	randomStruct := struct {
		XMLName xml.Name `xml:"random"`
		List    []struct {
			XMLName xml.Name `xml:"li"`
			Content string   `xml:",innerxml"`
		} `xml:"li"`
	}{}

	err := xml.Unmarshal([]byte(content), &randomStruct)

	if err != nil {
		return ret, err
	}

	randIdx := random(0, len(randomStruct.List))
	randContent := randomStruct.List[randIdx]
	li := randContent.Content

	ret, errT := aiml.processAllTemplateTags(li, []string{}, true)

	if errT != nil {
		return ret, errT
	}

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessThinkTag(content string, starContent []string) (string, error) {
	ret := ""
	thinkStruct := struct {
		XMLName xml.Name `xml:"think"`
		Content string   `xml:",innerxml"`
	}{}

	err := xml.Unmarshal([]byte(content), &thinkStruct)

	if err != nil {
		return ret, err
	}

	tmp := content
	tmp = strings.Replace(tmp, `<think>`, "", -1)
	tmp = strings.Replace(tmp, `</think>`, "", -1)

	_, errPattern := aiml.processBasicTemplateTags(tmp, starContent)

	if errPattern != nil {
		return tmp, errPattern
	}

	ret = regexp.MustCompile("(?s)<think>.*</think>").ReplaceAllString(content, "")
	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessSraiTag(content string) (string, error) {
	ret := ""
	sraiStruct := struct {
		XMLName xml.Name `xml:"srai"`
		Content string   `xml:",innerxml"`
	}{}

	err := xml.Unmarshal([]byte(content), &sraiStruct)

	if err != nil {
		return ret, err
	}

	sraiStruct.Content = strings.Replace(sraiStruct.Content, `<srai>`, "", -1)
	sraiStruct.Content = strings.Replace(sraiStruct.Content, `</srai>`, "", -1)

	ret, errPattern := aiml.findPattern(sraiStruct.Content, true)

	if errPattern != nil {
		return ret, errPattern
	}

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessConditionTag(content string) (string, error) {
	ret := ""
	conditionStruct := struct {
		XMLName xml.Name `xml:"condition"`
		List    []struct {
			XMLName xml.Name `xml:"li"`
			Name    string   `xml:"name,attr"`
			Value   string   `xml:"value,attr"`
			Content string   `xml:",innerxml"`
		} `xml:"li"`
	}{}

	err := xml.Unmarshal([]byte(content), &conditionStruct)

	if err != nil {
		return ret, err
	}

	for _, li := range conditionStruct.List {
		memValue, ok := aiml.Memory[li.Name]
		if ok {
			if li.Value == "" {
				ret = li.Content
				break
			} else if memValue == li.Value {
				ret = li.Content
				break
			}
		} else {
			ret = li.Content
		}
	}

	ret, errT := aiml.processAllTemplateTags(ret, []string{}, true)

	if errT != nil {
		return ret, errT
	}

	return ret, nil
}

func (aiml *AIMLInterpreter) ProcessLearnTag(content string) (string, error) {
	ret := ""
	learnStruct := struct {
		Name xml.Name `xml:"learns"`
		List []struct {
			XMLName xml.Name `xml:"learn"`
			Content string   `xml:",innerxml"`
		} `xml:"learn"`
	}{}

	tmpContent := fmt.Sprintf("<learns>%s</learns>", content)
	err := xml.Unmarshal([]byte(tmpContent), &learnStruct)

	if err != nil {
		return ret, err
	}

	for _, learnFile := range learnStruct.List {
		errT := aiml.learnFromMoreXML(learnFile.Content)
		if errT != nil {
			return ret, errT
		}
	}

	ret = regexp.MustCompile("(?s)<learn>.*</learn>").ReplaceAllString(content, "")
	return ret, nil
}
