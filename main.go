package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chanon-mike/stupid-discord-bot/discord"
)

func main() {
	lambda.Start(discord.Start)
}
