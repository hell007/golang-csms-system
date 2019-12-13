package system

import (
	"go-mvc/framework/utils/db"
)

// RoleMenu  角色-菜单关联表
type RoleMenu struct {
	Id  string `json:"id" xorm:"not null  comment('uuid') VARCHAR(32)"`
	Rid int64  `json:"rid" xorm:"not null comment('角色id') INT(10)"`
	Mid int64  `json:"mid" xorm:"not null comment('菜单id') INT(10)"`
}

type RoleMenus struct {
	Rid  int64      `json:"rid" `
	List []RoleMenu `json:"list" `
}

// CreateRoleMenu 建立角色菜单
func CreateRoleMenu(roleMenu ...*RoleMenu) (int64, error) {
	e := db.MasterEngine()
	return e.Insert(roleMenu)
}
