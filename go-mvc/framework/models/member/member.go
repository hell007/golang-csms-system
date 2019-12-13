package member

import "time"

type Member struct {
	Id         int       `json:"id,omitempty" xorm:"not null pk autoincr comment('会员id') INT(10)"`
	Name       string    `json:"name,omitempty" xorm:"not null comment('真实姓名') VARCHAR(20)"`
	Password   string    `json:"password,omitempty" xorm:"not null comment('密码') CHAR(32)"`
	Salt       string    `json:"salt,omitempty" xorm:" comment('盐值') VARCHAR(64)"`
	Mobile     string    `json:"mobile,omitempty" xorm:"not null comment('手机号码') unique VARCHAR(15)"`
	Gender     string    `json:"gender,omitempty" xorm:"not null default '保密' comment('性别') ENUM('保密','女','男')"`
	Ip         string    `json:"ip,omitempty" xorm:"not null comment('登录ip') VARCHAR(20)"`
	CreateTime time.Time `json:"createTime,omitempty" xorm:"not null comment('创建时间') DATETIME"`
	LoginTime  time.Time `json:"loginTime,omitempty" xorm:" comment('登录时间') DATETIME"`
	Status     int       `json:"status,omitempty" xorm:"not null default 2 comment('状态(1：已认证，2：未认证，3：已注销)') TINYINT(3)"`
	Note       string    `json:"note,omitempty" xorm:"comment('记录') VARCHAR(200)"`
}

type LoginUser struct {
	Member `xorm:"extends"`
	Code   string `json:"code,omitempty"`
	Token  string `json:"token,omitempty"`
}

type MemberDetail struct {
	Member *Member  `json:"member"`
	List  []Address `json:"list"`
}
