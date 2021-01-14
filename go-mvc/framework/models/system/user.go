package system

import (
	"time"
)

type User struct {
	Id         int       `json:"id,omitempty" xorm:"not null pk autoincr comment('系统id') unique INT(10)"`
	RoleId     int       `json:"roleId,omitempty" xorm:"not null comment('角色') INT(10)"`
	Username   string    `json:"username,omitempty" xorm:"not null comment('系统用户') unique VARCHAR(20)"`
	Password   string    `json:"password,omitempty" xorm:"not null comment('密码') CHAR(32)"`
	Status     int       `json:"status" xorm:"not null default 1 comment('状态：启用=1/停用=2') TINYINT(1)"`
	Salt       string    `json:"salt,omitempty" xorm:"comment('盐值') VARCHAR(64)"`
	Email      string    `json:"email,omitempty" xorm:"comment('邮箱') unique VARCHAR(50)"`
	Mobile     string    `json:"mobile,omitempty" xorm:"not null comment('手机号码') unique VARCHAR(11)"`
	Ip         string    `json:"ip,omitempty" xorm:"comment('登录ip') VARCHAR(20)"`
	CreateTime time.Time `json:"createTime" xorm:"not null comment('创建时间') DATETIME"`
	LoginTime  time.Time `json:"loginTime" xorm:"comment('登录时间') DATETIME"`
}

type UserDetail struct {
	User     `xorm:"extends"`
	RoleName string `json:"roleName"`
}
