package sys

import (
	"github.com/kataras/golog"

	"go-mvc/framework/middleware/casbin"
)

var (
	// 定义系统初始的角色
	Components = [][]string{
		{"admin", "/sys/*", "GET|POST|DELETE|PUT", ".*"},
		{"member", "/member/*", "GET|POST|DELETE|PUT", ".*"},
		{"goods", "/goods/*", "GET|POST|DELETE|PUT", ".*"},
		{"order", "/order/*", "GET|POST|DELETE|PUT", ".*"},
	}
)

// 创建系统默认角色
func CreateSystemRole() bool {
	e := casbin.GetEnforcer()

	for _, v := range Components {
		p := e.GetFilteredPolicy(0, v[0])
		if len(p) == 0 {
			if ok, _ := e.AddPolicy(v); !ok {
				golog.Errorf("初始化角色[%s]权限失败。[%s]", v)
			}
		}
	}
	return true
}
