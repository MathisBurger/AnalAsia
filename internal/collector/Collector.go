package collector

import (
	"github.com/MathisBurger/AnalAsia/internal/collector/user-words"
	"github.com/MathisBurger/AnalAsia/internal/collector/words"
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
			handleUserWords(word, m)
			handleGuildWords(word, m)
		}
	}
}

func handleUserWords(word string, m *discordgo.MessageCreate) {
	if user_words.CheckExistance(word, m.GuildID, m.Author.ID) {
		user_words.IncreaseWord(word, m.Author.ID, m.GuildID)
	} else {
		user_words.CreateWord(word, m.Author.ID, m.GuildID)
	}
}

func handleGuildWords(word string, m *discordgo.MessageCreate) {
	if words.CheckExistance(word, m.GuildID) {
		words.IncreaseWord(word, m.GuildID)
	} else {
		words.CreateWord(word, m.GuildID)
	}
}
