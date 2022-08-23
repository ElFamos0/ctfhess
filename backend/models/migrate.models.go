package models

import "backend/db"

func init() {
	// Migrate the schema
	db.DB.AutoMigrate(
		&User{},
		&Challenge{},
		&ChallengePage{},
		&Claims{},
		&Completion{},
		&File{},
	)
}
