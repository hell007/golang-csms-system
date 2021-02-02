package controllers

import (
	"go-mvc/framework/logs"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"

	models "go-mvc/framework/models/order"
	"go-mvc/framework/services"
	"go-mvc/framework/utils/page"
	"go-mvc/framework/utils/response"
)

type OrderController struct {
	Ctx     iris.Context
	Service services.OrderService
}

// list
func (c *OrderController) GetList() {
	var (
		err        error
		name       string
		orderState int
		p          *page.Pagination
		res        *page.Result
		list       []models.Order
		total      int64
	)

	// 分页设置
	p, err = page.NewPagination(c.Ctx)
	if err != nil {
		goto FAIL
	}

	// 查询
	name = c.Ctx.URLParam("name")
	orderState, _ = c.Ctx.URLParamInt("orderState")
	list, total, err = c.Service.List(name, orderState, p)

	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
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

// item
func (c *OrderController) GetItem() {
	var (
		err         error
		ordersn     string
		orderDetail = new(models.OrderDetail)
	)

	// 参数处理
	ordersn = c.Ctx.URLParam("ordersn")
	if err != nil {
		goto FAIL
	}

	// 查询
	orderDetail, err = c.Service.Get(ordersn)
	if err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "查询失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, orderDetail)
	return

	// 参数错误
FAIL:
	logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
	response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
	return
}

// save
func (c *OrderController) PostSave() {
	var (
		err     error
		effect  int64
		columns []string
		order   = new(models.Order)
	)

	if err = c.Ctx.ReadJSON(&order); err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, response.ParseParamsFailur)
		response.Error(c.Ctx, iris.StatusBadRequest, response.ParseParamsFailur, nil)
		return
	}

	if order.Id > 0 {
		order.ShipTime = time.Now()
		order.OrderState = 3 //已发货
		columns = append(columns, "ship_name", "ship_no", "ship_note", "ship_time", "order_state")
		effect, err = c.Service.Update(order, columns)
	}

	if effect < 0 || err != nil {
		logs.GetLogger().Error(logs.D{"err": err}, "保存失败")
		response.Failur(c.Ctx, response.OptionFailur, nil)
		return
	}

	response.Ok(c.Ctx, response.OptionSuccess, nil)
	return
}

//  delete?id=1,2
func (c *OrderController) GetDelete() {
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

// close?id=1,2
func (c *OrderController) GetClose() {
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
		logs.GetLogger().Error(logs.D{"err": err}, "关闭失败")
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
