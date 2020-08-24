package dao

import (
	"github.com/kyberorg/ang/database"
	"github.com/kyberorg/ang/database/model"
)

type HeroDao struct {
}

func NewHeroDao() HeroDao {
	return HeroDao{}
}

func (heroDao HeroDao) FindAll() ([]model.Hero, error) {
	var all []model.Hero
	err := database.DBConn.Find(&all).Error
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (heroDao HeroDao) FindHeroByName(heroName string) (model.Hero, error) {
	var hero model.Hero
	err := database.DBConn.Find(&hero).Where("name = ?", heroName)
	return hero, err.Error
}

func (heroDao HeroDao) FindHeroByID(id uint) (model.Hero, error) {
	var hero model.Hero
	err := database.DBConn.Find(&hero).Where("id = ?", id)
	return hero, err.Error
}

func (heroDao HeroDao) SaveHero(hero model.Hero) error {
	err := database.DBConn.Save(hero).Error
	return err
}

func (heroDao HeroDao) DeleteHero(id uint) error {
	err := database.DBConn.Unscoped().Delete(id).Error
	return err
}
