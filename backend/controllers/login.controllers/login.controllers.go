package logincontrollers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var scopes = []string{
	"https://www.googleapis.com/auth/userinfo.profile",
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/admin.directory.user.readonly",
}

type Education struct {
	Promo  int    `json:"Promotion"`
	Spe    string `json:"Approfondissement"`
	Statut int    `json:"Statut"`
}

var (
	googleProvider *google.Provider
	clientID       = os.Getenv("G_CLIENT_ID")
	clientSecret   = os.Getenv("G_CLIENT_SECRET")
	callbackURL    = os.Getenv("G_REDIRECT_URI")
	redirectFront  = os.Getenv("REDIRECT_FRONT")
	cookieHost     = os.Getenv("COOKIE_HOST")
)

func init() {
	token, err := uuid.NewV4()
	if err != nil {
		panic("impossible de générer une clé pour les cookies")
	}
	key := token.Bytes()
	maxAge := 86400 * 30 // 30 days cookie

	store := sessions.NewCookieStore(key)
	store.MaxAge(maxAge)
	store.Options.Domain = cookieHost
	store.Options.Path = "/"

	googleProvider = google.New(clientID, clientSecret, fmt.Sprintf("%sapi/login/callback", callbackURL), scopes...)
	googleProvider.SetHostedDomain("telecomnancy.net")
	googleProvider.SetPrompt("consent")
	googleProvider.SetAccessType("offline")
	goth.UseProviders(
		googleProvider,
	)
}

func googleLogin(ctx *gin.Context) {
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
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
		return
	}

	_, err = provider.FetchUser(sess)
	if err != nil {
		params := ctx.Request.URL.Query()
		if params.Encode() == "" && ctx.Request.Method == "POST" {
			ctx.Request.ParseForm()
			params = ctx.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)

		if err != nil {
			gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
			return
		}

		_, err := provider.FetchUser(sess)
		if err != nil {
			gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
			return
		}
	}

	ctx.Redirect(http.StatusFound, redirectFront)
}

func GetUser(ctx *gin.Context) (user goth.User, err error) {
	// try to get the user without re-authenticating
	provider, err := goth.GetProvider("google")
	if err != nil {
		return
	}
	values := ctx.Request.URL.Query()
	values.Add("provider", "google")
	ctx.Request.URL.RawQuery = values.Encode()

	value, err := gothic.GetFromSession(provider.Name(), ctx.Request)
	if err != nil {
		return
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return
	}

	user, err = provider.FetchUser(sess)
	if err != nil {
		params := ctx.Request.URL.Query()
		if params.Encode() == "" && ctx.Request.Method == "POST" {
			ctx.Request.ParseForm()
			params = ctx.Request.Form
		}

		// get new token and retry fetch
		_, err = sess.Authorize(provider, params)
		if err != nil {
			return
		}

		err = gothic.StoreInSession(provider.Name(), sess.Marshal(), ctx.Request, ctx.Writer)
		if err != nil {
			return
		}

		user, err = provider.FetchUser(sess)
		if err != nil {
			return
		}
	}
	return
}

func RegisterLoginRoutes(rg *gin.RouterGroup) {
	loginGroup := rg.Group("/login")
	loginGroup.GET("/", googleLogin)
	loginGroup.GET("/loggedin", googleCheck)
	loginGroup.GET("/callback", googleCallback)

	logoutGroup := rg.Group("/logout")
	logoutGroup.GET("/", googleLogout)
}
