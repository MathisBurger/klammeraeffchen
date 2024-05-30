package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/command"
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
	_, err = discord.ApplicationCommandCreate(discord.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "sound",
		Description: "Sound",
	})
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand && i.ApplicationCommandData().Name == "sound" {
			command.SoundboardCommand(s, i)
		} else if i.Type == discordgo.InteractionMessageComponent && i.MessageComponentData().CustomID == "soundboard" {
			action.PlayerSelectHandler(s, i)
		}
	})
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
