package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoshigayan/zocms/controllers/admin"
)

func Init(g *gin.RouterGroup) {
	admin.AdminController{}.Init(g.Group("/admin"))
}
