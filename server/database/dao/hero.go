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
