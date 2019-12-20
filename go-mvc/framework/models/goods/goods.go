package goods

import (
	morder "go-mvc/framework/models/order"
	"go-mvc/framework/utils/times"
)

type Goods struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('商品id') unique INT(10)"`
	GoodsSn     string    `json:"goodsSn,omitempty" xorm:"not null comment('商品编号') VARCHAR(20)"`
	GoodsName   string    `json:"goodsName,omitempty" xorm:"not null comment('商品名称') unique VARCHAR(50)"`
	PromoteWord string    `json:"promoteWord,omitempty" xorm:"not null comment('促销词') VARCHAR(100)"`
	Keywords    string    `json:"keywords,omitempty" xorm:"not null comment('关键词') VARCHAR(255)"`
	Description string    `json:"description,omitempty" xorm:"comment('描述') VARCHAR(255)"`
	Contents    string    `json:"contents,omitempty" xorm:"comment('商品内容') TEXT"`
	CategoryId  int       `json:"categoryId,omitempty" xorm:"not null comment('分类') INT(10)"`
	GoodsPrice  string    `json:"goodsPrice,omitempty" xorm:"not null default 999999.00 comment('商品价') DECIMAL(10,2)"`
	IsOnSale    int       `json:"isOnSale,omitempty" xorm:"not null default 2 comment('是/否 上架(1 => '是', 2 => '否')') TINYINT(3)"`
	IsFirst     int       `json:"isFirst,omitempty" xorm:"not null default 2 comment('是/否  主推(1 => '是', 2 => '否')') TINYINT(3)"`
	IsHot       int       `json:"isHot,omitempty" xorm:"not null default 2 comment('是/否  热卖(1 => '是', 2 => '否')') TINYINT(3)"`
	Unit        string    `json:"unit,omitempty" xorm:"not null default '20枝/扎' comment('单位') VARCHAR(20)"`
	Color       string    `json:"color,omitempty" xorm:"not null comment('颜色') VARCHAR(20)"`
	Views       int       `json:"views,omitempty" xorm:"not null default 188 comment('浏览数') INT(10)"`
	Concerns    int       `json:"concerns,omitempty" xorm:"not null default 100 comment('关注数') INT(10)"`
	CreateTime  *times.JsonTime `json:"createTime,omitempty" xorm:"not null comment('创建时间') DATETIME"`
	UpdateTime  *times.JsonTime `json:"updateTime,omitempty" xorm:"comment('更新时间') DATETIME"`
}


type GoodsDetail struct {
	Goods                 `xorm:"extends"`
	CategoryName string   `json:"categoryName,omitempty" xorm:"not null comment('分类名称') unique VARCHAR(20)"`
	Small     string      `json:"small,omitempty" xorm:"not null comment('小图') VARCHAR(100)"`
	Medium    string      `json:"medium,omitempty" xorm:"not null comment('中图') VARCHAR(100)"`
	Source    string      `json:"source,omitempty" xorm:"not null comment('中图') VARCHAR(100)"`
}

type Product struct {
	Goods   *Goods            `json:"goods"`
	Gallery []GoodsGallery    `json:"gallery"`
	Sku     []GoodsSku        `json:"sku"`
}

type ProductForm struct {
	Code   int                 `json:"code"`
	Token string			   `json:"token"`
	Order  *morder.Order       `json:"order"`
	Goods  []morder.OrderGoods `json:"goods"`
}