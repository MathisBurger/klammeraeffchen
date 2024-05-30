package pkg

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"log"
)

func PlayFile(dc *discordgo.Session, guildID string, channelID string, file string) {
	vc, err := dc.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		log.Println(err)
	}
	dgvoice.PlayAudioFile(vc, "./uploads/"+file, make(<-chan bool))
	_ = vc.Disconnect()
	if err != nil {
		log.Println(err)
	}
	defer vc.Close()
}
