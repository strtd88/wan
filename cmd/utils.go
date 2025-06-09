package cmd

import (
	"database/sql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		return nil, err
	}

	schema, _ := os.ReadFile("sqitch/deploy/001-init.sql")
	_, err = db.Exec(string(schema))
	if err != nil {
		return nil, err
	}

	return db, nil
}
