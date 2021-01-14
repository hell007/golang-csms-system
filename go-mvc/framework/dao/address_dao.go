package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/member"
	"go-mvc/framework/utils/page"
)

type AddressDao struct {
	engine *xorm.Engine
}

func NewAddressDao(engine *xorm.Engine) *AddressDao {
	return &AddressDao{
		engine: engine,
	}
}

// GetAll
func (d *AddressDao) GetAll() []models.Address {
	datalist := make([]models.Address, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *AddressDao) List(p *page.Pagination) ([]models.Address, int64, error) {

	list := make([]models.Address, 0)

	s := d.engine.Limit(p.Limit, p.Start)

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *AddressDao) Get(id int) (*models.Address, error) {
	data := &models.Address{Id: id}
	_, err := d.engine.Get(data)
	return data, err
}

// update
func (d *AddressDao) Update(address *models.Address, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(address.Id).Cols(columns...).Update(address)
	} else {
		effect, err = d.engine.Id(address.Id).AllCols().Update(address)
	}
	return effect, err
}

// insert
func (d *AddressDao) Create(address *models.Address) (int64, error) {
	effect, err := d.engine.Insert(address)
	return effect, err
}

// delete
func (d *AddressDao) Delete(id int) (int64, error) {
	var (
		effect int64
		err    error
	)

	address := new(models.Address)
	effect, err = d.engine.Where("id = ?", id).Delete(address)
	return effect, err
}
