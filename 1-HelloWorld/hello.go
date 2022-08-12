//Main package in go, is where the entry point ( main func) is located
package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
