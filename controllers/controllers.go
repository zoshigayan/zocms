package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/controllers/admin"
	"github.com/zoshigayan/zocms/controllers/site"
)

func Init(g *gin.RouterGroup) {
	admin.AdminController{}.Init(g.Group("/admin"))
	preview.PreviewController{}.Init(g.Group("/"))
}
