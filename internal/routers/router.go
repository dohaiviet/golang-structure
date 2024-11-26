package routers

import (
	"example.com/m/internal/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// Group router
	v1 := r.Group("/api/v1")
	{
		// Params in router
		v1.GET("/ping/:id", controllers.NewUserController().GetUserByID)

		// Query params in router
		v1.GET("/ping", controllers.NewPongController().Pong)
	}

	r.GET("/ping", Pong)

	return r
}

func Pong(c *gin.Context) {
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
