package userscontrollers

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(router *gin.RouterGroup) {

	Register_user_controllers_get(router)
}
