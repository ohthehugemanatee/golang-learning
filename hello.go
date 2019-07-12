package main

import "fmt"

var helloPrefix = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
}

func main() {
	fmt.Println(hello("world", ""))
}

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "" || helloPrefix[language] == "" {
		language = "English"
	}
	return helloPrefix[language] + name
}
