package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/goods"
)

type CategoryDao struct {
	engine *xorm.Engine
}

func NewCategoryDao(engine *xorm.Engine) *CategoryDao {
	return &CategoryDao{
		engine: engine,
	}
}

// list
func (d *CategoryDao) List(pid int) ([]models.Category, error) {
	var (
		err  error
		list = make([]models.Category, 0)
	)

	if pid == 0 {
		err = d.engine.Cols("id", "category_name").Where("status = ?", 1).And("pid = ?", 0).Find(&list)
	} else {
		err = d.engine.Table("jie_category").Alias("C").Select("C.*, GA.small").
			Join("LEFT", "jie_goods G", "C.id = G.category_id").
			Join("LEFT", "jie_goods_gallery GA", "G.id = GA.goods_id").
			Where("status = ?", 1).GroupBy("C.category_name").Find(&list)
	}

	return list, err
}

// Get
func (d *CategoryDao) Get(id int) (*models.Category, error) {
	data := &models.Category{Id: id}
	_, err := d.engine.Omit("small").Get(data)
	return data, err
}

// update
func (d *CategoryDao) Update(category *models.Category, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(category.Id).MustCols(columns...).Update(category)
	} else {
		effect, err = d.engine.Id(category.Id).AllCols().Update(category)
	}
	return effect, err
}

// insert
func (d *CategoryDao) Create(category *models.Category) (int64, error) {
	effect, err := d.engine.Insert(category)
	return effect, err
}

// delete
func (d *CategoryDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	category := new(models.Category)
	effect, err = d.engine.In("id", ids).Delete(category)
	return effect, err
}

// close
func (d *CategoryDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	category := new(models.Category)
	category.Status = 2
	effect, err = d.engine.In("id", ids).Cols("status").Update(category)
	return effect, err
}
