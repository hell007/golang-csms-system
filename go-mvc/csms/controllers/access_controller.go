package controllers

import (
	"github.com/kataras/iris/v12"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type AccessController struct {
	Ctx     iris.Context
	Service services.AccessService
}

//  访问列表
func (c *AccessController) GetList() {
	var (
		err  error
		rid  int64
		res  *page.Result
		list []models.RoleMenu
	)

	rid, err = c.Ctx.URLParamInt64("rid")
	if err != nil {
		goto FAIL
	}

	// 查询
	list, err = c.Service.List(rid)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("AccessController GetList 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	// 组装数据
	res = &page.Result{
		Rows: list,
	}

	response.Ok(c.Ctx, response.OptionSuccess, res)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("AccessController GetList 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// 访问保存
func (c *AccessController) PostSave() {
	var (
		err    error
		effect int64
		rms    models.RoleMenus
	)

	if err = c.Ctx.ReadJSON(&rms); err != nil {
		c.Ctx.Application().Logger().Errorf("AccessController PostSave Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	effect, err = c.Service.Create(rms.Rid, &rms)

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("AccessController PostSave 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}
