package main

import "os/exec"

type Notification struct {
	Title string
	Content string
}

func Notify(notification Notification) {
	command := exec.Command("notify-send", notification.Title)

	if notification.Content != "" {
		command.Args = append(command.Args, notification.Content)
	}

	command.Run()
}
