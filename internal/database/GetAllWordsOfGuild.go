package database

import "github.com/MathisBurger/AnalAsia/internal/database/models"

// This function returns all words of a specific guild
// in table `words` as type []models.WordModel
func GetAllWordsOfGuild(guildID string) []models.WordModel {
	conn := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `words` WHERE `guildID`=?;")
	defer stmt.Close()

	resp, _ := stmt.Query(guildID)
	defer resp.Close()
	return models.WordModel{}.ParseAll(resp)
}
