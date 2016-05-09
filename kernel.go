package goaiml

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (aiml *AIMLInterpreter) LearnFromXML(xmlBytes []byte) error {
	err := xml.Unmarshal(xmlBytes, &aiml.Root)

	if err != nil {
		return err
	}

	err = aiml.compilePatterns()

	if err != nil {
		return err
	}

	return nil
}

func (aiml *AIMLInterpreter) learnFromMoreXML(learnFile string) error {
	xmlFile, err := os.Open(learnFile)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)

	tmpAiml := NewAIMLInterpreter()
	err = xml.Unmarshal(bytes, &tmpAiml.Root)

	if err != nil {
		return err
	}

	err = tmpAiml.compilePatterns()

	if err != nil {
		return err
	}

	aiml.Root.Categories = append(aiml.Root.Categories, tmpAiml.Root.Categories...)
	aiml.Root.Topics = append(aiml.Root.Topics, tmpAiml.Root.Topics...)
	return nil
}

func (aiml *AIMLInterpreter) LearnFromFile(mainFile string) error {
	xmlFile, err := os.Open(mainFile)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)

	return aiml.LearnFromXML(bytes)
}

func (aiml *AIMLInterpreter) Brain() {
	for _, category := range aiml.Root.Categories {
		log.Println(fmt.Sprintf("Pattern %s - Template %s", category.Pattern.Content, category.Template.Content))
	}
}

func (aiml *AIMLInterpreter) compilePatterns() error {
	for _, category := range aiml.Root.Categories {
		if strings.Contains(category.Pattern.Content, "<bot") {
			var err error = nil
			category.Pattern.Content, err = aiml.ProcessBotTag(category.Pattern.Content)
			if err != nil {
				return err
			}
		}

		err := category.Pattern.Regexify()
		if err != nil {
			return err
		}
	}

	return nil
}

func (aiml *AIMLInterpreter) Respond(input string) (string, error) {
	ret, err := aiml.findPattern(input, false)

	if err != nil {
		return "", err
	}

	if strings.Contains(ret, "<srai") {
		return "", errors.New("Srai reference not found")
	}

	return postFormatInput(ret), nil
}

func (aiml *AIMLInterpreter) processBasicTemplateTags(template string, matchRes []string) (string, error) {
	var err error = nil

	if strings.Contains(template, "<star") {
		template = aiml.ProcessStarTag(template, matchRes)
	}

	if strings.Contains(template, "<set") {
		template, err = aiml.ProcessSetTag(template)
		if err != nil {
			return template, err
		}
	}

	if strings.Contains(template, "<get") {
		template, err = aiml.ProcessGetTag(template)
		if err != nil {
			return template, err
		}
	}

	if strings.Contains(template, "<bot") {
		template, err = aiml.ProcessBotTag(template)
		if err != nil {
			return template, err
		}
	}

	return template, err
}

func (aiml *AIMLInterpreter) processAllTemplateTags(template string, matchRes []string, looped bool) (string, error) {
	var err error = nil

	if strings.Contains(template, "<learn") {
		template, err = aiml.ProcessLearnTag(template)
		if err != nil {
			return template, err
		}
	}

	if strings.Contains(template, "<think") {
		template, err = aiml.ProcessThinkTag(template, matchRes)
		if err != nil {
			return template, err
		}
	}

	if err != nil {
		return template, err
	}

	if strings.Contains(template, "<random") {
		template, err = aiml.ProcessRandomTag(template, matchRes)
		if err != nil {
			return template, err
		}
	}

	if strings.Contains(template, "<condition") {
		template, err = aiml.ProcessConditionTag(template)
		if err != nil {
			return template, err
		}
	}

	if strings.Contains(template, "<srai") && !looped {
		return aiml.ProcessSraiTag(template)
	}

	return aiml.processBasicTemplateTags(template, matchRes)
}

func (aiml *AIMLInterpreter) findPattern(input string, looped bool) (string, error) {
	input = preFormatInput(input)
	for _, category := range aiml.Root.Categories {
		matchRes := category.Pattern.Re.FindStringSubmatch(input)
		if len(matchRes) > 0 {
			return aiml.processAllTemplateTags(category.Template.Content, matchRes, looped)
		}
	}

	return input, errors.New("Template not found")
}
