package goods

// sku-val
type GoodsSkuVal struct {
	Vid   int    `json:"vid" xorm:"not null pk autoincr comment('主键id') INT(10)"`
	Value string `json:"value,omitempty" xorm:"not null comment('sku规格值') VARCHAR(20)"`
	Stock int    `json:"stock,omitempty" xorm:"not null default 100 comment('库存') INT(10)"`
	Price string `json:"price,omitempty" xorm:"not null default 100.00 comment('价格') DECIMAL(10,2)"`
	Kid   int    `json:"kid,omitempty" xorm:"not null comment('sku规格id') INT(10)"`
}

type SkuVal struct {
	GoodsSkuVal `xorm:"extends"`
	Num         int `json:"num,omitempty"`
}
