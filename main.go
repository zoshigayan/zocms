package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-zglob"
	"github.com/zoshigayan/zocms/controllers"
	"github.com/zoshigayan/zocms/db"
	"github.com/zoshigayan/zocms/services"
	"log"
)

func main() {
	db.Init()
	services.Migrate()

	router := gin.Default()
	templates, err := zglob.Glob("views/**/*.html")
	if err != nil {
		log.Fatal(err)
	}
	router.LoadHTMLFiles(templates...)
	controllers.Init(router.Group(""))
	router.Static("/assets", "./assets")
	router.Run()
}
