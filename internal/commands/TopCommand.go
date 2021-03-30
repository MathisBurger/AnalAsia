package commands

import (
	"fmt"
	"github.com/MathisBurger/AnalAsia/internal/database"
	"github.com/MathisBurger/AnalAsia/internal/database/models"
	"github.com/MathisBurger/AnalAsia/pkg/algorithms"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"strconv"
	"strings"
)

// This command sends the top list of general words
// on the server (max 10 words)
func TopCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	var words []models.WordModel
	var description string

	spl := strings.Split(m.Content, " ")

	if len(spl) == 1 {
		description = "These are the top 10 Words used on this server"
		words = algorithms.BubbleSortWordModels(database.GetAllWordsOfGuild(m.GuildID))
	} else {
		if spl[1] == "self" {
			description = fmt.Sprintf("These are the top 10 Words used by `%s` on this server", m.Author.Username)
			words = algorithms.BubbleSortWordModels(database.GetWordsByUserAndGuildID(m.Author.ID, m.GuildID))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid command action")
			return
		}
	}

	var TopWords []models.WordModel

	if len(words) < 11 {
		TopWords = words
	} else {
		TopWords = words[len(words)-10:]
	}

	for i, j := 0, len(TopWords)-1; i < j; i, j = i+1, j-1 {
		TopWords[i], TopWords[j] = TopWords[j], TopWords[i]
	}

	emb := embed.NewEmbed()
	emb.SetTitle("Top Words")
	emb.SetDescription(description)
	for pos, word := range TopWords {
		emb.AddField(strconv.Itoa(pos+1)+":", "word: "+word.Word+"\ncounter: "+strconv.Itoa(word.Counter))
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}
