package controllers

import (
	"github.com/kataras/iris/v12"
	"strconv"
	"strings"
	"time"

	//"go-mvc/framework/middleware/jwt"
	models "go-mvc/framework/models/member"
	"go-mvc/framework/services"
	//"go-mvc/framework/utils/encrypt"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type MemberController struct {
	Ctx     iris.Context
	Service services.MemberService
}

// list?pageNumber=1&pageSize=2&name=曹操
func (c *MemberController) GetList() {
	var (
		err      error
		name string
		status   int
		p        *page.Pagination
		res      *page.Result
		list     []models.Member
		total    int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	name = c.Ctx.URLParam("name")
	status, _ = c.Ctx.URLParamInt("status")
	list, total, err = c.Service.List(name, status, p)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Member GetList 查询：[%s]", err)
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
	c.Ctx.Application().Logger().Errorf("Member GetList 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// item?id=1
func (c *MemberController) GetItem() {
	var (
		err  error
		id   int
		memberDetail = new(models.MemberDetail)
	)

	// 参数处理
	id, err = c.Ctx.URLParamInt("id")
	if err != nil {
		goto FAIL
	}

	// 查询
	memberDetail, err = c.Service.Get(id)
	if err != nil {
		c.Ctx.Application().Logger().Errorf("Member GetItem 查询：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, memberDetail)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Member GetItem 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// save
func (c *MemberController) PostSave() {
	var (
		err    error
		effect int64
		member   = new(models.Member)
	)

	if err = c.Ctx.ReadJSON(&member); err != nil {
		c.Ctx.Application().Logger().Errorf("Member PostSave Json：[%s]", err)
		response.Error(c.Ctx, iris.StatusBadRequest, response.OptionFailur, nil)
		return
	}

	if member.Id > 0 {
		member.Status = 3 //已注销
		effect, err = c.Service.Update(member, nil)
	} else {
		member.CreateTime = time.Now()
		effect, err = c.Service.Create(member)
	}

	if effect < 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Member PostSave 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

// delete?id=1,2
func (c *MemberController) GetDelete() {
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
		c.Ctx.Application().Logger().Errorf("Member GetDelete 删除：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Member GetDelete 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

func (c *MemberController) GetClose() {
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

	effect, err = c.Service.Close(ids)

	if effect <= 0 || err != nil {
		c.Ctx.Application().Logger().Errorf("Member GetClose 操作：[%s]", err)
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return

	// 参数错误
FAIL:
	c.Ctx.Application().Logger().Errorf("Member GetClose 参数：[%s]", err)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}
