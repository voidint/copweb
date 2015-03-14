package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

type BlogImage struct {
	Id      string    `xorm:"blog_img_id pk"`
	BlogId  string    `xorm:"blog_img_blog_id notnull"`
	Path    string    `xorm:"blog_img_path notnull"`
	Created time.Time `xorm:"blog_img_created created"`
}

func (this *BlogImage) TableName() string {
	return "t_blog_img"
}

func GetBlogImagesByBlogId(blogId string) (imgs []*BlogImage, err error) {
	imgs = make([]*BlogImage, 0, 1)
	err = x.Find(&imgs, &BlogImage{BlogId: blogId})
	return imgs, err
}

type Blog struct {
	Id        string     `xorm:"blog_id pk"`
	Title     string     `xorm:"blog_title notnull"`
	Intro     string     `xorm:"-"`
	Body      string     `xorm:"blog_body notnull"`
	BodyUseMd uint8      `xorm:"blog_body_use_markdown"` //是否使用了markdown来书写。1-是；2-否；
	Tags      string     `xorm:"blog_tags"`
	IsPublic  uint8      `xorm:"blog_is_public"` // 是否公开。1-公开;2-不公开。
	SortNo    uint       `xorm:"blog_sort_no"`
	Created   time.Time  `xorm:"blog_created created"`
	Modified  time.Time  `xorm:"blog_modified updated"`
	Cover     *BlogImage `xorm:"-"`
}

func (this *Blog) TableName() string {
	return "t_blog"
}

func AddBlog(blog *Blog) error {
	blog.Id = uuid.NewV4().String()

	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return err
	}
	_, err := sess.InsertOne(blog)
	if err != nil {
		beego.Error(err)
		sess.Rollback()
		return err
	}

	if blog.Cover != nil {
		blog.Cover.Id = uuid.NewV4().String()
		blog.Cover.BlogId = blog.Id
		_, err = sess.InsertOne(blog.Cover)
		if err != nil {
			beego.Error(err)
			sess.Rollback()
			return err
		}
	}

	return sess.Commit()
}

func GetBlogPage(cond *Blog, curPageNo, pageSize int, lazy bool) (page *Page, err error) {
	if cond == nil {
		cond = &Blog{}
	}

	totalRecords, err := CountBlogs(cond)
	if err != nil {
		beego.Error(fmt.Sprintf("CountBlogs(%#v) err:%s", cond, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(cond), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	rows, err := GetBlogs(cond, pager.BuildLimiter(), lazy)
	if err != nil {
		beego.Error(fmt.Sprintf("GetBlogs(%#v,%#v) err:%s", cond, pager.BuildLimiter(), err))
		return nil, err
	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}

func CountBlogs(cond *Blog) (int64, error) {
	if cond == nil {
		cond = &Blog{}
	}
	return x.Count(cond)
}

func GetBlogs(cond *Blog, lim *Limiter, lazy bool) (list []*Blog, err error) {
	if cond == nil {
		cond = &Blog{}
	}

	if lim != nil {
		list = make([]*Blog, 0, lim.Limit)
		err = x.Desc("blog_sort_no", "blog_created").Limit(lim.Limit, lim.Offset).Find(&list, cond)
	} else {
		list = make([]*Blog, 0, 10)
		err = x.Desc("blog_sort_no", "blog_created").Find(&list, cond)
	}

	if !lazy {
		for i := range list {
			imgs, err := GetBlogImagesByBlogId(list[i].Id)
			if err != nil {
				beego.Error(err)
				return list, nil
			}
			if len(imgs) > 0 {
				list[i].Cover = imgs[0]
			}
		}
	}

	return list, err
}

func GetBlogById(id string, lazy bool) (blog *Blog, has bool, err error) {
	blog = &Blog{Id: id}
	has, err = x.Get(blog)
	if err != nil {
		return nil, false, err
	}
	if !has {
		return nil, false, nil
	}
	if !lazy {
		imgs, err := GetBlogImagesByBlogId(id)
		if err != nil {
			beego.Error(err)
			return blog, has, nil
		}
		if len(imgs) > 0 {
			blog.Cover = imgs[0]
		}
	}
	return blog, has, err
}

func RmBlog(cond *Blog) (affected int64, err error) {
	return x.Delete(cond)
}

func ModBlog(blog *Blog) (affected int64, err error) {
	affected, err = x.Update(blog, &Blog{Id: blog.Id})
	return affected, err
}

func PushPinBlog(id string) (affected int64, err error) {
	sql := "UPDATE t_blog AS t1 SET t1.blog_sort_no = (SELECT MAX(blog_sort_no)+1 FROM (SELECT * FROM t_blog)AS t0) WHERE t1.blog_id = ?"
	result, err := x.Exec(sql, id)
	if err != nil {
		beego.Error(err)
		return 0, err
	}
	return result.RowsAffected()
}
