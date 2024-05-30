package command

import (
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

// SoundboardCommand is the initial command to render the soundboard options
// to the channel and provide the opportunity to select
func SoundboardCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var options []discordgo.SelectMenuOption
	entries, _ := os.ReadDir("./uploads")
	for _, entry := range entries {
		info, _ := entry.Info()
		if strings.Contains(info.Name(), ".mp3") {
			options = append(options, discordgo.SelectMenuOption{
				Label:   info.Name(),
				Value:   info.Name(),
				Default: false,
			})
		}
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags: discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							// Select menu, as other components, must have a customID, so we set it to this value.
							CustomID:    "soundboard",
							Placeholder: "Choose your sound",
							Options:     options,
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
