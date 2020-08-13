package service

import (
	"github.com/kyberorg/ang/database/dao"
	"github.com/kyberorg/ang/database/model"
)

var heroDao = dao.NewHeroDao()

type HeroService struct {
}

func NewHeroService() HeroService {
	return HeroService{}
}

func (srv HeroService) FindAllHeroes() ([]model.Hero, error) {
	return heroDao.FindAll()
}
