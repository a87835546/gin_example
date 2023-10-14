package service

import (
	"gin_example/logic"
	"gin_example/model"
)

type PermissionService struct {
}

func (ps *PermissionService) Add(model *model.NewPermissionModel) (err error) {
	err = logic.RuleDb.Table("permission").Debug().Create(&model).Error
	return
}
func (ps *PermissionService) Update(model *model.NewPermissionModel) (err error) {
	err = logic.RuleDb.Table("permission").Debug().Updates(&model).Error
	return
}
func (ps *PermissionService) QueryAll() (model []*model.NewPermissionModel, err error) {
	err = logic.RuleDb.Table("permission").Debug().Find(&model).Error
	return
}
