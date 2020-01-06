package goods

// sku-key
type GoodsSkuKey struct {
	Kid     int    `json:"kid" xorm:"not null pk autoincr comment('主键id') INT(10)"`
	GoodsSn string `json:"goodsSn,omitempty" xorm:"not null comment('商品编码') VARCHAR(20)"`
	Name    string `json:"name,omitempty" xorm:"not null comment('sku规格名称') VARCHAR(20)"`
	Keys    string `json:"keys,omitempty" xorm:"not null comment('sku规格选项') VARCHAR(200)"`
}
