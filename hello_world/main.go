package main

import (
	"fmt"
)

const (
	englishHelloPrefix = "Hello, "
	franchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	gawaiHelloPrefix   = "Aloha, "

	spanish = "Spanish"
	franch  = "Franch"
	gawai   = "Gawai"
)

func main() {
	fmt.Println(Hello("", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) string {
	switch language {
	case franch:
		return franchHelloPrefix
	case spanish:
		return spanishHelloPrefix
	case gawai:
		return gawaiHelloPrefix
	default:
		return englishHelloPrefix
	}
}
