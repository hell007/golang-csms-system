package controllers

import (
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type RoleController struct {
	Ctx     iris.Context
	Service services.RoleService
}

// 角色列表
func (c *RoleController) GetList() {
	var (
		err      error
		rolename string
		status   int
		p        *page.Pagination
		res      *page.Result
		list     []models.Role
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	rolename = c.Ctx.URLParam("name")
	status, _ = c.Ctx.URLParamInt("status")
	list, total, err = c.Service.List(rolename, status, p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Role GetList 查询：[%s]", err)
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
	c.Ctx.Application().Logger().Errorf("Role GetList 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 角色
func (c *RoleController) GetItem() {
	var (
		err  error
		id   int
		role = new(models.Role)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	role, err = c.Service.Get(id)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Role GetItem 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, role)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Role GetItem 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 角色保存
func (c *RoleController) PostSave() {
	var (
		err    error
		effect int64
		role   = models.Role{}
	)

	if err = c.Ctx.ReadJSON(&role); err != nil {
		c.Ctx.Application().Logger().Errorf("Role PostSave Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	if role.Id > 0 {
		effect, err = c.Service.Update(&role, nil)
	} else {
		effect, err = c.Service.Create(&role)
	}

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Role PostSave 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}

// 角色删除
func (c *RoleController) GetDelete() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int, 0)
		uid    int
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

		uid, err = strconv.Atoi(v)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, uid)
	}

	effect, err = c.Service.Delete(ids)

	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Role PostDelete 删除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Role PostDelete 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 角色停用
func (c *RoleController) GetClose() {
	var (
		err    error
		id     string
		idList = make([]string, 0)
		ids    = make([]int, 0)
		rid    int
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

		rid, err = strconv.Atoi(v)
		if err != nil {
			goto FAIL
		}

		ids = append(ids, rid)
	}

	effect, err = c.Service.Close(ids)

	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Role GetClose 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Role GetClose 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}
