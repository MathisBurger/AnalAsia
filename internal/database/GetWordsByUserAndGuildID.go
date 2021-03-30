package database

import "github.com/MathisBurger/AnalAsia/internal/database/models"

func GetWordsByUserAndGuildID(userID string, guildID string) []models.WordModel {
	conn := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `user-words` WHERE `userID`=? AND `guildID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(userID, guildID)
	defer resp.Close()
	cache := models.UserWordModel{}.ParseAll(resp)

	var words []models.WordModel

	for _, el := range cache {
		words = append(words, models.WordModel{
			el.ID,
			el.Word,
			el.Counter,
			el.GuildID,
		})
	}

	return words
}
