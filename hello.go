package main

import "fmt"

const helloPrefixConstant = "Hello, "
const holaPrefixConstant = "Hola, "
const Spanish = "Spanish"
const French = "French"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == Spanish {
		return holaPrefixConstant + name
	}
	if language == French {
		return "Bonjour, " + name
	}
	return helloPrefixConstant + name
}

func main() {
	fmt.Println(Hello("", ""))
}
