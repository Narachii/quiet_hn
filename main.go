package main

import (
	"fmt"
	"github.com/Narachii/quiet_hn/hn"
)

func main() {
	var client hn.Client
	ids, err := client.TopItems()
	if err != err {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		item, err := client.GetItem(ids[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s (by%s)\n", item.Title, item.By)
	}
}
