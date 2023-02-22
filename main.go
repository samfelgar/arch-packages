package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	updates, err := getUpdates(args...)

	if err != nil {
		log.Fatal(err.Error())
	}

	if len(updates) == 0 {
		Notify(Notification{
			Title: "No updates found",
		})
		return
	}

	var content string

	for _, update := range updates {
		content += fmt.Sprintf("%s\n", update.Title)
	}

	Notify(Notification{
		Title:   "Updates found",
		Content: content,
	})
}

func getUpdates(packages ...string) ([]Item, error) {
	items, err := FetchRss()

	if err != nil {
		return nil, err
	}

	if len(packages) == 0 {
		return items, nil
	}

	var matches []Item

	for _, item := range items {
		for _, pkg := range packages {
			normalizedTitle := strings.ToLower(item.Title)
			normalizedPackage := strings.ToLower(pkg)

			if strings.Contains(normalizedTitle, normalizedPackage) {
				matches = append(matches, item)
			}
		}
	}

	return matches, nil
}
