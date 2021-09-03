package main

import (
	"fmt"

	"github.com/miromax42/go-spotigram/cmd/tlg"
)

func main() {
	bot, err := tlg.Init()
	if err != nil {
		panic(err)
	}
	_ = bot
}

func must(e error) {
	fmt.Print(e.Error())
}
