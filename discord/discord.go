package discord

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/chanon-mike/reolym-discord-bot/config"
	"github.com/chanon-mike/reolym-discord-bot/handler"
)

// Start a discord bot
func Start() {
	cfg := config.Load()

	// Initialize discord
	discord, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		log.Fatal("Error initialize discord bot")
	}

	addEventHandlers(discord)

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		log.Fatalf("Error opening connection %v\n", err)
	}
	defer discord.Close()

	// Run until code is terminated
	log.Println("Bot is running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-c
	log.Println("Gracefully shutdown...")
}

// Add all event handlers
func addEventHandlers(s *discordgo.Session) {
	s.AddHandler(handler.OnMessageCreateHandler)
}
