package goaiml

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func (aiml *AIML) Learn(mainFile string) error {
	xmlFile, err := os.Open(mainFile)
	if err != nil {
		return err
	}
	defer xmlFile.Close()

	bytes, _ := ioutil.ReadAll(xmlFile)

	return xml.Unmarshal(bytes, &aiml.Root)
}

func (aiml *AIML) Respond(input string) (string, error) {
	aimlTemplate, err := aiml.findPattern(input, false)

	if err != nil {
		return "", err
	}

	if strings.Contains(aimlTemplate.Content, "<srai") {
		return "", errors.New("Srai reference not found")
	}

	return strings.TrimSpace(aimlTemplate.Content), nil
}

func (aiml *AIML) processTemplateTags(template *AIMLTemplate, matchRes []string, looped bool) (*AIMLTemplate, error) {
	if strings.Contains(template.Content, "<star") {
		template.ProcessStar(matchRes)
	}

	if strings.Contains(template.Content, "<set") {
		template.ProcessSet(aiml)
	}

	if strings.Contains(template.Content, "<get") {
		template.ProcessGet(aiml)
	}

	if strings.Contains(template.Content, "<bot") {
		template.ProcessBot(aiml)
	}

	if strings.Contains(template.Content, "<srai") && !looped {
		return template.ProcessSrai(aiml)
	}

	if strings.Contains(template.Content, "<random") {
		template.ProcessRandom(aiml)
	}

	return template, nil
}

func (aiml *AIML) findPattern(input string, looped bool) (*AIMLTemplate, error) {
	for _, category := range aiml.Root.Categories {
		input = " " + input + " "
		if strings.Contains(category.Pattern.Content, "<bot") {
			category.Pattern.ProcessBot(aiml)
		}

		matchRes := category.Pattern.Regexify().FindStringSubmatch(input)
		if len(matchRes) > 0 {
			return aiml.processTemplateTags(&category.Template, matchRes, looped)
		}
	}

	return nil, errors.New("Template not found")
}
