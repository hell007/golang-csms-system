package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/order"
	"go-mvc/framework/utils/page"
)

type OrderService interface {
	List(name string, orderState int, p *page.Pagination) ([]models.Order, int64, error)
	Get(ordersn string) (*models.OrderDetail, error)
	GetOrderListByMid(mid int, orderState int, p *page.Pagination) ([]models.Orders, int64, error)
	Update(order *models.Order, columns []string) (int64, error)
	Create(order *models.Order, goods []models.OrderGoods) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type orderService struct {
	dao *dao.OrderDao
}

func NewOrderService() OrderService {
	return &orderService{
		dao: dao.NewOrderDao(db.MasterEngine()),
	}
}

func (s *orderService) List(name string, orderState int, p *page.Pagination) ([]models.Order, int64, error) {
	return s.dao.List(name, orderState, p)
}

func (s *orderService) Get(ordersn string) (*models.OrderDetail, error) {
	return s.dao.Get(ordersn)
}

func (s *orderService) GetOrderListByMid(mid int, orderState int, p *page.Pagination) ([]models.Orders, int64, error) {
	return s.dao.GetOrderListByMid(mid, orderState, p)
}

func (s *orderService) Update(order *models.Order, columns []string) (int64, error) {
	return s.dao.Update(order, columns)
}

func (s *orderService) Create(order *models.Order, goods []models.OrderGoods) (int64, error) {
	return s.dao.Create(order, goods)
}

func (s *orderService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *orderService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}
