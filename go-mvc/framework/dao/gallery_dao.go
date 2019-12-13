package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/goods"
)

type GalleryDao struct {
	engine *xorm.Engine
}

func NewGalleryDao(engine *xorm.Engine) *GalleryDao {
	return &GalleryDao{
		engine: engine,
	}
}

// insert
func (d *GalleryDao) Create(gallery *models.GoodsGallery) (int64, error) {
	effect, err := d.engine.Insert(gallery)
	return effect, err
}

// delete
func (d *GalleryDao) Delete(id string) (int64, error) {
	var (
		effect int64
		err    error
	)
	gallery := new(models.GoodsGallery)
	effect, err = d.engine.Where("gallery_id = ?", id).Delete(gallery)
	return effect, err
}
