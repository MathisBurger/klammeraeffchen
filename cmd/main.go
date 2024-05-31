package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/command"
	"klammerAeffchen/internal/configuration"
	"log"
	"os"
	"os/signal"
)

// Initializes the main discord application
func main() {
	var config configuration.Config
	_ = envconfig.Process("klammeraeffchen", &config)
	discord, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildVoiceStates | discordgo.IntentsGuildMembers
	err = discord.Open()
	if err != nil {
		panic(err)
	}
	_, _ = discord.ApplicationCommandCreate(discord.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "sound",
		Description: "Sound",
	})
	discord.AddHandler(playSoundHandler)
	_ = discord.UpdateListeningStatus("Sounds")

	defer discord.Close()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")

	<-stop
}

// Handler for the play sound interaction
func playSoundHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand && i.ApplicationCommandData().Name == "sound" {
		command.SoundboardCommand(s, i)
	} else if i.Type == discordgo.InteractionMessageComponent && i.MessageComponentData().CustomID == "soundboard" {
		action.PlayerSelectHandler(s, i)
	}
}
