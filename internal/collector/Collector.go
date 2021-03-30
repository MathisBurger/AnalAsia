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

		var modifiedParts []string

		for _, raw := range parts {
			if strings.Contains(raw, "\n") {
				spl := strings.Split(raw, "\n")
				for _, el := range spl {
					modifiedParts = append(modifiedParts, el)
				}
			} else {
				modifiedParts = append(modifiedParts, raw)
			}
		}

		for _, raw := range modifiedParts {
			word := algorithms.RemovePunicationMarks(strings.ToLower(raw))
			if CheckExistance(word, m.GuildID) {
				IncreaseWord(word, m.GuildID)
			} else {
				CreateWord(word, m.GuildID)
			}
		}
	}
}
