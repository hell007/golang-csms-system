package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/goods"
)

type CategoryService interface {
	List(pid int) ([]models.Category, error)
	Get(id int) (*models.Category, error)
	Update(category *models.Category, columns []string) (int64, error)
	Create(category *models.Category) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type categoryService struct {
	dao *dao.CategoryDao
}

func NewCategoryService() CategoryService {
	return &categoryService{
		dao: dao.NewCategoryDao(db.MasterEngine()),
	}
}

func (s *categoryService) List(pid int) ([]models.Category, error) {
	return s.dao.List(pid)
}

func (s *categoryService) Get(id int) (*models.Category, error) {
	return s.dao.Get(id)
}

func (s *categoryService) Update(category *models.Category, columns []string) (int64, error) {
	return s.dao.Update(category, columns)
}

func (s *categoryService) Create(category *models.Category) (int64, error) {
	return s.dao.Create(category)
}

func (s *categoryService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *categoryService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}