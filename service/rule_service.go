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

type Tmp struct {
	Id          int `json:"id"`
	HasChildren int `json:"has_children"`
}

func (rs *RuleService) GetRules(sid int) (list []*model.RuleModel, err error) {
	err = logic.RuleDb.Table("rule").Debug().Where("super_id=?", sid).Find(&list).Error
	tmp := make([]*Tmp, 0)

	err = logic.RuleDb.Raw("select (count(id)>0) as has_children,super_id as id from rule where super_id in (select id from rule where super_id = ?) group by id", sid).Scan(&tmp).Error
	if len(tmp) > 0 {
		for i := 0; i < len(tmp); i++ {
			for j := 0; j < len(list); j++ {
				if list[j].Id == tmp[i].Id {
					list[j].HasChildren = tmp[i].HasChildren > 0
				}
			}
		}
	}
	return
}

func (rs *RuleService) GetRulesWithoutRoot() (list []*model.RuleModel, err error) {
	err = logic.RuleDb.Table("rule").Debug().Where("super_id!=?", 0).Find(&list).Error
	return
}
func (rs *RuleService) InsertRule(model model.RuleModel) (err error) {
	err = logic.RuleDb.Table("rule").Debug().Create(&model).Error
	return
}
func (rs *RuleService) UpdateRule(model model.RuleModel) (err error) {
	err = logic.RuleDb.Table("rule").Debug().Updates(&model).Error
	return
}

func (rs *RuleService) GetUsers() (model []*model.UserModel, err error) {
	err = logic.RuleDb.Table("user").Debug().Limit(10).Order("id desc").Find(&model).Error
	return
}
func (rs *RuleService) AddUser(model *model.UserModel) (err error) {
	err = logic.RuleDb.Table("user").Debug().Create(&model).Error
	return
}

func (rs *RuleService) UpdateUser(model *model.UserModel) (err error) {
	err = logic.RuleDb.Table("user").Debug().Updates(&model).Error
	return
}
