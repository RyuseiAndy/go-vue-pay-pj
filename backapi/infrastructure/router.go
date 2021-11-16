package infrastructure

import (
	"github.com/RyuseiAndy/go-vue-pay-pj/backapi/handler"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.GET("/api/v1/items", func(c *gin.Context) { handler.GetLists(c) })
	router.GET("/api/v1/items/:id", func(c *gin.Context) { handler.GetItem(c) })
	router.POST("/api/v1/charge/items/:id", func(c *gin.Context) { handler.Charge(c) })

	Router = router
}
