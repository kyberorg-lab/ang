package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kyberorg/ang/database/model"
	"github.com/kyberorg/ang/service"
	"github.com/kyberorg/ang/utils"
	"net/http"
	"strconv"
)

var (
	heroService = service.NewHeroService()
)

type HeroJson struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type HeroApi struct{}

func (api *HeroApi) GetAllHeroes(context *gin.Context) {
	httpUtils := utils.HttpUtils{Context: context}
	httpUtils.SetJsonHeader()

	allHeroes, err := heroService.FindAllHeroes()
	if err != nil {
		httpUtils.SendError(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusOK, allHeroes)
	}
}

func (api *HeroApi) AddHero(context *gin.Context) {
	httpUtils := utils.HttpUtils{Context: context}
	httpUtils.SetJsonHeader()

	var heroInput HeroJson
	err := context.ShouldBindJSON(&heroInput)
	if err != nil {
		httpUtils.SendError(http.StatusBadRequest, err.Error())
		return
	}
	if len(heroInput.Name) == 0 {
		httpUtils.SendError(http.StatusUnprocessableEntity, errors.New("name should be present").Error())
		return
	}
	hero := model.Hero{
		Name: heroInput.Name,
	}
	err = heroService.SaveHero(hero)
	if err != nil {
		httpUtils.SendError(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusCreated, "")
	}
}

func (api *HeroApi) UpdateHero(context *gin.Context) {
	httpUtils := utils.HttpUtils{Context: context}
	httpUtils.SetJsonHeader()

	var heroInput HeroJson
	err := context.ShouldBindJSON(&heroInput)
	if err != nil {
		httpUtils.SendError(http.StatusBadRequest, err.Error())
		return
	}
	id := context.Param("id")
	idAsNum, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		httpUtils.SendError(http.StatusUnprocessableEntity, err.Error())
		return
	}
	if len(heroInput.Name) == 0 {
		httpUtils.SendError(http.StatusUnprocessableEntity, errors.New("name should be present").Error())
		return
	}

	hero := model.Hero{
		Name: heroInput.Name,
	}
	hero.ID = uint(idAsNum)

	err = heroService.UpdateHero(hero)
	if err != nil {
		httpUtils.SendError(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusCreated, "")
	}
}

func (api *HeroApi) DeleteHero(context *gin.Context) {
	httpUtils := utils.HttpUtils{Context: context}
	httpUtils.SetJsonHeader()
	idStr := context.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		httpUtils.SendError(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = heroService.DeleteHero(uint(id))
	if err != nil {
		httpUtils.SendError(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusCreated, "")
	}
}
