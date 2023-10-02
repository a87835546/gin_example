package service

import (
	"gin_example/logic"
	"gin_example/model"
	"log"
)

type ActorService struct {
}

func (as *ActorService) Update(p *model.ActorModel) error {
	err := logic.Db.Table("actor").Updates(p).Error
	return err
}
func (as *ActorService) Insert(p *model.ActorModel) error {
	err := logic.Db.Table("actor").Create(p).Error
	return err
}

func (as *ActorService) QueryAll() (list []*model.ActorModel, err error) {
	rows, err := logic.Db.Table("actor").Rows()
	for rows.Next() {
		var l *model.ActorModel
		err = logic.Db.ScanRows(rows, &l)
		if err != nil {
			log.Println(err)
		}
		list = append(list, l)
	}
	return
}

func (as *ActorService) QueryByName(title string) (m *model.ActorModel, err error) {
	err = logic.Db.Debug().Table("actor").Where("name=?", title).First(&m).Error
	return
}
