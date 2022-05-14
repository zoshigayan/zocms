package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/models"
	"time"
)

type KnowledgeParams struct {
	ID uint `uri:"id" binding:"required"`
}

type KnowledgeController struct {
}

func (kc KnowledgeController) Init(g *gin.RouterGroup) {
	g.GET("/", kc.Index)

	g.GET("/:id", kc.Show)
	g.POST("/:id", kc.Update)

	g.GET("/new", kc.New)
	g.POST("/new", kc.Create)

	g.DELETE("/:id", kc.Destroy)
}

func (kc KnowledgeController) Index(c *gin.Context) {
	knowledges := models.KnowledgeAll()

	c.HTML(200, "admin/knowledge/index", gin.H{
		"Knowledges": knowledges,
	})
}

func (kc KnowledgeController) Show(c *gin.Context) {
	var knowledgeParams KnowledgeParams
	if err := c.ShouldBindUri(&knowledgeParams); err != nil {
		c.HTML(404, "admin/notFound", gin.H{})
		return
	}
	knowledge := models.KnowledgeFind(knowledgeParams.ID)

	c.HTML(200, "admin/knowledge/edit", gin.H{
		"Knowledge": knowledge,
	})
}

func (kc KnowledgeController) Update(c *gin.Context) {
	var knowledgeParams KnowledgeParams
	if err := c.ShouldBindUri(&knowledgeParams); err != nil {
		c.HTML(404, "admin/notFound", gin.H{})
		return
	}
	submitted := knowledgeByParams(c)
	models.KnowledgeUpdate(knowledgeParams.ID, submitted)

	c.Redirect(302, fmt.Sprintf("/admin/knowledge/%d", submitted.ID))
}

func (kc KnowledgeController) New(c *gin.Context) {
	c.HTML(200, "/admin/knowledge/edit", gin.H{})
}

func (kc KnowledgeController) Create(c *gin.Context) {
	submitted := knowledgeByParams(c)
	models.KnowledgeCreate(submitted)

	c.Redirect(302, fmt.Sprintf("/admin/knowledge/%d", submitted.ID))
}

func (kc KnowledgeController) Destroy(c *gin.Context) {
}

func knowledgeByParams(c *gin.Context) models.Entry {
	Slug := c.PostForm("Slug")
	Title := c.PostForm("Title")
	BodyMD := c.PostForm("BodyMD")
	BodyHTML := c.PostForm("BodyHTML")
	PublishedAt, _ := time.Parse("2006-01-02", c.PostForm("PublishedAt"))
	Draft := c.PostForm("Draft") == "on"

	return models.Entry{
		Slug:        Slug,
		Title:       Title,
		BodyMD:      BodyMD,
		BodyHTML:    BodyHTML,
		PublishedAt: PublishedAt,
		Draft:       Draft,
	}
}
