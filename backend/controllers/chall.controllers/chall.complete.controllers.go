package challcontrollers

import (
	"backend/config"
	"backend/controllers/middlewares.go"
	"backend/db"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type complReq struct {
	ID   int    `json:"id"`
	Flag string `json:"flag"`
}

func sendFirstBlood(u *models.User, chall *models.Challenge) {
	// To avoid race cond
	if u.Promo == config.Getenvint("PROMO_1A", 2025) {
		chall.FirstYearBlood = true
		chall.Save()
	}

	// We fill the challenge's completions
	db.DB.Model(&models.Completion{}).Where("chall_id = ?", chall.ID).Count(&chall.Completions)
	var err error
	var data []byte

	if chall.Completions == 0 {
		values := map[string]string{"content": fmt.Sprintf("First blood ! **%s %s** nous à compléter le chall **%s** pour **%d points** !", user.Name, user.Surname, chall.Title, chall.Points)}
		data, err = json.Marshal(values)
		if err != nil {
			fmt.Println("Cannot marshal values :", err)
			return
		}
	} else {
		if !chall.FirstYearBlood {
			values := map[string]string{"content": fmt.Sprintf("First blood (1A) ! **%s %s** nous à compléter le chall **%s** pour **%d points** !", user.Name, user.Surname, chall.Title, chall.Points)}
			data, err = json.Marshal(values)
			if err != nil {
				fmt.Println("Cannot marshal values :", err)
				return
			}
		}
	}

	if len(data) > 0 {
		resp, err := http.Post(os.Getenv("DISCORD_WEBHOOK_URL"), "application/json", bytes.NewBuffer(data))
		if err != nil {
			fmt.Println("Cannot send to webhook :", err)
			return
		}
		resp.Body.Close()
	}
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

	go sendFirstBlood(user, chall)

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
