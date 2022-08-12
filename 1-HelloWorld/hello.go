//Main package in go, is where the entry point ( main func) is located
package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello())
}
