package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() {
	// init environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func Getenv(key string, defaultvalue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultvalue
	}
	return value
}

func Getenvint(name string, def int) int {
	if v := os.Getenv(name); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
