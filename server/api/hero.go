package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kyberorg/ang/service"
	"github.com/kyberorg/ang/utils"
	"net/http"
)

var (
	heroService = service.NewHeroService()
)

type HeroJson struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetAllHeroes(context *gin.Context) {
	utils.SetJsonHeader(context)

	allHeroes, err := heroService.FindAllHeroes()
	if err != nil {
		utils.SendError(context, http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusOK, allHeroes)
	}
}
