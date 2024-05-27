package action

import (
	"github.com/bwmarrin/discordgo"
)

func ConnectToChannelWithUserId(dc *discordgo.Session, userId string) string {
	for _, guild := range dc.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userId {
				_, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
				if err != nil {
					return err.Error()
				}
				return "Successfully joined channel"
			}
		}
	}
	return "User not in channel"
}
