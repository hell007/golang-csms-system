package goods

type GoodsGallery struct {
	GalleryId string `json:"galleryId,omitempty" xorm:"not null pk comment('商品画廊id') unique VARCHAR(64)"`
	Small     string `json:"small,omitempty" xorm:"not null comment('小图') VARCHAR(100)"`
	Medium    string `json:"medium,omitempty" xorm:"not null comment('中图') VARCHAR(100)"`
	Source    string `json:"source,omitempty" xorm:"not null comment('中图') VARCHAR(100)"`
	GoodsId   int    `json:"goodsId,omitempty" xorm:"not null comment('商品id') int(10)"`
}
