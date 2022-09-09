package challcontrollers

import (
	"backend/controllers/middlewares.go"
	"backend/db"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type complReq struct {
	ID   int    `json:"id"`
	Flag string `json:"flag"`
}

func completeChall(ctx *gin.Context) {
	var req complReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error binding JSON"})
		return
	}

	chall, err := models.GetChallenge(req.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Error getting challenge"})
		return
	}

	if chall.Flag != req.Flag {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "Flag is incorrect"})
		return
	}

	u := ctx.MustGet("user")
	user := u.(*models.User)

	for _, compl := range user.Completions {
		if compl.ChallID == chall.ID {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"error": "You already did that."})
			return
		}
	}

	c := &models.Completion{
		ChallID: req.ID,
		UserID:  ctx.GetString("user_id"),
	}

	c.Save()

	go func() {
		var done []int
		for _, compl := range user.Completions {
			var found bool
			for _, id := range done {
				if compl.ChallID == id {
					found = true
					break
				}
			}
			if found {
				db.DB.Delete(&compl)
			} else {
				done = append(done, compl.ChallID)
			}
		}
	}()

	user.LastSubmission = time.Now().Unix()
	user.Save()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func RegisterCompleteChall(router *gin.RouterGroup) {
	router.POST("/complete", middlewares.EnsureLoggedIn(), completeChall)
}
