package main

import (
	"backend/db"
	"backend/server"
)

func main() {

	db.InitDatabase()
	/*
		router := gin.Default()
		router.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World!")
		})
		router.Run("localhost:5050")
	*/
	server.InitServer()

}
