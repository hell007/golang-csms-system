package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/utils/page"
)

type RoleDao struct {
	engine *xorm.Engine
}

func NewRoleDao(engine *xorm.Engine) *RoleDao {
	return &RoleDao{
		engine: engine,
	}
}

// exist
func (d *RoleDao) Exist(role *models.Role) (bool, error) {
	has, err := d.engine.Exist(role)
	return has, err
}

// GetAll
func (d *RoleDao) GetAll() []models.Role {
	datalist := make([]models.Role, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *RoleDao) List(name string, status int, p *page.Pagination) ([]models.Role, int64, error) {

	list := make([]models.Role, 0)

	s := d.engine.Omit("")

	if name != "" {
		s.Where("role_name like ?", "%"+name+"%")
	}

	if status > 0 {
		s.And("status = ?", status)
	}

	count, err := s.FindAndCount(&list)
	return list, count, err
}

// Get
func (d *RoleDao) Get(id int) (*models.Role, error) {
	data := &models.Role{Id: id}
	_, err := d.engine.Get(data)
	return data, err
}

// update
func (d *RoleDao) Update(role *models.Role, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(role.Id).MustCols(columns...).Update(role)
	} else {
		effect, err = d.engine.Id(role.Id).AllCols().Update(role)
	}
	return effect, err
}

// insert
func (d *RoleDao) Create(role *models.Role) (int64, error) {
	effect, err := d.engine.Insert(role)
	return effect, err
}

// delete
func (d *RoleDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	role := new(models.Role)
	effect, err = d.engine.In("id", ids).Delete(role)
	return effect, err
}

// close
func (d *RoleDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	role := new(models.Role)
	role.Status = 2
	effect, err = d.engine.In("id", ids).Cols("status").Update(role)
	return effect, err
}
