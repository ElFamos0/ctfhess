package logincontrollers

import (
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func googleCheck(ctx *gin.Context) {
	// try to get the user without re-authenticating
	provider, err := goth.GetProvider("google")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	values := ctx.Request.URL.Query()
	values.Add("provider", "google")
	ctx.Request.URL.RawQuery = values.Encode()

	value, err := gothic.GetFromSession(provider.Name(), ctx.Request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}

	user, err := provider.FetchUser(sess)
	if err != nil {
		params := ctx.Request.URL.Query()
		if params.Encode() == "" && ctx.Request.Method == "POST" {
			ctx.Request.ParseForm()
			params = ctx.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
			return
		}

		gu, err := provider.FetchUser(sess)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
			return
		}

		user = gu
	}

	token, err := provider.RefreshToken(user.RefreshToken)
	if err == nil {
		user.AccessToken = token.AccessToken
		user.RefreshToken = token.RefreshToken
		gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)
	}

	u, err := services.GetUserByID(user.UserID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"logged": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": u, "logged": true})
}
