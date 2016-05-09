package goaiml

import "encoding/xml"

const BOT_NAME string = "GOAIMLBot"
const GITHUB string = "https://github.com/eduardonunesp/goaiml"

type AIMLInterpreter struct {
	Memory map[string]string
	Bot    map[string]string
	Root   AIMLRoot
}

type AIMLRoot struct {
	XMLName    xml.Name       `xml:"aiml"`
	Version    string         `xml:"version,attr"`
	Encoding   string         `xml:"encoding,attr"`
	Categories []AIMLCategory `xml:"category"`
	Topics     []AIMLTopic    `xml:"topic"`
}

type AIMLTopic struct {
	XMLName    xml.Name       `xml:"topic"`
	Name       string         `xml:"name,attr"`
	Categories []AIMLCategory `xml:"category"`
}

type AIMLCategory struct {
	XMLName  xml.Name     `xml:"category"`
	Pattern  AIMLPattern  `xml:"pattern"`
	Template AIMLTemplate `xml:"template"`
}

type AIMLTemplate struct {
	XMLName xml.Name `xml:"template"`
	Content string   `xml:",innerxml"`
	Looped  bool
}

type AIMLPattern struct {
	XMLName xml.Name `xml:"pattern"`
	Content string   `xml:",innerxml"`
}

func NewAIMLInterpreter() *AIMLInterpreter {
	ret := &AIMLInterpreter{
		Memory: make(map[string]string),
		Bot:    make(map[string]string),
	}

	ret.Bot["name"] = BOT_NAME

	return ret
}
