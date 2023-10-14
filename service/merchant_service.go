package service

import (
	"gin_example/logic"
	"gin_example/model"
)

type MerchantService struct {
}

type NewRuleService struct {
}

func (ms *NewRuleService) CreateInBatches(userModel []*model.RuleModel) (err error) {
	err = logic.RuleDb.Table("rule").CreateInBatches(userModel, len(userModel)).Error
	return
}

type NewRuleUserService struct {
}

func (ms *NewRuleUserService) CreateInBatches(userModel []*model.RuleUserModel) (err error) {
	err = logic.RuleDb.Table("user_rule").CreateInBatches(userModel, len(userModel)).Error
	return
}
