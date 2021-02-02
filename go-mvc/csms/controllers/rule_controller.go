package controllers

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"strconv"
	"strings"

	"go-mvc/framework/middleware/casbin"
	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type RuleController struct {
	Ctx     iris.Context
	Service services.RuleService
}

// 规则列表
func (c *RuleController) GetList() {
	var (
		err   error
		p     *page.Pagination
		res   *page.Result
		list  []models.CasbinRule
		total int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	list, total, err = c.Service.List(p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询失败")
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Total: total,
		Rows:  list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 创建规则
func (c *RuleController) PostCreate() {
	var (
		err  error
		rule = new(models.CasbinRule)
	)

	// 读取
	if err = c.Ctx.ReadJSON(&rule); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	e := casbin.GetEnforcer()
	ok, err := e.AddPolicy(rule.V0, rule.V1, rule.V2, rule.V3, rule.V4, rule.V5)
	if !ok || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "添加失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
	}

	response.Ok_(c.Ctx, response.OptionSuccess)
	return
}

// 保存规则
func (c *RuleController) PostSave() {

	rule := models.CasbinRule{}

	// 读取
	if err := c.Ctx.ReadJSON(&rule); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err2 := c.Service.Update(&rule, nil)
	if effect < 0 || err2 != nil {
		logs.GetLogger().Error(logs.D{"err": err2}, "操作失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok_(c.Ctx, response.OptionSuccess)
	return
}

// 删除规则
func (c *RuleController) GetDelete() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int64, 0)
		uid    int64
		effect int64
	)

	id = c.Ctx.URLParam("id")
	idList = strings.Split(id, ",")
	if len(idList) == 0 {
		goto FAIL
	}

	for _, v := range idList {
		if v == "" {
			continue
		}

		uid, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Delete(ids)

	if effect <= 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "删除失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// RelationUserRole 给用户指定角色
func (c *RuleController) PostRelationuserrole() {
	groupDef := new(casbin.GroupDefine)

	if err := c.Ctx.ReadJSON(groupDef); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.ParseParamsFailur, nil)
		return
	}

	// TODO 校验前端的角色是否正确，和数据库的所有角色比较
	ok := true
	e := casbin.GetEnforcer()

	for _, v := range groupDef.Sub {
		// 给目标用户添加角色
		yes, _ := e.AddGroupingPolicy(groupDef.UID, v)
		if !yes {
			ok = false
		}
	}

	if !ok {
		logs.GetLogger().Error(nil, "给目标用户添加角色出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok_(c.Ctx, response.OptionSuccess)
}

// RoleUserList  角色用户查询
func (c *RuleController) GetRoleuserlist() {
	rKey := c.Ctx.URLParam("rKey")
	p, err := page.NewPagination(c.Ctx)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	// 获取角色名称 casbinrule v1(rkey) -> v0([]string)
	e := casbin.GetEnforcer()
	users, err := e.GetUsersForRole(rKey)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "获取用户角色名称出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 得到角色对应id
	roles := services.NewRoleService().GetAll()
	rids := make([]int, 0)
	for _, vr := range roles {
		for _, vu := range users {
			if vr.RoleName == vu {
				rids = append(rids, vr.Id)
			}
		}
	}

	// 根据角色id查询用户
	list, total, err := services.NewUserService().GetUsersByRids(rids, p)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "获取角色关联的用户表出错")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res := &page.Result{
		Total: total,
		Rows:  list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return
}

// 角色菜单
func (c *RuleController) GetMenus() {
	rid, err := c.Ctx.URLParamInt64("rid")
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	menus, err2 := services.NewMenuService().GetMenusByRoleid(rid)
	if err2 != nil {
		logs.GetLogger().Error(logs.D{"err": err2}, "查询失败")
		response.Failur(c.Ctx, response.OptionFailur, err2)
		return
	}

	// 组装数据
	res := &page.Result{
		Rows: menus,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
	return
}
