package order

import (
	models "go-mvc/framework/models/member"
	"time"
)

type Order struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('订单id') INT(10)"`
	Ordersn     string    `json:"ordersn,omitempty" xorm:"not null comment('订单编号') unique VARCHAR(20)"`
	OrderState  int       `json:"orderState,omitempty" xorm:"not null default 0000000001 comment('订单状态(1 => '待支付', 2 => '待发货',3 => '已发货',4 => '已取消',5 => '已完成')') INT(10)"`
	CreateTime  time.Time `json:"createTime,omitempty" xorm:"comment('下单时间') DATETIME"`
	PayTime     time.Time `json:"payTime,omitempty" xorm:"comment('付款时间') DATETIME"`
	ShipTime    time.Time `json:"shipTime,omitempty" xorm:"comment('发货时间') DATETIME"`
	ConfirmTime time.Time `json:"confirmTime,omitempty" xorm:"comment('订单确认时间') DATETIME"`
	TotalPrice  string    `json:"totalPrice,omitempty" xorm:"not null default 0.00 comment('订单总价') DECIMAL(10,2)"`
	ShipPrice   string    `json:"shipPrice,omitempty" xorm:"not null default 0.00 comment('物流价格') DECIMAL(5,2)"`
	OrderPrice  string    `json:"orderPrice,omitempty" xorm:"not null default 0.00 comment('订单总价') DECIMAL(10,2)"`
	PayPrice    string    `json:"payPrice,omitempty" xorm:"not null comment('支付总价') DECIMAL(10,2)"`
	ReturnPrice string    `json:"returnPrice,omitempty" xorm:"not null default 0.00 comment('退款') DECIMAL(10,2)"`
	PayName     string    `json:"payName,omitempty" xorm:"comment('支付方式（1=>'在线支付'，2=>'货到付款'）') VARCHAR(20)"`
	ShipName    string    `json:"shipName,omitempty" xorm:"comment('物流方式（1=>'快递'，2=>'上门取货'）') VARCHAR(45)"`
	ShipNo      string    `json:"shipNo,omitempty" xorm:"comment('物流编号') VARCHAR(20)"`
	ShipNote    string    `json:"shipNote,omitempty" xorm:"comment('发货备注') VARCHAR(200)"`
	Mid         int       `json:"mid,omitempty" xorm:"not null comment('会员id') INT(10)"`
	Consignee   string    `json:"consignee,omitempty" xorm:"not null comment('收件人') VARCHAR(20)"`
	Phone       string    `json:"phone,omitempty" xorm:"not null comment('手机号码') VARCHAR(15)"`
	Zipcode     string    `json:"zipcode,omitempty" xorm:"comment('邮政编码') VARCHAR(10)"`
	Tell        string    `json:"tell,omitempty" xorm:"comment('座机号码') VARCHAR(15)"`
	Province    string    `json:"province,omitempty" xorm:"not null comment('省份直辖市') VARCHAR(20)"`
	City        string    `json:"city,omitempty" xorm:"not null comment('市') VARCHAR(20)"`
	District    string    `json:"district,omitempty" xorm:"not null comment('区县') VARCHAR(20)"`
	Address     string    `json:"address,omitempty" xorm:"not null comment('详细地址') VARCHAR(100)"`
	Remark      string    `json:"remark,omitempty" xorm:"comment('订单备注') VARCHAR(200)"`
}

// 订单详情
type OrderDetail struct {
	Order  *Order             `json:"order"`
	Member *models.Member     `json:"member"`
	List   []OrderGoodsDetail `json:"list"`
}

// 订单列表
type Orders struct {
	Order `xorm:"extends"`
	List  []OrderGoodsDetail `json:"list"`
}
