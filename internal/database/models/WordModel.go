package models

import "database/sql"

// The database model for
// the word table
type WordModel struct {
	ID      int    `json:"id"`
	Word    string `json:"word"`
	Counter int    `json:"counter"`
}

// This function parses a single Row of
// the type WordModel
func (c WordModel) Parse(resp *sql.Rows) WordModel {
	var mdl WordModel
	_ = resp.Scan(&mdl.ID, &mdl.Word, &mdl.Counter)
	return mdl
}
