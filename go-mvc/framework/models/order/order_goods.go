package order

type OrderGoods struct {
	Id         string `json:"id" xorm:"not null pk comment('id') VARCHAR(64)"`
	GoodsId    int    `json:"goodsId,omitempty" xorm:"not null comment('商品id') INT(10)"`
	GoodsSku   string `json:"goodsSku,omitempty" xorm:"not null comment('商品规格') VARCHAR(145)"`
	GoodsPrice string `json:"goodsPrice,omitempty" xorm:"not null DECIMAL(10,2)"`
	GoodsNum   int    `json:"goodsNum,omitempty" xorm:"not null INT(10)"`
	OrderId    int    `json:"orderId,omitempty" xorm:"not null comment('订单id') INT(10)"`
}

type OrderGoodsDetail struct {
	OrderGoods `xorm:"extends"`
	GoodsSn    string `json:"goodsSn,omitempty"`
	GoodsName  string `json:"goodsName,omitempty"`
	Color      string `json:"color,omitempty"`
	Unit       string `json:"unit,omitempty"`
	Small      string `json:"small,omitempty"`
}
