package services

import (
	"go-mvc/framework/dao"
	models "go-mvc/framework/models/system"
	"go-mvc/framework/utils/db"
)

type AccessService interface {
	List(rid int64) ([]models.RoleMenu, error)
	Create(rid int64, rms *models.RoleMenus) (int64, error)
	CreateRoleMenu(roleMenu ...*models.RoleMenu) (int64, error)
}

type accessService struct {
	dao *dao.AccessDao
}

func NewAccessService() AccessService {
	return &accessService{
		dao: dao.NewAccessDao(db.MasterEngine()),
	}
}

// list
func (s *accessService) List(rid int64) ([]models.RoleMenu, error) {
	return s.dao.List(rid)
}

// insert
func (s *accessService) Create(rid int64, rms *models.RoleMenus) (int64, error) {
	return s.dao.Create(rid, rms)
}

// 批量 insert
func (s *accessService) CreateRoleMenu(roleMenu ...*models.RoleMenu) (int64, error) {
	return s.dao.CreateRoleMenu(roleMenu...)
}
