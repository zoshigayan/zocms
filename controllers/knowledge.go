package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/db"
	"github.com/zoshigayan/zocms/models"
	"time"
)

type KnowledgeController struct {
}

func (kc KnowledgeController) Init(g *gin.RouterGroup) {
	g.GET("/", kc.Index)

	g.GET("/:slug", kc.Show)
	g.POST("/:slug", kc.Update)

	g.GET("/new", kc.New)
	g.POST("/new", kc.Create)

	g.DELETE("/:slug", kc.Destroy)
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

	c.HTML(200, "knowledges/edit", gin.H{
		"Knowledge": knowledge,
	})
}

func (kc KnowledgeController) Update(c *gin.Context) {
	submitted := knowledgeByParams(c)

	db := db.DbManager()
	knowledge := models.Knowledge{}
	db.First(&knowledge, "slug = ?", c.Param("slug"))
	db.Model(&knowledge).Updates(submitted)

	c.Redirect(302, fmt.Sprintf("/knowledge/%s", submitted.Slug))
}

func (kc KnowledgeController) New(c *gin.Context) {
	c.HTML(200, "knowledges/edit", gin.H{})
}

func (kc KnowledgeController) Create(c *gin.Context) {
	submitted := knowledgeByParams(c)

	db := db.DbManager()
	db.Create(&submitted)
	c.Redirect(302, fmt.Sprintf("/knowledge/%s", submitted.Slug))
}

func (kc KnowledgeController) Destroy(c *gin.Context) {
}

func knowledgeByParams(c *gin.Context) models.Knowledge {
	Slug := c.PostForm("Slug")
	Title := c.PostForm("Title")
	BodyMD := c.PostForm("BodyMD")
	BodyHTML := c.PostForm("BodyHTML")
	PublishedAt, _ := time.Parse("2006-01-02", c.PostForm("PublishedAt"))
	Draft := c.PostForm("Draft") == "on"

	return models.Knowledge{
		Slug:        Slug,
		Title:       Title,
		BodyMD:      BodyMD,
		BodyHTML:    BodyHTML,
		PublishedAt: PublishedAt,
		Draft:       Draft,
	}
}
