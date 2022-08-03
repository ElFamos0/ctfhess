package controllers

import (
	users_controllers "backend/controllers/users.controllers"

	"github.com/gin-gonic/gin"
)

// permet d'init nos routes de manière propre

func InitControllers(router *gin.RouterGroup) {

	// init google authentification route
	registerLoginRoutes(router)
	// init user routes

	registerUsersLoggedInRoutes(router)

}

func registerUsersLoggedInRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users", setUserStatus(), ensureLoggedIn())
	users_controllers.RegisterUserRoutes(routerGroup)
}
