package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/system"
	"go-mvc/framework/utils/page"
)

type RoleService interface {
	Exist(role *models.Role) (bool, error)
	GetAll() []models.Role
	List(name string, status int, p *page.Pagination) ([]models.Role, int64, error)
	Get(id int) (*models.Role, error)
	Update(role *models.Role, columns []string) (int64, error)
	Create(role *models.Role) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type roleService struct {
	dao *dao.RoleDao
}

func NewRoleService() RoleService {
	return &roleService{
		dao: dao.NewRoleDao(db.MasterEngine()),
	}
}

func (s *roleService) Exist(role *models.Role) (bool, error) {
	return s.dao.Exist(role)
}

func (s *roleService) GetAll() []models.Role {
	return s.dao.GetAll()
}

func (s *roleService) List(name string, status int, p *page.Pagination) ([]models.Role, int64, error) {
	return s.dao.List(name, status, p)
}

func (s *roleService) Get(id int) (*models.Role, error) {
	return s.dao.Get(id)
}

func (s *roleService) Update(role *models.Role, columns []string) (int64, error) {
	return s.dao.Update(role, columns)
}

func (s *roleService) Create(role *models.Role) (int64, error) {
	return s.dao.Create(role)
}

func (s *roleService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *roleService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}