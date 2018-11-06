package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func main() {
	say, err := cowsay.Say(
		cowsay.Phrase("Go Secure"),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
		cowsay.Rainbow(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}
