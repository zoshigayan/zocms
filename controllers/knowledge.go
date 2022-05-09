package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/db"
	"github.com/zoshigayan/zocms/models"
)

type KnowledgeController struct {
}

func (kc KnowledgeController) Init(g *gin.RouterGroup) {
	g.GET("/", kc.Index)
}

func (kc KnowledgeController) Index(c *gin.Context) {
	db := db.DbManager()
	knowledges := []models.Knowledge{}
	db.Find(&knowledges)
	c.JSON(200, gin.H{
		"knowledges": knowledges,
	})
}
