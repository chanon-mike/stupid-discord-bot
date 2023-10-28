package handler

import "github.com/bwmarrin/discordgo"

// Add all event handlers
func AddEventHandlers(s *discordgo.Session) {
	s.AddHandler(OnMessageCreateHandler)
}
