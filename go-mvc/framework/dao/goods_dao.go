package dao

import (
	"github.com/go-xorm/xorm"

	models "go-mvc/framework/models/goods"
	"go-mvc/framework/utils/page"
	"fmt"
)

type GoodsDao struct {
	engine *xorm.Engine
}

func NewGoodsDao(engine *xorm.Engine) *GoodsDao {
	return &GoodsDao{
		engine: engine,
	}
}

// List
func (d *GoodsDao) List(name string, category int, isOnSale int, isFirst int, isHot int, p *page.Pagination) ([]models.GoodsDetail, int64, error) {
	list := make([]models.GoodsDetail, 0)

	// GA去重
	s := d.engine.Table("jie_goods").Alias("G").Select("G.*, C.category_name, GA.small").
		Join("LEFT", "jie_category C", "G.category_id = C.id").
		Join("LEFT", "(SELECT goods_id, small from jie_goods_gallery group by goods_id) as GA", "G.id = GA.goods_id")

	if name != "" {
		s.Where("G.goods_name like ?", "%"+name+"%").Or("G.goods_sn = ?", name)
	}

	if category > 0 {
		s.And("G.category_id = ?", category)
	}

	if isOnSale > 0 {
		s.And("G.is_on_sale = ?", isOnSale)
	}

	if isFirst > 0 {
		s.And("G.is_first = ?", isFirst)
	}

	if isHot > 0 {
		s.And("G.is_hot = ?", isHot)
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

	count, err :=s.FindAndCount(&list)

	return list, count, err
}

// Get
func (d *GoodsDao) Get(id int) (*models.Goods, []models.GoodsGallery, []models.GoodsSkuKey, []models.GoodsSkuVal, error) {
	var (
		err  error
		ok   bool
		goods = new(models.Goods)
		gallerys = make([]models.GoodsGallery, 0)
		skuKey = make([]models.GoodsSkuKey, 0)
		skuVal = make([]models.GoodsSkuVal, 0)
	)

	goods.Id = id
	ok, err = d.engine.Get(goods)

	if ok && err == nil {
		gallerySql := fmt.Sprintf(`SELECT GA.*
		FROM jie_goods_gallery GA WHERE GA.goods_id = %d`, goods.Id)
		err = d.engine.SQL(gallerySql).OrderBy("GA.gallery_id").Find(&gallerys)

		skuSql := fmt.Sprintf(`SELECT K.* FROM jie_goods_sku_key K WHERE K.goods_sn = "%s"`, goods.GoodsSn)
		err = d.engine.SQL(skuSql).OrderBy("K.kid").Find(&skuKey)

		skuvSql := fmt.Sprintf(`SELECT V.* FROM jie_goods_sku_val V WHERE V.kid in ( 
		SELECT K.kid FROM jie_goods_sku_key K WHERE K.goods_sn = "%s" )`, goods.GoodsSn)
		err = d.engine.SQL(skuvSql).OrderBy("V.vid").Find(&skuVal)

		return  goods, gallerys, skuKey, skuVal, err

	} else {
		return  goods, gallerys, skuKey, skuVal, nil
	}
}

// update
func (d *GoodsDao) Update(goods *models.Goods, columns []string) (int64, error) {
	var (
		err    error
		effect int64
	)

	if columns != nil && len(columns) > 0 {
		effect, err = d.engine.Id(goods.Id).MustCols(columns...).Update(goods)
	} else {
		effect, err = d.engine.Id(goods.Id).AllCols().Update(goods)
	}
	return effect, err
}

// insert
func (d *GoodsDao) Create(goods *models.Goods) (int64, error) {
	effect, err := d.engine.Insert(goods)
	return effect, err
}

// delete
func (d *GoodsDao) Delete(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	goods := new(models.Goods)
	effect, err = d.engine.In("id", ids).Delete(goods)
	return effect, err
}

// close
func (d *GoodsDao) Close(ids []int) (int64, error) {
	var (
		effect int64
		err    error
	)

	goods := new(models.Goods)
	goods.IsOnSale = 2
	effect, err = d.engine.In("id", ids).Cols("is_on_sale").Update(goods)
	return effect, err
}

// 热门商品 主推商品
func (d *GoodsDao) GetGoods(state string, val string, p *page.Pagination) ([]models.GoodsDetail, int64, error){
	var (
		list = make([]models.GoodsDetail, 0)
	)

	//builder.NotIn("a", 1, 2)
	// 去重方法二 非常有效
	s := d.engine.Table("jie_goods").Alias("G").Select("G.id, G.goods_name, G.unit, G.goods_price, GA.small").
		Join("LEFT", "(SELECT goods_id, small from jie_goods_gallery group by goods_id) as GA", "G.id = GA.goods_id")

	switch state {
	case "key":
		s.Where("G.goods_name like ?", "%"+val+"%")
	case "category":
		s.Where("G.category_id = ?", val)
	case "hot":
		s.Where("G.is_hot = ?", val)
	case "first":
		s.Where("G.is_first = ?", val)
	default:
	}

	s.And("is_on_sale = ?", 1)

	// 去重方法一，导致count错误
	//s.GroupBy("G.goods_name")

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

// 产品详情
func (d *GoodsDao) GetProduct(id int) (*models.Product, error) {
	var (
		err  error
		ok   bool
		goods = new(models.Goods)
		product = new(models.Product)
		gallerys = make([]models.GoodsGallery, 0)
		skus = make([]models.GoodsSku, 0)
		vlist = make([]models.GoodsSkuVal, 0)
	)

	goods.Id = id
	ok, err = d.engine.Omit("category_id", "is_first", "is_hot", "is_on_sale", "category_name", "Small", "Medium", "Source").Get(goods)

	if ok && err == nil {

		gallerySql := fmt.Sprintf(`SELECT GA.small, GA.source
		FROM jie_goods_gallery GA WHERE GA.goods_id = %d`, goods.Id)
		err = d.engine.SQL(gallerySql).Find(&gallerys)

		product.Goods = goods
		product.Gallery = gallerys

		skuSql := fmt.Sprintf(`SELECT K.kid ,K.name FROM jie_goods_sku_key K WHERE K.goods_sn = "%s"`, goods.GoodsSn)
		err = d.engine.SQL(skuSql).Find(&skus)

		skuvSql := fmt.Sprintf(`SELECT V.* FROM jie_goods_sku_val V WHERE V.kid in ( 
		SELECT K.kid FROM jie_goods_sku_key K WHERE K.goods_sn = "%s" )`, goods.GoodsSn)
		err = d.engine.SQL(skuvSql).Find(&vlist)

		// 构造需要的数据结构，减少数据库的查询
		for i,_ := range skus {
			for j,_ := range vlist {
				if skus[i].Kid == vlist[j].Kid {
					skus[i].List = append(skus[i].List, vlist[j])
				}
			}
		}

		product.Sku = skus
		return  product, err
	} else {
		return  product, err
	}
}