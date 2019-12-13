package services


import (
	"go-mvc/framework/dao"
	models "go-mvc/framework/models/member"
	"go-mvc/framework/utils/db"
	"go-mvc/framework/utils/page"
)

type MemberService interface {
	Exist(member *models.Member) (bool, error)
	List(name string, status int, p *page.Pagination) ([]models.Member, int64, error)
	Get(id int) (*models.MemberDetail, error)
	GetMemberByMobile(mobile string) (*models.Member, bool, error)
	Update(member *models.Member, columns []string) (int64, error)
	Create(member *models.Member) (int64, error)
	Delete(ids []int) (int64, error)
	Close(ids []int) (int64, error)
}

type memberService struct {
	dao *dao.MemberDao
}

func NewMemberService() MemberService {
	return &memberService{
		dao: dao.NewMemberDao(db.MasterEngine()),
	}
}

func (s *memberService) Exist(member *models.Member) (bool, error) {
	return s.dao.Exist(member)
}

func (s *memberService) List(name string, status int, p *page.Pagination) ([]models.Member, int64, error) {
	return s.dao.List(name, status, p)
}

func (s *memberService) Get(id int) (*models.MemberDetail, error) {
	return s.dao.Get(id)
}

func (s *memberService) GetMemberByMobile(mobile string) (*models.Member, bool, error) {
	return s.dao.GetMemberByMobile(mobile)
}

func (s *memberService) Update(member *models.Member, columns []string) (int64, error) {
	return s.dao.Update(member, columns)
}

func (s *memberService) Create(member *models.Member) (int64, error) {
	return s.dao.Create(member)
}

func (s *memberService) Delete(ids []int) (int64, error) {
	return s.dao.Delete(ids)
}

func (s *memberService) Close(ids []int) (int64, error) {
	return s.dao.Close(ids)
}