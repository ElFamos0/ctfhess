package logincontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func googleLogout(ctx *gin.Context) {
	gothic.Logout(ctx.Writer, ctx.Request)
	http.Redirect(ctx.Writer, ctx.Request, redirectFront, http.StatusTemporaryRedirect)

	ctx.JSON(http.StatusOK, gin.H{})
}
