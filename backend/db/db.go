package db

import (
	"backend/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.InitEnv()

	// Si on utilise ':=' on 'shadow' la variable DB qui sera définis localement dans la fonction (et pas globalement pour le reste du programme)
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
