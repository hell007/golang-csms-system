package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/goods"
	"go-mvc/framework/utils/page"
)

type GoodsService interface {
	List(name string, category int, isOnSale int, isFirst int, isHot int, p *page.Pagination) ([]models.GoodsDetail, int64, error)
	Get(id int) (*models.Goods, []models.GoodsGallery, []models.GoodsSkuKey, []models.GoodsSkuVal, error)
	Update(goods *models.Goods, columns []string) (int64, error)
	Create(goods *models.Goods) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
	//前端
	GetGoods(state string, val string, p *page.Pagination) ([]models.GoodsDetail, int64, error)
	GetProduct(id int) (*models.Product, error)
}

type goodsService struct {
	dao *dao.GoodsDao
}

func NewGoodsService() GoodsService {
	return &goodsService{
		dao: dao.NewGoodsDao(db.MasterEngine()),
	}
}

func (s *goodsService) List(name string, category int, isOnSale int, isFirst int, isHot int, p *page.Pagination) ([]models.GoodsDetail, int64, error) {
	return s.dao.List(name, category, isOnSale, isFirst, isHot, p)
}

func (s *goodsService) Get(id int) (*models.Goods, []models.GoodsGallery, []models.GoodsSkuKey, []models.GoodsSkuVal, error)  {
	return s.dao.Get(id)
}

func (s *goodsService) Update(goods *models.Goods, columns []string) (int64, error) {
	return s.dao.Update(goods, columns)
}

func (s *goodsService) Create(goods *models.Goods) (int64, error) {
	return s.dao.Create(goods)
}

func (s *goodsService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *goodsService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}

func (s *goodsService) GetGoods(state string, val string, p *page.Pagination) ([]models.GoodsDetail, int64, error){
	return s.dao.GetGoods(state, val, p)
}

func (s *goodsService) GetProduct(id int) (*models.Product, error) {
	return s.dao.GetProduct(id)
}



