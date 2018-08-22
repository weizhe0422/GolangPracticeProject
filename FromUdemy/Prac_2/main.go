package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
	getBotInfo() botInfo
}
type botInfo struct {
	language string
	platform string
}

type englishBot struct {
	enable bool
	botInfo
}

type spanishBot struct {
	enable bool
	botInfo
}

func main() {
	eb := englishBot{
		enable: true,
		botInfo: botInfo{
			language: "English",
			platform: "Facebook",
		},
	}

	sb := spanishBot{
		enable: false,
		botInfo: botInfo{
			language: "Spanish",
			platform: "Line",
		},
	}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
	fmt.Println(b.getBotInfo())
}

func (eb englishBot) getBotInfo() botInfo {
	return eb.botInfo
}

func (sb spanishBot) getBotInfo() botInfo {
	return sb.botInfo
}

func (englishBot) getGreeting() string {
	return "Welcom!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
