package service

import (
	"gin_example/logic"
	"gin_example/models"
)

type ActorService struct {
}

func (as *ActorService) Update(p *models.ActorModel) error {
	err := logic.Db.Table("actor").Updates(p).Error
	return err
}
func (as *ActorService) Insert(p *models.ActorModel) error {
	err := logic.Db.Table("actor").Create(p).Error
	return err
}

func (as *ActorService) QueryAll() (list []*models.ActorModel, err error) {
	//rows,err := logic.Db.Table("actor").Find().Rows()
	return
}

func (as *ActorService) QueryByName(title string) (m *models.ActorModel, err error) {
	err = logic.Db.Debug().Table("actor").Where("name=?", title).First(&m).Error
	return
}
