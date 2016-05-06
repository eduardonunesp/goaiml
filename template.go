package goaiml

import (
	"encoding/xml"
	"errors"
	"math/rand"
	"strings"
	"time"
)

func (aimlTemplate *AIMLTemplate) ProcessSet(aiml *AIML) error {
	setStruct := struct {
		XMLName xml.Name `xml:"set"`
		Name    string   `xml:"name,attr"`
		Content string   `xml:",innerxml"`
	}{}

	err := xml.Unmarshal([]byte(aimlTemplate.Content), &setStruct)

	if err != nil {
		return err
	}

	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<set name="`+setStruct.Name+`">`, "", -1)
	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `</set>`, "", -1)

	aiml.Memory[setStruct.Name] = setStruct.Content
	return nil
}

func (aimlTemplate *AIMLTemplate) ProcessGet(aiml *AIML) error {
	getStruct := struct {
		XMLName xml.Name `xml:"get"`
		Name    string   `xml:"name,attr"`
	}{}

	err := xml.Unmarshal([]byte(aimlTemplate.Content), &getStruct)

	if err != nil {
		return err
	}

	content, ok := aiml.Memory[getStruct.Name]

	if !ok {
		return errors.New("Key not found in memory")
	}

	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<get name="`+getStruct.Name+`"/>`, content, -1)
	return nil
}

func (aimlTemplate *AIMLTemplate) ProcessBot(aiml *AIML) error {
	botStruct := struct {
		XMLName xml.Name `xml:"bot"`
		Name    string   `xml:"name,attr"`
	}{}

	err := xml.Unmarshal([]byte(aimlTemplate.Content), &botStruct)

	if err != nil {
		return err
	}

	content, ok := aiml.Bot[botStruct.Name]

	if !ok {
		return errors.New("Key not found in memory")
	}

	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, `<bot name="`+botStruct.Name+`"/>`, content, -1)
	return nil
}

func (aimlTemplate *AIMLTemplate) ProcessSrai(aiml *AIML) (*AIMLTemplate, error) {
	sraiStruct := struct {
		XMLName xml.Name `xml:"srai"`
		Content string   `xml:",innerxml"`
	}{}

	err := xml.Unmarshal([]byte(aimlTemplate.Content), &sraiStruct)

	if err != nil {
		return nil, err
	}

	sraiStruct.Content = strings.Replace(sraiStruct.Content, `<srai>`, "", -1)
	sraiStruct.Content = strings.Replace(sraiStruct.Content, `</srai>`, "", -1)

	ret, errPattern := aiml.findPattern(sraiStruct.Content, true)

	if errPattern != nil {
		return nil, errPattern
	}

	return ret, nil
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (aimlTemplate *AIMLTemplate) ProcessRandom(aiml *AIML) error {
	randomStruct := struct {
		XMLName xml.Name `xml:"random"`
		List    []struct {
			XMLName xml.Name `xml:"li"`
			Content string   `xml:",innerxml"`
		} `xml:"li"`
	}{}

	err := xml.Unmarshal([]byte(aimlTemplate.Content), &randomStruct)

	if err != nil {
		return err
	}

	randIdx := random(0, len(randomStruct.List))
	randContent := randomStruct.List[randIdx]
	aimlTemplate.Content = randContent.Content
	arr := []string{}

	_, errT := aiml.processTemplateTags(aimlTemplate, arr, true)

	if errT != nil {
		return errT
	}

	return nil
}

func (aimlTemplate *AIMLTemplate) ProcessStar(starContent []string) {
	for idx, sContent := range starContent {
		if idx > 0 {
			aimlTemplate.Content = strings.Replace(aimlTemplate.Content, "<star/>", strings.TrimSpace(sContent), 1)
		}
	}

	aimlTemplate.Content = strings.Replace(aimlTemplate.Content, "<star/>", "", -1)
}
