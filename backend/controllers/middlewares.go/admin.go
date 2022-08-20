package middlewares

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loggedIn := ctx.GetBool("is_logged_in")
		if !loggedIn {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			return
		}

		u, err := ctx.Get("user")
		if !err {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			return
		}
		if u.(*models.User).Type != models.AdminUser {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token."})
			return
		}
	}
}
