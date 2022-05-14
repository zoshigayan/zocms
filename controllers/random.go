package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/models"
	"time"
)

type RandomController struct {
}

type RandomParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (rc RandomController) Init(g *gin.RouterGroup) {
	g.GET("/", rc.Index)

	g.GET("/:id", rc.Show)
	g.POST("/:id", rc.Update)

	g.GET("/new", rc.New)
	g.POST("/new", rc.Create)

	g.DELETE("/:id", rc.Destroy)
}

func (rc RandomController) Index(c *gin.Context) {
	randoms := models.RandomAll()

	c.HTML(200, "random/index", gin.H{
		"Randoms": randoms,
	})
}

func (rc RandomController) Show(c *gin.Context) {
	var randomParams RandomParams
	if err := c.ShouldBindUri(&randomParams); err != nil {
		c.HTML(404, "notFound", gin.H{})
		return
	}
	random := models.RandomFind(randomParams.ID)

	c.HTML(200, "random/edit", gin.H{
		"Random": random,
	})
}

func (rc RandomController) Update(c *gin.Context) {
	var randomParams RandomParams
	if err := c.ShouldBindUri(&randomParams); err != nil {
		c.HTML(404, "notFound", gin.H{})
		return
	}
	submitted := randomByParams(c)
	updated := models.RandomUpdate(randomParams.ID, submitted)

	c.Redirect(302, fmt.Sprintf("/random/%d", updated.ID))
}

func (rc RandomController) New(c *gin.Context) {
	c.HTML(200, "random/edit", gin.H{})
}

func (rc RandomController) Create(c *gin.Context) {
	submitted := randomByParams(c)
	created := models.RandomCreate(submitted)

	c.Redirect(302, fmt.Sprintf("/random/%d", created.ID))
}

func (rc RandomController) Destroy(c *gin.Context) {
}

func randomByParams(c *gin.Context) models.Entry {
	Title := c.PostForm("Title")
	BodyMD := c.PostForm("BodyMD")
	BodyHTML := c.PostForm("BodyHTML")
	PublishedAt, _ := time.Parse("2006-01-02", c.PostForm("PublishedAt"))
	Draft := c.PostForm("Draft") == "on"

	return models.Entry{
		Title:       Title,
		BodyMD:      BodyMD,
		BodyHTML:    BodyHTML,
		PublishedAt: PublishedAt,
		Draft:       Draft,
	}
}
