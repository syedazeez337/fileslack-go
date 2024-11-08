package main

import "github.com/shomali11/slacker"

func main() {
	botToken := "SLACK_BOT_TOKEN"
	appToken := "SLACK_APP_TOKEN"

	bot := slacker.NewClient(botToken, appToken)

	
}
