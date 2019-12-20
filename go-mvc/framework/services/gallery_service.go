package services

import (
	"go-mvc/framework/dao"
	"go-mvc/framework/db"
	models "go-mvc/framework/models/goods"
)

type GalleryService interface {
	Create(gallery *models.GoodsGallery) (int64, error)
	Delete(id string) (int64, error)
}

type galleryService struct {
	dao *dao.GalleryDao
}

func NewGalleryService() GalleryService {
	return &galleryService{
		dao: dao.NewGalleryDao(db.MasterEngine()),
	}
}

func (s *galleryService) Create(gallery *models.GoodsGallery) (int64, error) {
	return s.dao.Create(gallery)
}

func (s *galleryService) Delete(id string) (int64, error) {
	return s.dao.Delete(id)
}
