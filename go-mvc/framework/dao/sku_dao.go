package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/goods"
)

type SkuDao struct {
	engine *xorm.Engine
}

func NewSkuDao(engine *xorm.Engine) *SkuDao {
	return &SkuDao{
		engine: engine,
	}
}

// update
func (d *SkuDao) UpdateVal(val *models.GoodsSkuVal, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(val.Vid).MustCols(columns...).Update(val)
	} else {
		effect, err = d.engine.Id(val.Vid).AllCols().Update(val)
	}
	return effect, err
}

// insert
func (d *SkuDao) CreateVal(val *models.GoodsSkuVal) (int64, error) {
	effect, err := d.engine.Insert(val)
	return effect, err
}

// update
func (d *SkuDao) UpdateKey(key *models.GoodsSkuKey, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(key.Kid).MustCols(columns...).Update(key)
	} else {
		effect, err = d.engine.Id(key.Kid).AllCols().Update(key)
	}
	return effect, err
}

// insert
func (d *SkuDao) CreateKey(key *models.GoodsSkuKey) (int64, error) {
	effect, err := d.engine.Insert(key)
	return effect, err
}
