package commands

import (
	"github.com/MathisBurger/AnalAsia/internal/database"
	"github.com/MathisBurger/AnalAsia/internal/database/models"
	"github.com/MathisBurger/AnalAsia/pkg/algorithms"
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"strconv"
)

// This command sends the top list of general words
// on the server (max 10 words)
func TopCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	words := algorithms.BubbleSortWordModels(database.GetAllWords())

	var TopWords []models.WordModel

	if len(words) < 11 {
		TopWords = words
	} else {
		TopWords = words[:10]
	}

	for i, j := 0, len(TopWords)-1; i < j; i, j = i+1, j-1 {
		TopWords[i], TopWords[j] = TopWords[j], TopWords[i]
	}

	emb := embed.NewEmbed()
	emb.SetTitle("Top Words")
	emb.SetDescription("These are the top 10 Words used on this server")
	for pos, word := range TopWords {
		emb.AddField(strconv.Itoa(pos)+":", "word: "+word.Word+"\ncounter: "+strconv.Itoa(word.Counter))
	}
	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}
