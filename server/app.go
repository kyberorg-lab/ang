package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

const (
	Webapp_root = "./webapp/dist/webapp/"
)

type App struct {
	router *gin.Engine
}

type Ping struct {
	Ping string
}

func (app *App) start() {
	app.router.GET("/ping", app.getPing)
	app.router.NoRoute(app.serveStatic)
	log.Fatal(app.router.Run(":8080"))
}

func (app *App) serveStatic(c *gin.Context) {
	dir, file := path.Split(c.Request.RequestURI)
	ext := filepath.Ext(file)
	if file == "" || ext == "" {
		c.File(Webapp_root + "/index.html")
	} else {
		c.File(Webapp_root + path.Join(dir, file))
	}
}

func (app *App) getPing(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, Ping{Ping: "pong"})
}
