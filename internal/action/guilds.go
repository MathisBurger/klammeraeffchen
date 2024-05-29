package action

import (
	"github.com/bwmarrin/discordgo"
)

type GuildResponse struct {
	Guilds []GuildResponseModel `json:"guilds"`
}

type GuildResponseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetAllCommonGuilds(dc *discordgo.Session, userId string) GuildResponse {
	guilds := make([]GuildResponseModel, 0)
	for _, guild := range dc.State.Guilds {
		_, err := dc.GuildMember(guild.ID, userId)
		if err == nil {
			guilds = append(guilds, GuildResponseModel{
				ID:   guild.ID,
				Name: guild.Name,
			})
		}
	}
	return GuildResponse{
		Guilds: guilds,
	}
}
