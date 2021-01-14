package dao

import (
	"github.com/go-xorm/xorm"
	models "go-mvc/framework/models/system"
)

type AccessDao struct {
	engine *xorm.Engine
}

func NewAccessDao(engine *xorm.Engine) *AccessDao {
	return &AccessDao{
		engine: engine,
	}
}

// list
func (d *AccessDao) List(rid int64) ([]models.RoleMenu, error) {
	var (
		err  error
		list = make([]models.RoleMenu, 0)
	)

	err = d.engine.Where("rid = ?", rid).Find(&list)
	return list, err
}

// insert
func (d *AccessDao) Create(rid int64, rms *models.RoleMenus) (int64, error) {
	var (
		effect int64
		rows   int64
		err    error
		rm     = models.RoleMenu{}
	)

	effect, err = d.engine.Where("rid = ?", rid).Delete(rm)
	if effect >= 0 {
		rows, err = d.engine.Insert(rms.List)
	}

	return rows, err
}

// 批量 insert
func (d *AccessDao) CreateRoleMenu(roleMenu ...*models.RoleMenu) (int64, error) {
	return d.engine.Insert(roleMenu)
}
