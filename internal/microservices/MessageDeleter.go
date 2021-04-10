package microservices

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func DeleteMessageAfterTime(s *discordgo.Session, msg *discordgo.Message, secs int) {
	c := time.Tick(time.Duration(secs) * time.Second)

	for _ = range c {
		_ = s.ChannelMessageDelete(msg.ChannelID, msg.ID)
	}
}
