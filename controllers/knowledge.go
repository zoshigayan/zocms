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
	g.GET("/:slug", kc.Show)
}

func (kc KnowledgeController) Index(c *gin.Context) {
	db := db.DbManager()
	knowledges := []models.Knowledge{}
	db.Find(&knowledges)

	c.HTML(200, "knowledges/index", gin.H{
		"knowledges": knowledges,
	})
}

func (kc KnowledgeController) Show(c *gin.Context) {
	db := db.DbManager()
	knowledge := models.Knowledge{}
	db.First(&knowledge, "slug = ?", c.Param("slug"))

	c.HTML(200, "knowledges/show", gin.H{
		"Knowledge": knowledge,
	})
}
