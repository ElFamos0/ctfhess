package controllers

import (
	chall_controllers "backend/controllers/chall.controllers"
	users_controllers "backend/controllers/users.controllers"

	"github.com/gin-gonic/gin"
)

// permet d'init nos routes de manière propre

func InitControllers(router *gin.RouterGroup) {

	// init google authentification route
	registerLoginRoutes(router)
	// init user routes

	registerUsersLoggedInRoutes(router)

	registerChallRoutes(router)
}

func registerChallRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/chall")
	chall_controllers.RegisterChallRoutes(routerGroup)
}

func registerUsersLoggedInRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users", setUserStatus(), ensureLoggedIn())
	users_controllers.RegisterUserRoutes(routerGroup)
}
