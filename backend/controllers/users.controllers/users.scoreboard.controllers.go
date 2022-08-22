package userscontrollers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ScoreboardUsers struct {
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func UserToScoreboardUser(user *models.User) *ScoreboardUsers {
	return &ScoreboardUsers{
		Name:   fmt.Sprintf("%s %s", user.Name, user.Surname),
		Points: user.Points,
	}
}

func GetScoreboard(ctx *gin.Context) {
	u := []*models.User{}
	db.DB.
		//Preload("Completions.Challenge").
		Preload("Completions").
		Find(&u)

	var scoreboard []*ScoreboardUsers
	// Insert users in the scoreboard ordering by points
	for _, user := range u {
		user.Points = user.GetPoints()
		if len(scoreboard) == 0 {
			scoreboard = append(scoreboard, UserToScoreboardUser(user))
		} else {
			inserted := false
			for i, sb := range scoreboard {
				if user.Points > sb.Points {
					// insert in the scoreboard
					inserted = true
					scoreboard = append(scoreboard[:i], append([]*ScoreboardUsers{UserToScoreboardUser(user)}, scoreboard[i:]...)...)
					break
				}
			}
			// put at the end
			if !inserted {
				scoreboard = append(scoreboard, UserToScoreboardUser(user))
			}
		}
	}

	ctx.JSON(http.StatusOK, scoreboard)
}

func RegisterScoreboardControllersGet(router *gin.RouterGroup) {
	router.GET("/scoreboard", GetScoreboard)
}
