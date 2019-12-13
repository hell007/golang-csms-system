package inits

import (
	"github.com/kataras/golog"

	"go-mvc/framework/models/system"
	"go-mvc/csms/inits/sys"
)

func init() {
	initRootUser()
}

func initRootUser() {
	// root is existed?
	if sys.CheckRootExit() {
		return
	}

	// create root user
	sys.CreateRoot()
	ok := sys.CreateSystemRole()
	if ok {
		addRoleMenu()
	}
}

func addRoleMenu() {
	// 添加role-menu关系
	rMenus := []*system.RoleMenu{
		{Rid: 2, Mid: 1},
		{Rid: 2, Mid: 2},
		{Rid: 2, Mid: 3},
		{Rid: 2, Mid: 4},
		{Rid: 2, Mid: 5},
		{Rid: 2, Mid: 10},
		{Rid: 2, Mid: 11},
		{Rid: 2, Mid: 12},
		{Rid: 2, Mid: 13},
		{Rid: 2, Mid: 19},
		{Rid: 2, Mid: 20},
		{Rid: 2, Mid: 21},
		{Rid: 2, Mid: 22},
	}

	effect, err := system.CreateRoleMenu(rMenus...)
	if err != nil {
		golog.Errorf("===> %d, %s", effect, err.Error())
	}
	golog.Infof("===> %s, %s", effect, err.Error())
}
