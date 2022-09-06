package userscontrollers

import (
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GraphUsers struct {
	RawUser *models.User `json:"-"`

	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	Graph  []int  `json:"graph"`
}

type GraphResp struct {
	Users []*GraphUsers `json:"users"`
	Days  []string      `json:"days"`
}

func UserToGraphUser(user *models.User) *GraphUsers {
	return &GraphUsers{
		RawUser: user,
		ID:      user.ID,
		Name:    fmt.Sprintf("%s %s", user.Name, user.Surname),
		Points:  user.Points,
	}
}

func GetGraph(ctx *gin.Context) {
	promo := ctx.Param("promo")

	u := []*models.User{}
	switch promo {
	case "all":
		db.DB.
			//Preload("Completions.Challenge").
			Preload("Completions").
			Find(&u)
	default:
		db.DB.
			//Preload("Completions.Challenge").
			Preload("Completions").
			Where("promo = ?", premiereAnnee).
			Find(&u)
	}

	var times []time.Time
	db.DB.
		Table("completions").
		Select("created_at").
		Order("created_at asc").
		Scan(&times)

	// get days from times
	var days []string
	for _, t := range times {
		f := t.Format("02 Jan")
		if len(days) == 0 {
			// On ajoute le jour d'avant pour le point 0
			days = append(days, t.Add(-24*time.Hour).Format("02 Jan"))
			days = append(days, f)
		} else if days[len(days)-1] != f {
			days = append(days, f)
		}
	}

	var graph []*GraphUsers
	// Insert users in the scoreboard ordering by points
	for _, user := range u {
		user.Points = user.GetPoints()
		if len(graph) == 0 {
			graph = append(graph, UserToGraphUser(user))
		} else {
			inserted := false
			for i, sb := range graph {
				if user.Points > sb.Points {
					// insert in the scoreboard
					inserted = true
					graph = append(graph[:i], append([]*GraphUsers{UserToGraphUser(user)}, graph[i:]...)...)
					break
				}
			}
			// put at the end
			if !inserted {
				graph = append(graph, UserToGraphUser(user))
			}
		}
	}

	// We keep the top 5
	if len(graph) > 5 {
		graph = graph[:5]
	}

	// We now generate the data for their graphs
	for _, user := range graph {
		user.Graph = make([]int, len(days))
		for i, day := range days {
			if i > 0 {
				user.Graph[i] = user.Graph[i-1]
			}
			for _, completion := range user.RawUser.Completions {
				if completion.CreatedAt.Format("02 Jan") == day {
					user.Graph[i] += completion.Challenge.Points
				}
			}
		}
	}

	resp := &GraphResp{
		Users: graph,
		Days:  days,
	}

	ctx.JSON(http.StatusOK, resp)
}

func RegisterGraphControllersGet(router *gin.RouterGroup) {
	router.GET("/graph/:promo", GetGraph)
}
