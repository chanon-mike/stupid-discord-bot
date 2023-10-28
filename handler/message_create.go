package handler

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

// This function will be called every time a new message is created
// on any channel that the authenticated bot has access to.
func OnMessageCreateHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == s.State.User.ID {
		return
	}

	switch {
	// If เป็น or เปง or เปน
	case regexp.MustCompile(`เป[็นง]`).MatchString(msg.Content):
		s.ChannelMessageSend(msg.ChannelID, "เป็นตาธรรม")
	}
}
