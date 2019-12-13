package controllers

import (
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"

	models "go-mvc/framework/models/goods"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/response"
)

type CategoryController struct {
	Ctx     iris.Context
	Service services.CategoryService
}

// list
func (c *CategoryController) GetList() {
	var (
		err   error
		list  []models.Category
		pid int
	)

	pid, _ = c.Ctx.URLParamInt("pid")

	// 查询
	list, err  = c.Service.List(pid)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Category GetList 查询：[%s]", err)
		response.Error(c.Ctx, iris.StatusInternalServerError, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, list)
	return
}

// item
func (c *CategoryController) GetItem() {
	var (
		err  error
		id   int
		category = new(models.Category)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	category, err = c.Service.Get(id)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Category GetItem 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, category)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Category GetItem 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// save
func (c *CategoryController) PostSave() {
	var (
		err    error
		effect int64
		category   = models.Category{}
	)

	if err = c.Ctx.ReadJSON(&category); err != nil {
		c.Ctx.Application().Logger().Errorf("Category PostSave Json：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	if category.Id > 0 {
		effect, err = c.Service.Update(&category, nil)
	} else {
		effect, err = c.Service.Create(&category)
	}

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Category PostSave 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
}

// delete
func (c *CategoryController) GetDelete() {
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
		c.Ctx.Application().Logger().Errorf("Category PostDelete 删除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Category PostDelete 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// close
func (c *CategoryController) GetClose() {
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
		c.Ctx.Application().Logger().Errorf("Category GetClose 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Category GetClose 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

