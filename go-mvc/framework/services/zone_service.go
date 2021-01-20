package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/zone"
	"go-mvc/framework/utils/page"
)

type ZoneService interface {
	GetAll() ([]models.Zone, error)
	List(p *page.Pagination) ([]models.Zone, int64, error)
	GetCity(pid int) ([]models.City, error)
	Get(id int) (*models.Zone, error)
	Delete(id int) (int64, error)
	Update(zone *models.Zone, columns []string) (int64, error)
	Create(zone *models.Zone) (int64, error)
}

type zoneService struct {
	dao *dao.ZoneDao
}

func NewZoneService() ZoneService {
	return &zoneService{
		dao: dao.NewZoneDao(db.MasterEngine()),
	}
}

func (s *zoneService) GetAll() ([]models.Zone, error) {
	return s.dao.GetAll()
}

func (s *zoneService) List(p *page.Pagination) ([]models.Zone, int64, error) {
	return s.dao.List(p)
}

func (s *zoneService) GetCity(pid int) ([]models.City, error) {
	return s.dao.GetCity(pid)
}

func (s *zoneService) Get(id int) (*models.Zone, error) {
	return s.dao.Get(id)
}

func (s *zoneService) Update(address *models.Zone, columns []string) (int64, error) {
	return s.dao.Update(address, columns)
}

func (s *zoneService) Create(address *models.Zone) (int64, error) {
	return s.dao.Create(address)
}

func (s *zoneService) Delete(id int) (int64, error) {
	return s.dao.Delete(id)
}
