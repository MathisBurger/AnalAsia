package user_words

import (
	"github.com/MathisBurger/AnalAsia/internal/database"
	"github.com/MathisBurger/AnalAsia/internal/database/models"
)

// Checks if given word already exists in database
func CheckExistance(word string, guildID string, userID string) bool {

	conn := database.Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `user-words` WHERE `word`=? AND `guildID`=? AND `userID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(word, guildID, userID)
	defer resp.Close()
	return resp.Next()
}

// Creates a new row for non existing
// word in words table
func CreateWord(word string, userID string, guildID string) {

	conn := database.Connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("INSERT INTO `user-words` (`ID`, `word`, `counter`, `userID`, `guildID`) VALUES (NULL, ?, '1', ?, ?);")
	defer stmt.Close()
	stmt.Exec(word, userID, guildID)
}

// this function increases the counter
// of existing word
func IncreaseWord(word string, userID string, guildID string) {

	conn := database.Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `user-words` WHERE `word`=? AND `guildID`=? AND `userID`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(word, guildID, userID)
	defer resp.Close()
	resp.Next()
	counter := models.UserWordModel{}.Parse(resp).Counter + 1
	stmt, _ = conn.Prepare("UPDATE `user-words` SET `counter`=? WHERE `word`=? AND `guildID`=? AND `userID`=?")
	defer stmt.Close()
	stmt.Exec(counter, word, guildID, userID)
}
