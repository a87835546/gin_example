package controllers

import (
	"gin_example/doreamon/param"
	"gin_example/doreamon/utils"
	"gin_example/model"
	"gin_example/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RuleController struct {
	rs service.RuleService
}

func NewRuleController() *RuleController {
	return &RuleController{
		rs: service.RuleService{},
	}
}
func (rc *RuleController) AddPolicy(ctx *gin.Context) {
	req := &param.AddPolicyReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		res := rc.rs.AddPolicy(req)
		RespOk(ctx, res)
	}
}

func (rc *RuleController) AddGroupPolicy(ctx *gin.Context) {
	req := &param.AddGroupPolicyReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		res := rc.rs.AddGroupPolicy(req)
		if res {
			RespOk(ctx, res)
		} else {
			RespErrorWithMsg(ctx, utils.AddPermissionErrorCode, "add grouping policy error", nil)
		}
	}
}

func (rc *RuleController) AddNamedGroupPolicy(ctx *gin.Context) {
	req := &param.AddGroupPolicyReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		res := rc.rs.AddGroupPolicy(req)
		RespOk(ctx, res)
	}
}

func (rc *RuleController) GetPolicy(ctx *gin.Context) {
	title := ctx.Query("title")
	RespOk(ctx, rc.rs.GetPolicy(title))
}
func (rc *RuleController) GetGroupPolicy(ctx *gin.Context) {
	title := ctx.Query("group")
	RespOk(ctx, rc.rs.GetGroupPolicy(title))
}
func (rc *RuleController) GetNamedGroupingPolicy(ctx *gin.Context) {
	RespOk(ctx, rc.rs.GetNamedGroupingPolicy())
}
func (rc *RuleController) GetGroupPolicyByRule(ctx *gin.Context) {
	title := ctx.Query("rule")
	list, err := rc.rs.GetGroupPolicyByRule(title)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}
func (rc *RuleController) DeleteGroupPolicy(ctx *gin.Context) {
	title := ctx.Query("title")
	RespOk(ctx, rc.rs.DeleteGroupPolicy(title))
}

func (rc *RuleController) GetRules(ctx *gin.Context) {
	sid := ctx.Query("super_id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		id = 0
	}
	list, err := rc.rs.GetRules(id)
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}

func (rc *RuleController) GetRulesWithoutRoot(ctx *gin.Context) {
	list, err := rc.rs.GetRulesWithoutRoot()
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}
func (rc *RuleController) UpdateRule(ctx *gin.Context) {
	req := model.RuleModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = rc.rs.UpdateRule(req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.UpdateDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
func (rc *RuleController) InsertRule(ctx *gin.Context) {
	req := model.RuleModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = rc.rs.InsertRule(req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
func (rc *RuleController) GetUsers(ctx *gin.Context) {
	list, err := rc.rs.GetUsers()
	if err != nil {
		RespErrorWithMsg(ctx, utils.QueryDBErrorCode, err.Error(), nil)
	} else {
		RespOk(ctx, list)
	}
}
func (rc *RuleController) AddUsers(ctx *gin.Context) {
	req := model.UserModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = rc.rs.AddUser(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
func (rc *RuleController) UpdateUser(ctx *gin.Context) {
	req := model.UserModel{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		RespErrorWithMsg(ctx, utils.ParameterErrorCode, err.Error(), nil)
	} else {
		err = rc.rs.UpdateUser(&req)
		if err != nil {
			RespErrorWithMsg(ctx, utils.InsertDBErrorCode, err.Error(), nil)
		} else {
			RespOk(ctx, nil)
		}
	}
}
