package goods

type Category struct {
	Id          int       `json:"id" xorm:"not null pk autoincr comment('分类id') unique INT(10)"`
	CategoryName   string `json:"categoryName,omitempty" xorm:"not null comment('分类名称') unique VARCHAR(20)"`
	Pid         int       `json:"pid" xorm:"not null comment('分类父级id') unique INT(10)"`
	IsNav       int       `json:"isNav,omitempty" xorm:"not null default 2 comment('导航(1 => '是', 2 => '否')') TINYINT(3)"`
	Status      int       `json:"status,omitempty" xorm:"not null default 1 comment('状态：启用=1/停用=2') TINYINT(3)"`
}

type CategoryDetail struct {
	Category `xorm:"extends"`
	Small       string    `json:"small,omitempty"`
}
