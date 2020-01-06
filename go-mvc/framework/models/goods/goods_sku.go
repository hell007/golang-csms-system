package goods

// sku
type GoodsSku struct {
	GoodsSkuKey `xorm:"extends"`
	List        []GoodsSkuVal `json:"list"`
}
