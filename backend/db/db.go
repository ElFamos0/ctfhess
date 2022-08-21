package db

import (
	"backend/config"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.InitEnv()

	// Si on utilise ':=' on 'shadow' la variable DB qui sera définis localement dans la fonction (et pas globalement pour le reste du programme)
	var err error
	if os.Getenv("APP_ENV") == "prod" {
		// mysql database
		DBUSER := os.Getenv("DB_USER")
		DBPASS := os.Getenv("DB_PASS")
		DBNAME := os.Getenv("DB_NAME")
		DBHOST := os.Getenv("DB_HOST")
		DBPORT := os.Getenv("DB_PORT")
		DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASS, DBHOST, DBPORT, DBNAME)), &gorm.Config{})
		fmt.Println("Database connected")
	} else {
		DB, err = gorm.Open(sqlite.Open("./db/prod.db"), &gorm.Config{})
	}

	if err != nil {
		panic(err)
	}
}
