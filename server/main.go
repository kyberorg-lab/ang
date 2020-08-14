package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kyberorg/ang/api"
	"github.com/kyberorg/ang/database"
	"log"
	"path"
	"path/filepath"
)

const (
	WebappRoot = "./webapp/dist/webapp/"
)

var (
	router = gin.Default()
)

func init() {
	database.InitDatabase()
}

func setupRoutes() {
	heroesApiGroup := router.Group("/api/heroes")
	heroApi := api.HeroApi{}
	{
		heroesApiGroup.GET("", heroApi.GetAllHeroes)
		heroesApiGroup.POST("", heroApi.AddHero)
		heroesApiGroup.PUT(":id", heroApi.UpdateHero)
		heroesApiGroup.DELETE(":id", heroApi.DeleteHero)
	}

	router.NoRoute(serveStatic)
}

func start() {
	defer database.CloseDatabase()
	log.Fatal(router.Run(":8080"))
}

func serveStatic(c *gin.Context) {
	dir, file := path.Split(c.Request.RequestURI)
	ext := filepath.Ext(file)
	if file == "" || ext == "" {
		c.File(WebappRoot + "/index.html")
	} else {
		c.File(WebappRoot + path.Join(dir, file))
	}
}

func main() {
	setupRoutes()
	start()
}
