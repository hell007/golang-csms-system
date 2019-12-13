package dao

import (
	"fmt"
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/member"
	"go-mvc/framework/utils/page"
)

type MemberDao struct {
	engine *xorm.Engine
}

func NewMemberDao(engine *xorm.Engine) *MemberDao {
	return &MemberDao{
		engine: engine,
	}
}

// exist
func (d *MemberDao) Exist(member *models.Member) (bool, error) {
	has, err := d.engine.Exist(member)
	return has, err
}

// List
func (d *MemberDao) List(name string, status int, p *page.Pagination) ([]models.Member, int64, error) {
	list := make([]models.Member, 0)

	s := d.engine.Omit("password, salt")

	if name != "" {
		s.Where("name like ?", "%"+name+"%")
	}

	if status > 0 {
		s.And("status = ?", status)
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

// Get
func (d *MemberDao) Get(id int) (*models.MemberDetail, error) {
	var (
		err  error
		ok   bool
		member = new(models.Member)
		memberDetail  = new(models.MemberDetail)
		list = make([]models.Address, 0)
	)
	member.Id = id
	ok, err = d.engine.Omit("password, salt").Get(member)

	if ok && err == nil {
		memberDetail.Member = member

		sql := fmt.Sprintf(`SELECT  A.id, A.consignee, A.phone, A.tell, A.zipcode, A.state, 
		(SELECT area_name from jie_zone WHERE province = area_id) as province,
		(SELECT area_name from jie_zone WHERE city = area_id) as city,
		(SELECT area_name from jie_zone WHERE district = area_id) as district,
		A.address FROM jie_address A WHERE A.mid = %d`, member.Id)
		err = d.engine.SQL(sql).Find(&list)
		memberDetail.List = list
	}
	return memberDetail, err
}

// getMemberByMobile
func (d *MemberDao) GetMemberByMobile(mobile string) (*models.Member, bool, error) {
	var (
		member = new(models.Member)
	)
	member.Mobile = mobile
	has, err := d.engine.Get(member)
	return member, has, err
}

// update
func (d *MemberDao) Update(member *models.Member, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(member.Id).Cols(columns...).Update(member)
	} else {
		effect, err = d.engine.Id(member.Id).AllCols().Update(member)
	}
	return effect, err
}

// insert
func (d *MemberDao) Create(member *models.Member) (int64, error) {
	effect, err := d.engine.Insert(member)
	return effect, err
}

// delete
func (d *MemberDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	u := new(models.Member)
	effect, err = d.engine.In("id", ids).Delete(u)
	return effect, err
}

// close
func (d *MemberDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	u := new(models.Member)
	u.Status = 2
	effect, err = d.engine.In("id", ids).Cols("status").Update(u)
	return effect, err
}