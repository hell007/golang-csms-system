package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/system"
	"fmt"
)

type MenuDao struct {
	engine *xorm.Engine
}

func NewMenuDao(engine *xorm.Engine) *MenuDao {
	return &MenuDao{
		engine: engine,
	}
}

func (d *MenuDao) List(level int) ([]models.Menu, error) {
	var (
		err  error
		list = make([]models.Menu, 0)
	)

	if level > 0 {
		err = d.engine.Where("status = ?", 1).In("level", 1, 2).Find(&list)
	} else {
		err = d.engine.Where("status = ?", 1).Find(&list)
	}

	return list, err
}

func (d *MenuDao) GetMenusByRoleid(rid int64) ([]models.Menu, error) {
	var (
		err   error
		sql   string
		menus = make([]models.Menu, 0)
	)

	if rid == 1 {
		err = d.engine.Where("hidden = ?", 1).And("status = ?", 1).Find(&menus)

	} else {
		sql = fmt.Sprintf(`
				SELECT M.id, M.name, M.pid, M.path, M.level, M.redirect, M.hidden, M.icon FROM jie_menu M
				WHERE M.hidden = 1 AND M.status = 1 AND M.id in
				(
				SELECT mid FROM jie_role_menu WHERE rid=%d
				)
				`, rid)
		err = d.engine.SQL(sql).Find(&menus)
	}

	return menus, err
}

// Get
func (d *MenuDao) Get(id int) (*models.Menu, error) {
	data := &models.Menu{Id: id}
	_, err := d.engine.Get(data)
	return data, err
}

// update
func (d *MenuDao) Update(menu *models.Menu, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(menu.Id).MustCols(columns...).Update(menu)
	} else {
		effect, err = d.engine.Id(menu.Id).AllCols().Update(menu)
	}
	return effect, err
}

// insert
func (d *MenuDao) Create(menu *models.Menu) (int64, error) {
	effect, err := d.engine.Insert(menu)
	return effect, err
}

// delete
func (d *MenuDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	menu := new(models.Menu)
	effect, err = d.engine.In("id", ids).Delete(menu)
	return effect, err
}

// close
func (d *MenuDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	menu := new(models.Menu)
	menu.Status = 2
	effect, err = d.engine.In("id", ids).Cols("status").Update(menu)
	return effect, err
}
