package member

type Address struct {
	Id        int    `json:"id" xorm:"not null pk autoincr comment('地址id') unique INT(10)"`
	Consignee string `json:"consignee,omitempty" xorm:"not null comment('收件人') VARCHAR(20)"`
	Phone     string `json:"phone,omitempty" xorm:"not null comment('手机号码') VARCHAR(15)"`
	Tell      string `json:"tell,omitempty" xorm:"comment('座机号码') VARCHAR(15)"`
	Zipcode   string `json:"zipcode,omitempty" xorm:"comment('邮政编码') VARCHAR(10)"`
	Province  string `json:"province,omitempty" xorm:"not null comment('省份直辖市') VARCHAR(20)"`
	City      string `json:"city,omitempty" xorm:"not null comment('市') VARCHAR(20)"`
	District  string `json:"district,omitempty" xorm:"not null comment('区县') VARCHAR(20)"`
	Address   string `json:"address,omitempty" xorm:"not null comment('详细地址') VARCHAR(100)"`
	State     int    `json:"state,omitempty" xorm:"not null default 1 comment('状态(1：默认，2：非)') TINYINT(3)"`
	Mid       int    `json:"mid,omitempty" xorm:"not null comment('会员id') INT(10)"`
}

type AddressDetail struct {
	Address *Address `json:"address"`
	Token   string   `json:"token" `
}
