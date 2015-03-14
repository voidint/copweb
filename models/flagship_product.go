package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

type FlagshipProduct struct {
	Id        string    `xorm:"fp_id pk"`
	ProductId string    `xorm:"fp_prod_id"`
	SortNo    uint      `xorm:"fp_sort_no notnull"`
	Created   time.Time `xorm:"fp_created created"`
	Modified  time.Time `xorm:"fp_modified updated"`
	Product   *Product  `xorm:"-"`
}

func (this *FlagshipProduct) TableName() string {
	return "t_home_flagship_product"
}

func GetAllFlagshipProducts(lazy bool) (fProds []*FlagshipProduct, err error) {
	err = x.Desc("fp_sort_no", "fp_created").Find(&fProds, &FlagshipProduct{})
	if err != nil {
		return nil, err
	}
	if !lazy {
		for i := range fProds {
			prod, _, err := GetProductById(fProds[i].ProductId, false)
			if err != nil {
				return nil, err
			}
			fProds[i].Product = prod
		}
	}

	return fProds, nil
}

func CountFlagshipProducts(cond *FlagshipProduct) (affected int64, err error) {
	if cond == nil {
		cond = &FlagshipProduct{}
	}
	return x.Count(cond)
}

func GetFlagshipProducts(cond *FlagshipProduct, limit, offset int, lazy bool) (fProds []*FlagshipProduct, err error) {
	if cond == nil {
		cond = &FlagshipProduct{}
	}
	err = x.Desc("fp_sort_no", "fp_created").Limit(limit, offset).Find(&fProds, cond)
	if err != nil {
		return nil, err
	}
	if !lazy {
		for i := range fProds {
			prod, _, err := GetProductById(fProds[i].ProductId, false)
			if err != nil {
				return nil, err
			}
			fProds[i].Product = prod
		}
	}
	return fProds, nil
}

func AddFlagshipProducts(prods ...*FlagshipProduct) (affected int64, err error) {
	for idx := range prods {
		prods[idx].Id = uuid.NewV4().String()
	}

	return x.Insert(prods)
}

func RmFlagshipProductById(id string) (affected int64, err error) {
	cond := &FlagshipProduct{Id: id}
	return x.Delete(cond)
}

func PushPinFlagshipProduct(id string) (affected int64, err error) {
	sql := "UPDATE t_home_flagship_product AS t1 SET t1.fp_sort_no = (SELECT MAX(fp_sort_no)+1 FROM (SELECT * FROM t_home_flagship_product)AS t0) WHERE t1.fp_id = ?"
	result, err := x.Exec(sql, id)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}
