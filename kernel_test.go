package goaiml

import (
	"log"
	"testing"
)

func TestLib_Loader(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}
}

func TestLib_Respond_Star(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("My DOGs Name is Bela")
	if err != nil {
		t.Error(err)
	}

	if result != "That is interesting that you have a dog named Bela" {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_Star_Star(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("MAN WHATSUP TO YOU")
	if err != nil {
		t.Error(err)
	}

	if result != "My friends call me "+BOT_NAME {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_Star_Star_Maybe(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("MAN WHATSUP")
	if err != nil {
		t.Error(err)
	}

	if result != "My friends call me "+BOT_NAME {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_Memory(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("My DOGs Name is Bela")
	if err != nil {
		t.Error(err)
	}

	if result != "That is interesting that you have a dog named Bela" {
		log.Println(result)
		t.Error("Result not match")
	}

	result, err = aiml.Respond("WHaT IS My DOGS nAME")
	if err != nil {
		t.Error(err)
	}

	if result != "Your dog's name is Bela" {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_Bot_At_Template(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("Do you have any idea")
	if err != nil {
		t.Error(err)
	}

	if result != "No, I'm sorry, What you saying ?" {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_Bot_At_Pattern(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond(BOT_NAME + " are you there")
	if err != nil {
		t.Error(err)
	}

	if result != "What's up ?" {
		log.Println(result)
		t.Error("Result not match")
	}
}

func TestLib_Respond_At_Srai(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("BLL ASD 123 ASD")
	if err != nil {
		t.Error(err)
	}

	if result != "No, I'm sorry, What you saying ?" {
		t.Error("Result not match")
	}
}

func TestLib_Respond_At_Random(t *testing.T) {
	aiml := NewAIML()
	err := aiml.Learn("test.aiml.xml")
	if err != nil {
		t.Error(err)
	}

	result, err := aiml.Respond("I AM Joking")
	if err != nil {
		t.Error(err)
	}

	if result != "Why are you Joking" && result != "Do your friends call you" && result != "My name is GOAIMLBot" {
		t.Error("Result not match")
	}
}
