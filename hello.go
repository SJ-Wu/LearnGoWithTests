package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

var languageLookup = map[string]string{
	spanish: spanishHelloPrefix,
	french:  frenchHelloPrefix,
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(name, language)
}

func greetingPrefix(name string, language string) string {
	if language == "" {
		return englishHelloPrefix + name
	}
	return languageLookup[language] + name
}
