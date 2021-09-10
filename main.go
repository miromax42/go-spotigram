package main

import (
	//"context"
	"fmt"

	"github.com/miromax42/go-spotigram/cmd/sp"
	"github.com/miromax42/go-spotigram/cmd/tlg"
	//"github.com/zmb3/spotify/v2"
	//"github.com/miromax42/go-spotigram/cmd/tlg"
)

func main() {

	client, user, err := sp.Init()
	if err != nil {
		panic(err)
	}
	_, _, _ = client, user, err
	tlg.Init(client, user)
	// if err != nil {
	// 	panic(err)
	// }
	// _ = bot

	fmt.Print("\t------------------------------ bot init !!!!!!!!!!!!!!!")

	// client, user, err := sp.Init()
	// _, _, _ = client, user, err

	// pl, err := client.CreatePlaylistForUser(context.Background(), user.ID, "TESTT", "testt", true, false)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("New Playlist cleated %v", pl.ID)
}

func must(e error) {
	fmt.Print(e.Error())
}
