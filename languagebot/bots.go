package languagebot

import "fmt"

type EnglishBot struct{}

type SpanishBot struct{}

type Bot interface {
	getGreeting() string
}

func (EnglishBot) getGreeting() string {
	return "Hi there."
}

func (SpanishBot) getGreeting() string {
	return "Hola"
}

func PrintGreeting(bot Bot) {
	fmt.Println(bot.getGreeting())
}

func UseBots() {
	var eBot EnglishBot
	var sBot SpanishBot

	PrintGreeting(eBot)
	PrintGreeting(sBot)
}
