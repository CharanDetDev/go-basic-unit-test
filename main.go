package main

import "fmt"

func main() {
	hello := "Hello"
	fmt.Println(hello, HelloWorld(hello))
}

func HelloWorld(hello string) string {
	if hello != "" {
		return "World"
	}

	return ""
}
