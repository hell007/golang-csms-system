package controllers

import (
	"github.com/kataras/iris/v12"
	"go-mvc/framework/logs"
	"strconv"
	"strings"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type PermissionController struct {
	Ctx     iris.Context
	Service services.MenuService
}

// 菜单列表
func (c *PermissionController) GetList() {
	var (
		err   error
		res   *page.Result
		list  []models.Menu
		level int
	)

	// 查询
	level, _ = c.Ctx.URLParamInt("level")
	if level < 0 {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	list, err = c.Service.List(level)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Rows: list,
	}
	response.Ok(c.Ctx, response.OptionSuccess, res)
}

// 菜单
func (c *PermissionController) GetItem() {
	var (
		err  error
		id   int
		menu = new(models.Menu)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	menu, err = c.Service.Get(id)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, menu)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 菜单保存
func (c *PermissionController) PostSave() {
	var (
		err    error
		effect int64
		menu   = models.Menu{}
	)

	if err = c.Ctx.ReadJSON(&menu); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	if menu.Id > 0 {
		effect, err = c.Service.Update(&menu, nil)
	} else {
		effect, err = c.Service.Create(&menu)
	}

	if effect < 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "保存失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

// 菜单删除
func (c *PermissionController) GetDelete() {
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

// 菜单停用
func (c *PermissionController) GetClose() {
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
		logs.GetLogger().Error(logs.D{"err": err}, "close失败")
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
