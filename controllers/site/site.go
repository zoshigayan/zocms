package preview

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/models"
	// "time"
)

type PreviewController struct {
}

type PreviewParams struct {
	ID uint `uri:"id" binding:"required"`
}

func (pc PreviewController) Init(g *gin.RouterGroup) {
	g.GET("/", pc.Index)
	g.GET("/knowledge", pc.KnowledgeIndex)
	g.GET("/random", pc.RandomIndex)
	g.GET("/quote", pc.QuoteIndex)
	g.GET("/entries/:id", pc.Show)
}

func (pc PreviewController) Index(c *gin.Context) {
	entries := models.EntryAll()
	c.HTML(200, "site/index", gin.H{
		"Entries": entries,
	})
}

func (pc PreviewController) KnowledgeIndex(c *gin.Context) {
	entries := models.KnowledgeAll()
	c.HTML(200, "site/knowledge/index", gin.H{
		"Entries": entries,
	})
}

func (pc PreviewController) RandomIndex(c *gin.Context) {
}

func (pc PreviewController) QuoteIndex(c *gin.Context) {
}

func (pc PreviewController) Show(c *gin.Context) {
	var previewParams PreviewParams
	if err := c.ShouldBindUri(&previewParams); err != nil {
		c.HTML(404, "site/notFound", gin.H{})
		return
	}
	entry := models.EntryFind(previewParams.ID)
	c.HTML(200, "site/show", gin.H{
		"Entry": entry,
	})
}
