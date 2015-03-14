package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

type Carousel struct {
	Id       string    `xorm:"caro_id pk"`
	ImgPath  string    `xorm:"caro_img_path notnull"`
	Caption  string    `xorm:"caro_caption notnull"`
	SortNo   uint      `xorm:"caro_sort_no"`
	Created  time.Time `xorm:"caro_created created"`
	Modified time.Time `xorm:"caro_modified updated"`
}

func (this *Carousel) TableName() string {
	return "t_home_carousel"
}

func GetCarouselPage(cond *Carousel, curPageNo, pageSize int) (page *Page, err error) {
	if cond == nil {
		cond = &Carousel{}
	}

	totalRecords, err := CountCarouses(cond)
	if err != nil {
		beego.Error(fmt.Sprintf("CountCarouses(%#v) err:%s\n", cond, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(cond), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	rows, err := GetCarousels(cond, pager.BuildLimiter())
	if err != nil {
		beego.Error(fmt.Sprintf("GetCarouselList(%#v,%#v) err:%s\n", cond, pager.BuildLimiter(), err))
		return nil, err
	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}

func CountCarouses(cond *Carousel) (int64, error) {
	if cond == nil {
		cond = &Carousel{}
	}
	return x.Count(cond)
}

/*func GetCarousels(cond *Carousel, limit, offset int) ([]*Carousel, error) {
	if cond == nil {
		cond = &Carousel{}
	}
	carousels := make([]*Carousel, 0, limit)
	err := x.Desc("caro_sort_no", "caro_created").Limit(limit, offset).Find(&carousels, cond)
	if err != nil {
		beego.Error(err)
		return carousels, err
	}
	return carousels, nil
}*/

func GetCarousels(cond *Carousel, lim *Limiter) (list []*Carousel, err error) {
	if cond == nil {
		cond = &Carousel{}
	}

	if lim != nil {
		list = make([]*Carousel, 0, lim.Limit)
		err = x.Desc("caro_sort_no", "caro_created").Limit(lim.Limit, lim.Offset).Find(&list, cond)
	} else {
		list = make([]*Carousel, 0, 10)
		err = x.Desc("caro_sort_no", "caro_created").Find(&list, cond)
	}

	return list, err
}

func GetCarouselById(id string) (carousel *Carousel, has bool, err error) {
	carousel = &Carousel{Id: id}
	has, err = x.Get(carousel)
	if err != nil {
		return nil, false, err
	}

	return carousel, has, nil
}

func AddCarousel(caro *Carousel) (affected int64, err error) {
	caro.Id = uuid.NewV4().String()
	return x.Insert(caro)
}

func RmCarousel(id string) (affected int64, err error) {
	return x.Delete(&Carousel{Id: id})
}

func ModCarousel(carousel *Carousel) (affected int64, err error) {
	affected, err = x.Update(carousel, &Carousel{Id: carousel.Id})
	return
}

func PushPinCarousel(id string) (affected int64, err error) {
	sql := "UPDATE t_home_carousel AS t1 SET t1.caro_sort_no = (SELECT MAX(caro_sort_no)+1 FROM (SELECT * FROM t_home_carousel)AS t0) WHERE t1.caro_id = ?"
	result, err := x.Exec(sql, id)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}
