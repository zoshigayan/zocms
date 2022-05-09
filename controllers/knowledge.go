package controllers

import "github.com/gin-gonic/gin"

type KnowledgeController struct {
}

func (kc KnowledgeController) Init(g *gin.RouterGroup) {
	g.GET("/", kc.Index)
}

func (kc KnowledgeController) Index(c *gin.Context) {
	c.String(200, "ok from knowledge")
}
