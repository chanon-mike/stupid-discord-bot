package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreateHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case strings.Contains(msg.Content, "ping"):
		s.ChannelMessageSend(msg.ChannelID, "pong")
	}
}
