package routers

import (
	"chatbot/api"
	"chatbot/messenger"

	"net/http"

	"github.com/gin-gonic/gin"
)

// return a new gin Router instance
func NewRouter() *gin.Engine {
	router := gin.Default()
	// facebook messenger webhook
	router.Any("/webhook", gin.WrapF(messenger.GetInstance().Handler))
	// facebook app policy required
	router.GET("/privacy", func(c *gin.Context) {
		c.String(http.StatusOK, "privacy")
	})
	router.GET("/service", func(c *gin.Context) {
		c.String(http.StatusOK, "service")
	})
	// router group
	v1 := router.Group("/api/")
	{
		v1.GET("/consumer", api.GetConsumer)
		v1.GET("/message", api.GetMessages)
		v1.POST("/message", api.SendMessage)
	}
	return router
}
