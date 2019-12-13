package dao

import (
	"fmt"
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/zone"
	"go-mvc/framework/utils/page"
)

type ZoneDao struct {
	engine *xorm.Engine
}

func NewZoneDao(engine *xorm.Engine) *ZoneDao {
	return &ZoneDao{
		engine: engine,
	}
}

// GetAll
func (d *ZoneDao) GetAll() []models.Zone {
	datalist := make([]models.Zone, 0)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

// List
func (d *ZoneDao) List(p *page.Pagination) ([]models.Zone, int64, error) {

	list := make([]models.Zone, 0)

	s := d.engine.Limit(p.Limit, p.Start)

	count, err := s.FindAndCount(&list)

	return list, count, err
}

// GetCity
func (d *ZoneDao) GetCity(areaType interface{}) ([]models.City, error) {
	var (
		err error
		sql string
		list = make([]models.City, 0)
	)

	if areaType != nil {
		sql = fmt.Sprintf(`SELECT Z.id, Z.area_id AS value, Z.area_name AS text, Z.pid FROM jie_zone Z 
			WHERE Z.area_type = %d GROUP BY id asc`, areaType)
		err = d.engine.SQL(sql).Find(&list)
	} else {
		sql = `SELECT Z.id, Z.area_id AS value, Z.area_name AS text, Z.pid FROM jie_zone Z GROUP BY id asc`
		err = d.engine.SQL(sql).Find(&list)
	}

	return list, err
}

// Get
func (d *ZoneDao) Get(id int) (*models.Zone, error) {
	data := &models.Zone{Id: id}
	_, err := d.engine.Get(data)
	return data, err
}

// update
func (d *ZoneDao) Update(zone *models.Zone, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(zone.Id).MustCols(columns...).Update(zone)
	} else {
		effect, err = d.engine.Id(zone.Id).AllCols().Update(zone)
	}
	return effect, err
}

// insert
func (d *ZoneDao) Create(zone *models.Zone) (int64, error) {
	effect, err := d.engine.Insert(zone)
	return effect, err
}

// delete
func (d *ZoneDao) Delete(id int) (int64, error) {
	var (
		effect int64
		err    error
	)
	zone := new(models.Zone)
	effect, err = d.engine.Where("id = ?", id).Delete(zone)
	return effect, err
}