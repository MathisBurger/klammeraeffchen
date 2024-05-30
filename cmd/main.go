package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/pkg"
	"log"
	"os"
	"os/signal"
)

func main() {
	var config configuration.Config
	err := envconfig.Process("klammeraeffchen", &config)
	if err != nil {
		log.Fatal(err)
	}
	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		panic(err)
	}
	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMembers
	err = discord.Open()
	if err != nil {
		panic(err)
	}
	authChannel := make(chan *pkg.ShortAuthMessage, 1)
	go pkg.ShortShortAuthHandler(authChannel)
	go internal.InitializeWebServer(config, discord, authChannel)
	defer discord.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
