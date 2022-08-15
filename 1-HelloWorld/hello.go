//Main package in go, is where the entry point ( main func) is located
package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Ola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "English"))
}
