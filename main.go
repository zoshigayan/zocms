package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	controllers.KnowledgeController{}.Init(router.Group("/knowledge"))
	controllers.RandomController{}.Init(router.Group("/random"))
	router.Run()
}
