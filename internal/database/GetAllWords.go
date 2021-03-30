package database

import "github.com/MathisBurger/AnalAsia/internal/database/models"

// This function returns all words
// in table `words` as type []models.WordModel
func GetAllWords() []models.WordModel {
	conn := Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `words`;")
	defer stmt.Close()

	resp, _ := stmt.Query()
	defer resp.Close()
	return models.WordModel{}.ParseAll(resp)
}
