package http

import (
	"fmt"
	"net/http"

	"github.com/demig00d/zakaty-service/internal/usecase"
	"github.com/demig00d/zakaty-service/pkg/logger"
	"github.com/demig00d/zakaty-service/pkg/puzzlebot"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, l logger.Interface, uc usecase.Tournament, pb puzzlebot.Puzzlebot) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	r := &tournamentRouter{uc, pb, l}

	handler.POST("/", func(c *gin.Context) {

		var hook puzzlebot.WebHook
		if err := c.BindJSON(&hook); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		fmt.Println(hook)
		switch hook.Command.Name {
		case "Турнирная таблица":
			rating, _ := r.uc.GetRating(hook.User)
			_ = pb.SendMessage(hook.User, rating)
			c.JSON(http.StatusOK, rating)

		default:

		}

	})
}
