package system

type Menu struct {
	Id       int    `json:"id" xorm:"not null pk autoincr comment('权限id') INT(10)"`
	Name     string `json:"name" xorm:"not null comment('权限名称') VARCHAR(32)"`
	Pid      int    `json:"pid" xorm:"not null comment('权限父id') INT(10)"`
	Path     string `json:"path" xorm:"comment('权限路径') VARCHAR(100)"`
	Redirect string `json:"redirect" xorm:"comment('url') VARCHAR(100)"`
	Level    int    `json:"level" xorm:"not null comment('权限级别') TINYINT(3)"`
	Code     string `json:"code" xorm:"comment('权限标识') VARCHAR(100)"`
	Icon     string `json:"icon" xorm:"comment('图标') VARCHAR(100)"`
	Status   int    `json:"status" xorm:"not null default 1 comment('状态：启用=1/停用=2') TINYINT(1)"`
	Hidden   int    `json:"hidden" xorm:"not null default 2 comment('显示： 显示=1/隐藏=2') TINYINT(1)"`
}
