package collector

import (
	"github.com/MathisBurger/AnalAsia/pkg/algorithms"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
)

func Collector(s *discordgo.Session, m *discordgo.MessageCreate) {

	if !strings.HasPrefix(m.Content, os.Getenv("botPrefix")) && !m.Author.Bot && m.Author.ID != s.State.User.ID {
		parts := strings.Split(m.Content, " ")
		for _, raw := range parts {
			word := algorithms.RemovePunicationMarks(strings.ToLower(raw))
			if CheckExistance(word, m.GuildID) {
				IncreaseWord(word, m.GuildID)
			} else {
				CreateWord(word, m.GuildID)
			}
		}
	}
}
