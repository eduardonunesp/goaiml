package goaiml

import (
	"strings"
	"testing"
)

func Test_Pattern_Regexify_1(t *testing.T) {
	aimlPattern := &AIMLPattern{
		Content: "TEST * IT",
	}

	matches := aimlPattern.Regexify().FindStringSubmatch(preFormatInput("Test NOW THAT'S it"))

	if len(matches) != 2 {
		t.Error("There's no matches:", matches)
	}

	if matches[1] != "NOW THAT'S" {
		t.Error("Result different of value:", matches[1])
	}
}

func Test_Pattern_Regexify_2(t *testing.T) {
	aimlPattern := &AIMLPattern{
		Content: "* IT *",
	}

	matches := aimlPattern.Regexify().FindStringSubmatch(preFormatInput("BEFORE IT AFTER"))

	if len(matches) != 3 {
		t.Error("There's no matches:", matches)
	}

	if matches[1] != "BEFORE" && matches[2] != "AFTER" {
		t.Error("Result different of value")
	}
}

func Test_Pattern_Regexify_3(t *testing.T) {
	aimlPattern := &AIMLPattern{
		Content: "* IT *",
	}

	matches := aimlPattern.Regexify().FindStringSubmatch(preFormatInput("BEFORE IT"))

	if len(matches) != 3 {
		t.Error("There's no matches:", matches)
	}

	if strings.TrimSpace(matches[1]) != "BEFORE" {
		t.Error("Result different of value")
	}
}

func Test_Pattern_Regexify_4(t *testing.T) {
	aimlPattern := &AIMLPattern{
		Content: "* IT *",
	}

	matches := aimlPattern.Regexify().FindStringSubmatch(preFormatInput("IT AFTER"))

	if len(matches) != 3 {
		t.Error("There's no matches:", matches)
	}

	if strings.TrimSpace(matches[2]) != "AFTER" {
		t.Error("Result different of value")
	}
}
