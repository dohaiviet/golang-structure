package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	id := c.Param("id")
	uid := c.Query("uid")

	// Default query, if param = null return default;
	defaultId := c.DefaultQuery("id", "1")

	c.JSON(http.StatusOK, gin.H{ //Map string
		"message":    "pong" + id,
		"uid":        uid,
		"default id": defaultId,
		//Array
		"users": []string{"david", "messi"},
	})
}
