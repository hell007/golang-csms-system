package services

import (
	"go-mvc/framework/dao"
	models "go-mvc/framework/models/system"
	"go-mvc/framework/utils/db"
)

type MenuService interface {
	List(level int) ([]models.Menu, error)
	GetMenusByRoleid(rid int64) ([]models.Menu, error)
	Get(id int) (*models.Menu, error)
	Update(menu *models.Menu, columns []string) (int64, error)
	Create(menu *models.Menu) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type menuService struct {
	dao *dao.MenuDao
}

func NewMenuService() MenuService {
	return &menuService{
		dao: dao.NewMenuDao(db.MasterEngine()),
	}
}

func (s *menuService) List(level int) ([]models.Menu, error) {
	return s.dao.List(level)
}

func (s *menuService) GetMenusByRoleid(rid int64) ([]models.Menu, error) {
	return s.dao.GetMenusByRoleid(rid)
}

func (s *menuService) Get(id int) (*models.Menu, error) {
	return s.dao.Get(id)
}

func (s *menuService) Update(menu *models.Menu, columns []string) (int64, error) {
	return s.dao.Update(menu, columns)
}

func (s *menuService) Create(menu *models.Menu) (int64, error) {
	return s.dao.Create(menu)
}

func (s *menuService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *menuService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}
