package main

import (
	//"context"
	"flag"
	"fmt"

	"github.com/miromax42/go-spotigram/cmd/sp"
	"github.com/miromax42/go-spotigram/cmd/tlg"
	//"github.com/zmb3/spotify/v2"
	//"github.com/miromax42/go-spotigram/cmd/tlg"
)

func main() {
	redirectUri := flag.String("host", "localhost", "sets address for callback spotify")
	flag.Parse()

	client, user, err := sp.Init(*redirectUri)
	if err != nil {
		panic(err)
	}
	// defer tlg.Init(client, user)
	tlg.Init(client, user)
}

func must(e error) {
	fmt.Print(e.Error())
}
