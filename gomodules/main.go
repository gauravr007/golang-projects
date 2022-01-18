package main

import (
	"fmt"

	"rsc.io/quote/v3"
)

func main() {
	fmt.Println("Hello")
}

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
