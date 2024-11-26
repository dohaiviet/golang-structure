package controllers

import (
	"example.com/m/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	uid := c.Query("uid")

	// Default query, if param = null return default;
	defaultId := c.DefaultQuery("id", "1")

	c.JSON(http.StatusOK, gin.H{ //Map string
		"message":    "pong" + id,
		"uid":        uid,
		"default id": defaultId,
		// Array
		"users": []string{"david", "messi"},
		// User
		"info": uc.userService.GetInfoUserService(),
	})
}
