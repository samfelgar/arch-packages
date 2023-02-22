package main

import (
	"encoding/xml"
	"errors"
	"net/http"
)

type Rss struct {
	Channel Items `xml:"channel"`
}

type Items struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

func FetchRss() ([]Item, error) {
	response, err := http.Get("https://archlinux.org/feeds/packages/")

	if err != nil {
		return nil, errors.New("unable to retrieve the RSS")
	}

	defer response.Body.Close()

	var rss Rss
	decoder := xml.NewDecoder(response.Body)
	err = decoder.Decode(&rss)

	if err != nil {
		return nil, errors.New("unable to parse the RSS body")
	}

	return rss.Channel.Items, nil
}
