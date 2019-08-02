package main

import "fmt"

var helloPrefix = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
}

func _main() {
	fmt.Println(Hello("", ""))
}

// Hello returns "Hello world" in an appropriate language.
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" || helloPrefix[language] == "" {
		language = "English"
	}
	return helloPrefix[language] + name
}
