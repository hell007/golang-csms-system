package dao

import (
	"fmt"
	//"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/system"
	"go-mvc/framework/utils/page"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	// tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "jie_")
	// engine.SetTableMapper(tbMapper)
	return &UserDao{
		engine: engine,
	}
}

// exist
func (d *UserDao) Exist(user *models.User) (bool, error) {
	has, err := d.engine.Exist(user)
	return has, err
}

// GetAll
func (d *UserDao) GetAll() []models.User {
	datalist := make([]models.User, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *UserDao) List(name string, status int, p *page.Pagination) ([]models.UserDetail, int64, error) {
	list := make([]models.UserDetail, 0)

	s := d.engine.Table("jie_user").Alias("U").
		Select("U.id, U.username, U.mobile, U.status, U.email, U.ip, U.create_time, U.login_time, R.role_name").
		Join("LEFT", "jie_role R", "U.role_id = R.id")

	if name != "" {
		s.Where("U.username like ?", "%"+name+"%")
	}

	if status > 0 {
		s.And("U.status = ?", status)
	}

	s.Limit(p.Limit, p.Start)

	if p.SortName != "" {
		switch p.SortOrder {
		case "asc":
			s.Asc(p.SortName)
		case "desc":
			s.Desc(p.SortName)
		}
	}

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// GetUsersByRids
func (d *UserDao) GetUsersByRids(rids []int, page *page.Pagination) ([]models.User, int64, error) {
	users := make([]models.User, 0)
	s := d.engine.In("id", rids).Limit(page.Limit, page.Start)
	count, err := s.FindAndCount(&users)
	return users, count, err
}

// Get
func (d *UserDao) Get(id int) (*models.User, bool, error) {
	user := &models.User{Id: id}
	has, err := d.engine.Get(user)
	return user, has, err
}

// GetUserByName
func (d *UserDao) GetUserByName(name string, user *models.UserToken) (bool, error) {
	sql := fmt.Sprintf(`
SELECT USER.*, ROLE.ROLE_NAME AS ROLENAME, ROLE.ROLE_NOTE AS ROLENOTE FROM jie_user USER, jie_role ROLE 
WHERE USER.USERNAME = "%s" AND USER.ROLE_ID = ROLE.ID
`, name)
	//d.engine.Where("username = ?", name).Get(user)
	return d.engine.SQL(sql).Get(user)
}

// GetRoleNameById
func (d *UserDao) GetRoleNameByRId(rid int) (string, error) {
	var rolename string
	_, err := d.engine.Table("jie_role").Where("role_id = ?", rid).Cols("role_name").Get(&rolename)
	return rolename, err
}

// update
func (d *UserDao) Update(user *models.User, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(user.Id).Cols(columns...).Update(user)
	} else {
		effect, err = d.engine.Id(user.Id).AllCols().Update(user)
	}
	return effect, err
}

// insert
func (d *UserDao) Create(user *models.User) (int64, error) {
	effect, err := d.engine.Insert(user)
	return effect, err
}

// delete
func (d *UserDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	u := new(models.User)

	effect, err = d.engine.In("id", ids).Delete(u)
	return effect, err
}

// close
func (d *UserDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	u := new(models.User)
	u.Status = 2

	effect, err = d.engine.In("id", ids).Cols("status").Update(u)
	return effect, err
}
