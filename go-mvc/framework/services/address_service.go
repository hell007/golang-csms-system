package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/member"
	"go-mvc/framework/utils/page"
)

type AddressService interface {
	GetAll() []models.Address
	List(p *page.Pagination) ([]models.Address, int64, error)
	Get(id int) (*models.Address, error)
	Delete(id int) (int64, error)
	Update(address *models.Address, columns []string) (int64, error)
	Create(address *models.Address) (int64, error)
}

type addressService struct {
	dao *dao.AddressDao
}

func NewAddressService() AddressService {
	return &addressService{
		dao: dao.NewAddressDao(db.MasterEngine()),
	}
}

func (s *addressService) GetAll() []models.Address {
	return s.dao.GetAll()
}

func (s *addressService) List(p *page.Pagination) ([]models.Address, int64, error) {
	return s.dao.List(p)
}

func (s *addressService) Get(id int) (*models.Address, error) {
	return s.dao.Get(id)
}

func (s *addressService) Update(address *models.Address, columns []string) (int64, error) {
	return s.dao.Update(address, columns)
}

func (s *addressService) Create(address *models.Address) (int64, error) {
	return s.dao.Create(address)
}

func (s *addressService) Delete(id int) (int64, error) {
	return s.dao.Delete(id)
}
