package dao

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"go-mvc/framework/utils/uuids"

	members "go-mvc/framework/models/member"
	models "go-mvc/framework/models/order"
	"go-mvc/framework/utils/page"
)

type OrderDao struct {
	engine *xorm.Engine
}

func NewOrderDao(engine *xorm.Engine) *OrderDao {
	return &OrderDao{
		engine: engine,
	}
}

// List
func (d *OrderDao) List(name string, orderState int, p *page.Pagination) ([]models.Order, int64, error) {
	var (
		list = make([]models.Order, 0)
	)

	s := d.engine.Table("jie_order").Alias("o").Select("o.id, o.ordersn, o.order_state, " +
			"o.create_time, o.pay_time, o.ship_time, o.confirm_time, o.pay_name, " +
			"o.total_price, o.ship_price, o.pay_price, o.order_price, o.return_price, " +
			"o.pay_name, o.ship_name, o.ship_no, o.ship_note, o.consignee, o.phone, o.tell, o.remark, " +
			"(SELECT area_name from jie_zone WHERE province = area_id) as province, " +
			"(SELECT area_name from jie_zone WHERE city = area_id) as city, " +
			"(SELECT area_name from jie_zone WHERE district = area_id) as district, o.address ")

	if name != "" {
		s.Where("o.ordersn like ?", "%"+name+"%")
	}

	if orderState > 0 {
		s.And("o.order_state = ?", orderState)
	}

	s.Limit(p.Limit, p.Start)

	if p.SortName != "" {
		switch p.SortOrder {
		case "asc":
			s.Asc(p.SortName)
		case "desc":
			s.Desc(p.SortName)
		}
	}

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *OrderDao) Get(ordersn string) (*models.OrderDetail, error) {
	var (
		err         error
		ok          bool
		order       = new(models.Order)
		member      = new(members.Member)
		orderDetail = new(models.OrderDetail)
		list        = make([]models.OrderGoodsDetail, 0)
	)

	orderSql := fmt.Sprintf(`SELECT  o.id, o.ordersn, o.order_state, 
		o.create_time, o.pay_time, o.ship_time, o.confirm_time,
		o.pay_name, o.ship_name,  o.ship_no, o.ship_note,
		o.total_price, o.ship_price, o.pay_price, o.order_price, o.return_price,
		o.consignee, o.phone, o.tell, o.remark,
		(SELECT area_name from jie_zone WHERE province = area_id) as province,
		(SELECT area_name from jie_zone WHERE city = area_id) as city,
		(SELECT area_name from jie_zone WHERE district = area_id) as district,
		o.address FROM jie_order o WHERE o.ordersn = %s`, ordersn)
	ok, err = d.engine.SQL(orderSql).Get(order)

	if ok && err == nil {
		orderDetail.Order = order

		ogSql := fmt.Sprintf(`SELECT OG.*, G.goods_sn, G.goods_name, G.color, G.unit, GA.small 
        FROM jie_order O 
		LEFT JOIN jie_order_goods OG ON O.id = OG.order_id 
		LEFT JOIN jie_goods G ON OG.goods_id = G.id 
		LEFT JOIN (SELECT goods_id, small from jie_goods_gallery group by goods_id) as GA ON OG.goods_id = GA.goods_id 
		WHERE O.id = %d`, order.Id)
		err = d.engine.SQL(ogSql).Find(&list)
		orderDetail.List = list

		mSql := fmt.Sprintf(`SELECT M.name, M.mobile FROM jie_order O 
		LEFT JOIN jie_member M ON O.mid = M.id 
		WHERE O.id = %d`, order.Id)
		_, err = d.engine.SQL(mSql).Get(member)
		orderDetail.Member = member
	}
	return orderDetail, err
}

// getOrderByMid
// 查询：一个用户对应多个订单，一个订单对应多个商品
func (d *OrderDao) GetOrderListByMid(mid int, orderState int, p *page.Pagination) ([]models.Orders, int64, error) {
	var (
		olist = make([]models.Orders, 0)
		glist = make([]models.OrderGoodsDetail, 0)
		ids   []int
	)

	s := d.engine.Table("jie_order").Cols("id", "ordersn", "order_state", "total_price", "ship_price", "pay_price", "order_price").
		Where("mid = ?", mid).And("DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(create_time)")

	if orderState > 0 {
		s.And("order_state = ?", orderState)
	}

	s.Limit(p.Limit, p.Start)
	count, err := s.FindAndCount(&olist)

	// 二次查询订单对应的商品
	if count > 0 {
		for _, v := range olist {
			ids = append(ids, v.Id)
		}

		err := d.engine.Table("jie_order").Alias("O").
			Select("OG.*, G.goods_sn, G.goods_name, G.color, G.unit, GA.small").
			Join("LEFT", "jie_order_goods OG", "O.id = OG.order_id").
			Join("LEFT", "jie_goods G", "OG.goods_id = G.id").
				Join("LEFT", "(SELECT goods_id, small from jie_goods_gallery group by goods_id) as GA",
					"OG.goods_id = GA.goods_id").In("O.id", ids).Find(&glist)

		// 组装数据
		// 根据order_id 合并多个商品到每个订单
		if err == nil {
			for k1, v1 := range olist {

				for k2, v2 := range glist {

					if v1.Id == v2.OrderId {
						olist[k1].List = append(olist[k1].List, glist[k2])
					}
				}
			}
		}
	}

	return olist, count, err
}

// update
func (d *OrderDao) Update(order *models.Order, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(order.Id).Cols(columns...).Update(order)
	} else {
		effect, err = d.engine.Id(order.Id).AllCols().Update(order)
	}
	return effect, err
}

// insert
func (d *OrderDao) Create(order *models.Order, goods []models.OrderGoods) (int64, error) {
	var (
		err error
	)

	// 事务处理开始
	session := d.engine.NewSession()
	defer session.Close()

	if err = session.Begin(); err != nil {
		return 0, errors.New("订单事务处理：")
	}

	// 新增订单
	if _, err = session.Insert(order); err != nil {
		err = session.Rollback()
		return 0, errors.New("订单新增失败")
	}

	// 设置新增成功返回的主键id
	for k, _ := range goods {
		goods[k].Id = uuids.UUid32()
		goods[k].OrderId = order.Id
	}

	// 新增订单商品
	if _, err = session.Table("jie_order_goods").Insert(goods); err != nil {
		err = session.Rollback()
		return 0, errors.New("订单商品新增失败")
	}

	err = session.Commit()
	return 1, err
}

// delete
func (d *OrderDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	order := new(models.Order)
	effect, err = d.engine.In("id", ids).Delete(order)
	return effect, err
}

// close
func (d *OrderDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	order := new(models.Order)
	order.OrderState = 4 //取消
	effect, err = d.engine.In("id", ids).Cols("order_state").Update(order)
	return effect, err
}
