package models

import (
	"fmt"
	"html"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

// ProductImage 产品图片
type ProductImage struct {
	Id        string    `xorm:"prod_img_id pk"`
	ProductId string    `xorm:"prod_img_prod_id notnull"`  // 产品ID
	Path      string    `xorm:"prod_img_path notnull"`     // 图片存放相对路径
	PlaceAt   uint8     `xorm:"prod_img_place_at notnull"` // 图片放置的网页位置。1-封面；2-详情；
	Created   time.Time `xorm:"prod_img_created created"`
}

func (this *ProductImage) TableName() string {
	return "t_product_img"
}

func (this *ProductImage) HtmlEscape() *ProductImage {
	this.Path = html.EscapeString(this.Path)
	return this
}

// AddProductImage 添加产品图片记录。此时并未将产品图片与产品进行关联。
func AddProductImage(img *ProductImage) error {
	img.Id = uuid.NewV4().String()
	_, err := x.Insert(img)
	if err != nil {
		return err
	}
	return nil
}

func GetProductImages(cond *ProductImage) (images []*ProductImage, err error) {
	if cond == nil {
		cond = &ProductImage{}
	}
	// images := make([]*ProductImage)
	err = x.Find(&images, cond)
	return images, err
}

// Product 产品
type Product struct {
	Id              string          `xorm:"prod_id pk"`
	Title           string          `xorm:"prod_title notnull"`
	Intro           string          `xorm:"prod_intro notnull"`             // 简介
	Desc            string          `xorm:"prod_desc notnull"`              // 描述
	DescUseMarkdown uint8           `xorm:"prod_desc_use_markdown notnull"` //描述内容是否使用了markdown来书写。1-是；2-否；
	IsPublic        uint8           `xorm:"prod_is_public notnull"`         // 是否公开。
	SortNo          uint            `xorm:"prod_sort_no notnull"`           // 排序序号，序号越大越靠前。
	Created         time.Time       `xorm:"prod_created created"`
	Modified        time.Time       `xorm:"prod_modified updated"`
	CoverImg        *ProductImage   `xorm:"-"`
	DetailImgs      []*ProductImage `xorm:"-"`
}

func (this *Product) TableName() string {
	return "t_product"
}

func (this *Product) HtmlEscape() *Product {
	this.Title = html.EscapeString(this.Title)
	this.Intro = html.EscapeString(this.Intro)
	this.Desc = html.EscapeString(this.Desc)
	this.CoverImg = this.CoverImg.HtmlEscape()
	for i, _ := range this.DetailImgs {
		this.DetailImgs[i] = this.DetailImgs[i].HtmlEscape()
	}
	return this
}

func GetProductById(id string, lazy bool) (prod *Product, has bool, err error) {
	prod = new(Product)
	has, err = x.Id(id).Get(prod)
	if err != nil {
		beego.Error(err)
		return nil, has, err
	}
	if has && !lazy {
		// 加载封面信息
		coverImgs, err := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_COVER, ProductId: prod.Id})
		if err != nil {
			beego.Error(err)
			return nil, has, err
		}
		if len(coverImgs) > 0 {
			prod.CoverImg = coverImgs[0]
		}
		//加载详情图片信息
		detailImgs, err := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_DETAIL, ProductId: prod.Id})
		prod.DetailImgs = detailImgs
		if err != nil {
			beego.Error(err)
			return nil, has, err
		}
	}

	return prod, has, nil
}

func PushPinProduct(id string) (affected int64, err error) {
	sql := "UPDATE t_product AS t1 SET t1.prod_sort_no = (SELECT MAX(prod_sort_no)+1 FROM (SELECT * FROM t_product)AS t0) WHERE t1.prod_id = ?"
	result, err := x.Exec(sql, id)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}

// 添加产品及其图片信息
func AddProduct(prod *Product) error {
	prod.Id = uuid.NewV4().String()

	prod.CoverImg.Id = uuid.NewV4().String()
	prod.CoverImg.ProductId = prod.Id
	prod.CoverImg.PlaceAt = IMG_PLACE_AT_COVER

	for i, _ := range prod.DetailImgs {
		prod.DetailImgs[i].Id = uuid.NewV4().String()
		prod.DetailImgs[i].ProductId = prod.Id
		prod.DetailImgs[i].PlaceAt = IMG_PLACE_AT_DETAIL
	}

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}
	// 插入封面及详情图片
	// imgs := append(prod.DetailImgs, prod.CoverImg)
	// _, err := x.Insert(&imgs)
	_, err := sess.Insert(prod.CoverImg, prod.DetailImgs)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return err
	}

	// 插入产品信息
	_, err = sess.Insert(prod)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return err
	}

	return sess.Commit()
}

func RmProducts(prodIds ...string) (affected int64, err error) {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return 0, err
	}

	prod := &Product{}
	affected, err = sess.In("prod_id", prodIds).Delete(prod)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return affected, err
	}

	prodImg := &ProductImage{}
	_, err = sess.In("prod_img_prod_id", prodIds).Delete(prodImg)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return affected, err
	}

	return affected, sess.Commit()
}

func ModProduct(prod *Product) (affected int64, err error) {
	sess := x.NewSession()
	defer sess.Close()

	if err = sess.Begin(); err != nil {
		beego.Error(err)
		return affected, err
	}

	affected, err = sess.Update(prod, &Product{Id: prod.Id})
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return affected, err
	}

	if prod.CoverImg != nil && len(prod.DetailImgs) != 0 {
		// 删除产品下所有图片
		_, err = sess.Delete(&ProductImage{ProductId: prod.Id})
		if err != nil {
			beego.Error(err)
			sess.Rollback()
			return affected, err
		}
		//再重新添加图片
		prod.CoverImg.Id = uuid.NewV4().String()
		prod.CoverImg.ProductId = prod.Id
		prod.CoverImg.PlaceAt = IMG_PLACE_AT_COVER

		for i, _ := range prod.DetailImgs {
			prod.DetailImgs[i].Id = uuid.NewV4().String()
			prod.DetailImgs[i].ProductId = prod.Id
			prod.DetailImgs[i].PlaceAt = IMG_PLACE_AT_DETAIL
		}
		_, err = sess.Insert(prod.CoverImg, prod.DetailImgs)
		if err != nil {
			beego.Error(err)
			sess.Rollback()
			return affected, err
		}
	}

	return affected, sess.Commit()
}

func CountProducts(cond *Product) (int64, error) {
	if cond == nil {
		cond = &Product{}
	}
	return x.Count(cond)
}

func GetProducts(cond *Product, limit, offset int, lazy bool) ([]*Product, error) {
	if cond == nil {
		cond = &Product{}
	}

	prods := make([]*Product, 0, limit)
	err := x.Desc("prod_sort_no", "prod_created").Limit(limit, offset).Find(&prods, cond)
	if err != nil {
		return prods, err
	}

	if !lazy {
		for i, _ := range prods {
			// 加载封面信息
			coverImgs, err := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_COVER, ProductId: prods[i].Id})
			if err != nil {
				beego.Error(err)
				continue
			}
			if len(coverImgs) > 0 {
				prods[i].CoverImg = coverImgs[0]
			}
			//加载详情图片信息
			detailImgs, _ := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_DETAIL, ProductId: prods[i].Id})
			prods[i].DetailImgs = detailImgs
		}
	}
	return prods, err
}

func GetProductsButFlagshipsPage(cond *Product, curPageNo, pageSize int, lazy bool) (page *Page, err error) {
	if cond == nil {
		cond = &Product{}
	}

	totalRecords, err := CountProductsButFlagships(cond)
	if err != nil {
		beego.Error(fmt.Sprintf("CountProductsButFlagships(%#v) err:%s", cond, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(cond), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	lim := pager.BuildLimiter()
	rows, err := GetProductsButFlagships(cond, lim.Limit, lim.Offset, lazy)
	if err != nil {
		beego.Error(fmt.Sprintf("GetProductsButFlagships(%#v,%d,%d,%t) err:%s", cond, lim.Limit, lim.Offset, lazy, err))
		return nil, err
	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}

func CountProductsButFlagships(cond *Product) (int64, error) {
	if cond == nil {
		cond = &Product{}
	}
	return x.Where("prod_id NOT IN(SELECT fp_prod_id FROM t_home_flagship_product)").Count(cond)
}

func GetProductsButFlagships(cond *Product, limit, offset int, lazy bool) ([]*Product, error) {
	if cond == nil {
		cond = &Product{}
	}
	prods := make([]*Product, 0, limit)
	err := x.Where("prod_id NOT IN(SELECT fp_prod_id FROM t_home_flagship_product)").Limit(limit, offset).Find(&prods, cond)
	if err != nil {
		return prods, err
	}
	if !lazy {
		for i, _ := range prods {
			// 加载封面信息
			coverImgs, err := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_COVER, ProductId: prods[i].Id})
			if err != nil {
				beego.Error(err)
				continue
			}
			if len(coverImgs) > 0 {
				prods[i].CoverImg = coverImgs[0]
			}
			//加载详情图片信息
			detailImgs, _ := GetProductImages(&ProductImage{PlaceAt: IMG_PLACE_AT_DETAIL, ProductId: prods[i].Id})
			prods[i].DetailImgs = detailImgs
		}
	}
	return prods, err
}

func GetProductPage(cond *Product, curPageNo, pageSize int, lazy bool) (page *Page, err error) {
	if cond == nil {
		cond = &Product{}
	}

	totalRecords, err := CountProducts(cond)
	if err != nil {
		beego.Error(fmt.Sprintf("CountProducts(%#v) err:%s", cond, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(cond), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	lim := pager.BuildLimiter()
	rows, err := GetProducts(cond, lim.Limit, lim.Offset, lazy)
	if err != nil {
		beego.Error(fmt.Sprintf("GetProducts(%#v,%d,%d,%t) err:%s", cond, lim.Limit, lim.Offset, lazy, err))
		return nil, err

	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}
