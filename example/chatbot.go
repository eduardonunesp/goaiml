package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/eduardonunesp/goaiml"
)

func main() {
	aiml := goaiml.NewAIMLInterpreter()
	err := aiml.LearnFromFile("example.aiml.xml")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You: ")
	for scanner.Scan() {
		line := scanner.Text()
		resp, _ := aiml.Respond(line)
		fmt.Println("Robot: " + resp)
		fmt.Print("You: ")
	}
}
