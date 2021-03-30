package collector

import (
	"github.com/MathisBurger/AnalAsia/internal/database"
	"github.com/MathisBurger/AnalAsia/internal/database/models"
)

// Checks if given word already exists in database
func CheckExistance(word string) bool {
	conn := database.Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `words` WHERE `word`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(word)
	defer resp.Close()
	return resp.Next()
}

// Creates a new row for non existing
// word in words table
func CreateWord(word string) {
	conn := database.Connect()
	defer conn.Close()
	stmt, _ := conn.Prepare("INSERT INTO `words` (`ID`, `word`, `counter`) VALUES (NULL, ?, '1');")
	defer stmt.Close()
	stmt.Exec(word)
}

// this function increases the counter
// of existing word
func IncreaseWord(word string) {
	conn := database.Connect()
	defer conn.Close()

	stmt, _ := conn.Prepare("SELECT * FROM `words` WHERE `word`=?")
	defer stmt.Close()

	resp, _ := stmt.Query(word)
	defer resp.Close()
	resp.Next()
	counter := models.WordModel{}.Parse(resp).Counter + 1
	stmt, _ = conn.Prepare("UPDATE `words` SET `counter`=? WHERE `word`=?")
	defer stmt.Close()
	stmt.Exec(counter, word)
}
