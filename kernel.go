package goaiml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func (aiml *AIMLInterpreter) LearnFromXML(xmlBytes []byte) error {
	return xml.Unmarshal(xmlBytes, &aiml.Root)
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

	if strings.Contains(template, "<srai") && !looped {
		return aiml.ProcessSraiTag(template)
	}

	return aiml.processBasicTemplateTags(template, matchRes)
}

func (aiml *AIMLInterpreter) findPattern(input string, looped bool) (string, error) {
	input = preFormatInput(input)
	for _, category := range aiml.Root.Categories {
		if strings.Contains(category.Pattern.Content, "<bot") {
			var err error = nil
			category.Pattern.Content, err = aiml.ProcessBotTag(category.Pattern.Content)
			if err != nil {
				return category.Pattern.Content, err
			}
		}

		matchRes := category.Pattern.Regexify().FindStringSubmatch(input)
		if len(matchRes) > 0 {
			return aiml.processAllTemplateTags(category.Template.Content, matchRes, looped)
		}
	}

	return input, errors.New("Template not found")
}
