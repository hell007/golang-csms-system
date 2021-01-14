package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"net/http"
	"sync"

	"github.com/casbin/casbin/v2/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12/context"

	"go-mvc/framework/conf"
	"go-mvc/framework/db"
	"go-mvc/framework/middleware/jwt"
	"go-mvc/framework/utils/response"
)

var (
	adt *Adapter // Your driver and data source.
	e   *casbin.Enforcer

	adtLook sync.Mutex
	eLook   sync.Mutex

	rbacModel string
)

func SetRbacModel(rootID string) {
	rbacModel = fmt.Sprintf(`
[request_definition]
r = sub, obj, act, suf

[policy_definition]
p = sub, obj, act, suf

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch(r.obj, p.obj) && regexMatch(r.suf, p.suf) && regexMatch(r.act, p.act) || r.sub == "%s"
`, rootID)
}

// 获取Enforcer
func GetEnforcer() *casbin.Enforcer {
	if e != nil {
		e.LoadPolicy()
		return e
	}
	eLook.Lock()
	defer eLook.Unlock()
	if e != nil {
		e.LoadPolicy()
		return e
	}

	m, _ := model.NewModelFromString(rbacModel)
	e, _ = casbin.NewEnforcer(m, singleAdapter())

	e.EnableLog(true)
	return e
}

func singleAdapter() *Adapter {
	if adt != nil {
		return adt
	}
	adtLook.Lock()
	defer adtLook.Unlock()
	if adt != nil {
		return adt
	}

	url := db.GetConnURL()
	adt = NewAdapter(conf.GlobalConfig.MysqlDialect, url, true)
	return adt
}

func CheckPermissions(ctx context.Context) bool {
	ut, ok := jwt.ParseToken(ctx)

	if !ok {
		return false
	}

	//uid := strconv.Itoa(int(ut.Id))
	rolename := ut.Rolename
	yes, _ := GetEnforcer().Enforce(rolename, ctx.Path(), ctx.Method(), ".*")

	if !yes {
		response.Unauthorized(ctx, response.PermissionsLess, nil)
		ctx.StopExecution()
		return false
	}

	return true
}

func Wrapper() func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		//if !c.Check(r) {
		//	w.WriteHeader(http.StatusForbidden)
		//	w.Write([]byte("403 Forbidden"))
		//	return
		//}
		router(w, r)
	}
}
