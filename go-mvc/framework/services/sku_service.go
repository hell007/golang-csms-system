package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/goods"
)

type SkuService interface {
	UpdateVal(val *models.GoodsSkuVal, columns []string) (int64, error)
	CreateVal(val *models.GoodsSkuVal) (int64, error)
	UpdateKey(key *models.GoodsSkuKey, columns []string) (int64, error)
	CreateKey(key *models.GoodsSkuKey) (int64, error)
}

type skuService struct {
	dao *dao.SkuDao
}

func NewSkuService() SkuService {
	return &skuService{
		dao: dao.NewSkuDao(db.MasterEngine()),
	}
}

func (s *skuService) UpdateVal(val *models.GoodsSkuVal, columns []string) (int64, error) {
	return s.dao.UpdateVal(val, columns)
}

func (s *skuService) CreateVal(val *models.GoodsSkuVal) (int64, error) {
	return s.dao.CreateVal(val)
}

func (s *skuService) UpdateKey(key *models.GoodsSkuKey, columns []string) (int64, error) {
	return s.dao.UpdateKey(key, columns)
}

// insert
func (s *skuService) CreateKey(key *models.GoodsSkuKey) (int64, error) {
	return s.dao.CreateKey(key)
}
