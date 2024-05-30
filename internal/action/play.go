package action

import (
	"github.com/bwmarrin/discordgo"
	"klammerAeffchen/pkg"
)

func PlayerSelectHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_ = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Playing sound",
		},
	})
	data := i.MessageComponentData()
	value := data.Values[0]
	guild, _ := s.State.Guild(i.GuildID)
	for _, state := range guild.VoiceStates {
		if state.UserID == i.Member.User.ID {
			pkg.PlayFile(s, i.GuildID, state.ChannelID, value)
		}
	}
}
