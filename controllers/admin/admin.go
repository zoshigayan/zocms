package admin

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (ac AdminController) Init(g *gin.RouterGroup) {
	KnowledgeController{}.Init(g.Group("/knowledge"))
	RandomController{}.Init(g.Group("/random"))
}
