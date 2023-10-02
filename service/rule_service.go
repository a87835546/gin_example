package service

import (
	"gin_example/doreamon/param"
	"gin_example/logic"
	"gin_example/model"
	"log"
)

type RuleService struct {
}

func (rs *RuleService) AddPolicy(req *param.AddPolicyReq) bool {
	res, err := logic.E.AddPolicy(req.Name, req.Source, req.Permission)
	if err != nil {
		log.Printf("添加权限失败 --->>> %v\n", err.Error())
	}
	_ = logic.E.LoadPolicy()
	return res
}
func (rs *RuleService) AddGroupPolicy(req *param.AddGroupPolicyReq) bool {
	res, err := logic.E.AddGroupingPolicy(req.Name, req.Group)
	_ = logic.E.LoadPolicy()
	if err != nil && !res {
		log.Printf("添加权限失败 --->>> %v\n", err.Error())
	}
	return res
}
func (rs *RuleService) GetPolicy(title string) [][]string {
	res := logic.E.GetFilteredPolicy(0, title)
	log.Printf("res--->>%v", res)
	_ = logic.E.LoadPolicy()
	return res
}
func (rs *RuleService) GetGroupPolicy(title string) [][]string {
	res := logic.E.GetNamedPolicy(title)
	log.Printf("res--->>%v", res)
	_ = logic.E.LoadPolicy()
	return res
}
func (rs *RuleService) GetGroupPolicyByRule(title string) (list []*model.PermissionModel, err error) {
	res := logic.E.GetFilteredGroupingPolicy(0, title)
	//for i := 0; i < len(res); i++ {
	//	m := model.PermissionModel{}
	//	for j := 0; j < len(res[i]); j++ {
	//		l := strings.Split(res[i][j], " ")
	//		if len(l) > 1 {
	//			m.User = l[0]
	//			m.Value = l[1]
	//		}
	//	}
	//	list = append(list, &m)
	//}
	if len(title) > 0 {
		err = logic.Db.Table("casbin_rule").Where("v0=? and p_type = ?", title, "p").Find(&list).Error
	} else {
		err = logic.Db.Table("casbin_rule").Where("p_type = ?", "p").Find(&list).Error
	}
	log.Printf("res--->>%v", res)
	_ = logic.E.LoadPolicy()
	return
}
func (rs *RuleService) DeleteGroupPolicy(title string) bool {
	res, _ := logic.E.RemoveFilteredPolicy(0, title)
	logic.E.GetNamedGroupingPolicy("g")
	log.Printf("res--->>%v", res)
	_ = logic.E.LoadPolicy()
	return res
}
func (rs *RuleService) GetNamedGroupingPolicy() [][]string {
	res := logic.E.GetNamedGroupingPolicy("g")
	log.Printf("res--->>%v", res)
	_ = logic.E.LoadPolicy()
	return res
}
