package internal

import (
	"github.com/bwmarrin/discordgo"
	"klammerAeffchen/internal/configuration"
	"log"
	"os"
	"os/signal"
)

func InitializeDiscordBot(config *configuration.Config) {

	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		panic(err)
	}
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates
	err = discord.Open()
	if err != nil {
		panic(err)
	}

	defer discord.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
