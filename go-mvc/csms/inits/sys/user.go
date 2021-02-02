package sys

import (
	"strconv"
	"time"

	"github.com/kataras/golog"

	"go-mvc/framework/conf"
	"go-mvc/framework/db"
	"go-mvc/framework/middleware/casbin"
	"go-mvc/framework/models/system"
	"go-mvc/framework/utils/encrypt"
)

const (
	username = "root"
	password = "123456"
)

// 检查超级用户是否存在
func CheckRootExit() bool {
	e := db.MasterEngine()
	// root is existed?
	exit, err := e.Exist(&system.User{Username: username})
	if err != nil {
		golog.Errorf("===>When check Root User is exited? happened error. [%s]", err)
	}
	if exit {
		golog.Info("===>Root User is existed.")

		// 初始化rbac_model
		r := system.User{Username: username}

		if exit, _ := e.Get(&r); exit {
			casbin.SetRbacModel(strconv.Itoa(r.Id))
			CreateSystemRole()
		}
	}
	return exit
}

// CreateRoot 建立root用户
func CreateRoot() {
	newRoot := system.User{
		Username:   username,
		Password:   encrypt.AESEncrypt(password, conf.Global.JWTSalt),
		RoleId:     1,
		Status:     1,
		Email:      "root@sina.com",
		Mobile:     "13888888888",
		CreateTime: time.Now(),
	}

	e := db.MasterEngine()
	if _, err := e.Insert(&newRoot); err != nil {
		golog.Errorf("===>When create Root User happened error. [%s]", err)
	}
	rooId := strconv.Itoa(newRoot.Id)
	casbin.SetRbacModel(rooId)

	addAllpolicy(rooId)
}

func addAllpolicy(rootId string) {
	rootId = "superadmin"
	// add policy for root
	e := casbin.GetEnforcer()
	p, _ := e.AddPolicy(rootId, "/*", "ANY", ".*", "", "", "", "", "", "超级管理员")
	if !p {
		golog.Errorf("初始化用户[%s]权限失败.", username)
	}

	for _, v := range Components {
		e.AddGroupingPolicy(rootId, v[0])
	}

	// admin
	for _, v := range Components {
		e.AddGroupingPolicy("admin", v[0])
	}
}
