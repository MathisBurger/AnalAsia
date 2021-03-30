package models

import "database/sql"

// The database model for
// the word table
type UserWordModel struct {
	ID      int    `json:"id"`
	Word    string `json:"word"`
	Counter int    `json:"counter"`
	UserID  string `json:"user_id"`
	GuildID string `json:"guild_id"`
}

// This function parses a single Row of
// the type UserWordModel
func (c UserWordModel) Parse(resp *sql.Rows) UserWordModel {
	var mdl UserWordModel
	_ = resp.Scan(&mdl.ID, &mdl.Word, &mdl.Counter, &mdl.UserID, &mdl.GuildID)
	return mdl
}

// This function parses all words in the table `words`
// of the type UserWordModel
func (c UserWordModel) ParseAll(resp *sql.Rows) []UserWordModel {
	var answers []UserWordModel
	for resp.Next() {
		var mdl UserWordModel
		_ = resp.Scan(&mdl.ID, &mdl.Word, &mdl.Counter, &mdl.UserID, &mdl.GuildID)
		answers = append(answers, mdl)
	}
	return answers
}
