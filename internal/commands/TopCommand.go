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

	spl := strings.Split(m.Content, " ")

	// check if there are more parameters
	// than the single command
	if len(spl) == 1 {
		description := "These are the top 10 Words used on this server"
		words := algorithms.BubbleSortWordModels(database.GetAllWordsOfGuild(m.GuildID))
		TopWords := getTopWords(words)
		buildTopEmbed(description, TopWords, s, m)
	} else {

		// Generate and send the message embed for every user
		// mentioned in the message
		if len(m.Mentions) > 0 {
			for _, usr := range m.Mentions {
				description := fmt.Sprintf("These are the top 10 Words used by `%s` on this server", usr.Username)
				words := algorithms.BubbleSortWordModels(database.GetWordsByUserAndGuildID(usr.ID, m.GuildID))
				TopWords := getTopWords(words)
				buildTopEmbed(description, TopWords, s, m)
			}
		} else {

			// send error message
			_, _ = s.ChannelMessageSend(m.ChannelID, "Invalid command action")
			return
		}
	}
}

// This function generates a TopWords message embed
// and sends it into the command channel
func buildTopEmbed(description string, TopWords []models.WordModel, s *discordgo.Session, m *discordgo.MessageCreate) {
	emb := embed.NewEmbed()
	emb.SetTitle("Top Words")
	emb.SetDescription(description)
	for pos, word := range TopWords {
		emb.AddField(strconv.Itoa(pos+1)+":", "word: "+word.Word+"\ncounter: "+strconv.Itoa(word.Counter))
	}

	_, _ = s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
}

// This function gets the top words
// from the sorted word list
func getTopWords(words []models.WordModel) []models.WordModel {
	var TopWords []models.WordModel

	if len(words) < 11 {
		TopWords = words
	} else {
		TopWords = words[len(words)-10:]
	}

	for i, j := 0, len(TopWords)-1; i < j; i, j = i+1, j-1 {
		TopWords[i], TopWords[j] = TopWords[j], TopWords[i]
	}
	return TopWords
}
