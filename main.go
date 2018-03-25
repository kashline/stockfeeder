package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL("http://articlefeeds.nasdaq.com/nasdaq/symbols?symbol=AMTD")
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(feed.Items); i++ {

		fmt.Println(feed.Items[i].Link)
	}
}
