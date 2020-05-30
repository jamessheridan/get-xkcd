package main

import (
	"flag"
	"fmt"
	"get-xkcd/client"
	"log"
	"time"
)

func main() {
	comicNo := flag.Int(
		"n", int(client.LatestComic), "Comic number to fetch",
	)
	clientTimeout := flag.Int64(
		"t", int64(client.DefaultClientTimeout.Seconds()), "Client timeout in seconds",
	)
	saveImage := flag.Bool(
		"s", false, "Save image to current directory",
	)
	outputType := flag.String(
		"o", "text", "Print output in format: text/json",
	)
	flag.Parse()

	xkcdclient := client.NewXKCDClient()
	xkcdclient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	comic, err := xkcdclient.Fetch(client.ComicNumber(*comicNo), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(comic.JSON())
	} else {
		fmt.Println(comic.PrettyString())
	}
}
