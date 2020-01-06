package zone

type Zone struct {
	Id       int    `json:"id,omitempty" xorm:"not null pk autoincr MEDIUMINT(6)"`
	AreaId   string `json:"areaId,omitempty" xorm:"not null default '0' comment('地区标识') VARCHAR(20)"`
	AreaName string `json:"areaName,omitempty" xorm:"not null comment('地区名称') VARCHAR(100)"`
	AreaType int    `json:"areaType,omitempty" xorm:"not null default 1 comment('地区等级') index TINYINT(1)"`
	Pid      string `json:"pid,omitempty" xorm:"not null default '0' comment('上级地区') VARCHAR(20)"`
	Lng      string `json:"lng,omitempty" xorm:"comment('经度') VARCHAR(20)"`
	Lat      string `json:"lat,omitempty" xorm:"comment('纬度') VARCHAR(20)"`
}

type City struct {
	Id    int    `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
	Pid   string `json:"pid,omitempty"`
}
