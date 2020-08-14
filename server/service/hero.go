package service

import (
	"errors"
	"github.com/kyberorg/ang/database/dao"
	"github.com/kyberorg/ang/database/model"
)

const (
	ErrHeroAlreadyExists = "already have given hero"
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

func (srv HeroService) SaveHero(hero model.Hero) error {
	var err error
	_, err = heroDao.FindHeroByName(hero.Name)
	if err != nil {
		return errors.New(ErrHeroAlreadyExists)
	}
	return srv.UpdateHero(hero)
}

func (srv HeroService) UpdateHero(hero model.Hero) error {
	err := heroDao.SaveHero(hero)
	return err
}

func (srv HeroService) DeleteHero(id uint) error {

	err := heroDao.DeleteHero(id)
	return err
}
