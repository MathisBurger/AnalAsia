package database

import "fmt"

// Initializes all database tables
// and checks if they are already existing
func InitDatabase() {

	conn := Connect()
	defer conn.Close()

	stmt, err := conn.Prepare("CREATE TABLE `words` ( `ID` INT NOT NULL AUTO_INCREMENT , `word` TEXT NOT NULL , `counter` INT NOT NULL , `guildID` TEXT NOT NULL, PRIMARY KEY (`ID`));")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Table `words` already exists")
	}

}
