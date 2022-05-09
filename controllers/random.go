package controllers

import "github.com/gin-gonic/gin"

type RandomController struct {
}

func (rc RandomController) Init(g *gin.RouterGroup) {
	g.GET("/", rc.Index)
}

func (rc RandomController) Index(c *gin.Context) {
	c.String(200, "ok from random")
}
