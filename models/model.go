package models

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	// 性别常量定义
	GENDER_MALE   Gender = "male"
	GENDER_FEMALE Gender = "female"
	GENDER_OTHER  Gender = "other"

	// 可访问性
	ACCESSABLE_PUBLIC  uint8 = 1
	ACCESSABLE_PRIVATE uint8 = 2

	// 图片放置位置。1-封面；2-详情；
	IMG_PLACE_AT_COVER  uint8 = 1
	IMG_PLACE_AT_DETAIL uint8 = 2

	// 使用markdown。1-使用；2-未使用；
	USE_MD_YES uint8 = 1
	USE_MD_NO  uint8 = 2

	// 逻辑运算符
	LOGICAL_OP_AND string = "AND"
	LOGICAL_OP_OR  string = "OR"

	//消息状态
	MSG_STAT_UNPROCESSED uint8 = 1
	MSG_STAT_PROCESSED   uint8 = 2

	// ISO通用日期格式
	LAYOUT_DATE_TIME = "2006-01-02 15:04:05"
	LAYOUT_NOSEC     = "2006-01-02 15:04"
	LAYOUT_DATE      = "2006-01-02"
)

// 性别
type Gender string

// 逻辑运算符
type LogicalOp string

// Limiter 数据层分页对象
type Limiter struct {
	// Limit 分页中每页记录数
	Limit int
	// Offset 分页记录开始序号（偏移量）
	Offset int
}

// Pager 业务逻辑层分页接口
type Pager interface {
	// BuildLimiter 构建数据层分页对象
	BuildLimiter() *Limiter
	// AddRow 向当前分页中增加行记录
	AddRow(interface{}) bool
	// BuildPage 构建业务逻辑层分页对象
	BuildPage() *Page
}

// pagerImpl 业务逻辑层默认分页接口实现
type pagerImpl struct {
	curPageNo    int
	pageSize     int
	totalRecords int
	totalPages   int
	rows         []interface{}
	elemType     reflect.Type
}

// NewPager 新建业务逻辑层分页对象。
// elemType 分页中每条记录所代表的对象的反射类型
// curPageNo 当前页页号
// pageSize 每页记录数
// totalRecords 总记录数
func NewPager(elemType reflect.Type, curPageNo, pageSize, totalRecords int) Pager {
	if pageSize <= 0 {
		pageSize = 10
	}
	if curPageNo <= 0 {
		curPageNo = 1
	}

	totalPages, _ := CalcTotalPages(pageSize, totalRecords)
	return &pagerImpl{
		curPageNo:    curPageNo,
		pageSize:     pageSize,
		totalRecords: totalRecords,
		totalPages:   totalPages,
		rows:         make([]interface{}, 0, pageSize),
		elemType:     elemType,
	}
}

// BuildLimiter 构建数据层分页对象
func (this *pagerImpl) BuildLimiter() *Limiter {
	startRecordNo, _ := CalcStartRecordNo(this.curPageNo, this.pageSize, this.totalRecords)
	return &Limiter{
		Limit:  this.pageSize,
		Offset: startRecordNo,
	}
}

// isAcceptableElem 判断对象反射类型是否与分页记录对象反射类型一致。
func (this *pagerImpl) isAcceptableElem(k interface{}) bool {
	return reflect.TypeOf(k) == this.elemType
}

// AddRow 向当前分页中增加行记录
func (this *pagerImpl) AddRow(row interface{}) bool {
	if ok := this.isAcceptableElem(row); !ok {
		return false
	}
	this.rows = append(this.rows, row)
	return true
}

// BuildPage 构建业务逻辑层分页对象
func (this *pagerImpl) BuildPage() (page *Page) {
	page = &Page{
		CurPageNo:    this.curPageNo,
		PageSize:     this.pageSize,
		TotalRecords: this.totalRecords,
		TotalPages:   this.totalPages,
		Rows:         this.rows,
	}
	return page
}

// Page 业务逻辑层分页对象
type Page struct {
	CurPageNo    int           // 当前页页号
	PageSize     int           // 每页记录数
	TotalRecords int           // 总记录数
	TotalPages   int           // 总页数
	Rows         []interface{} // 当前页记录
}

func EmptyPage(curPageNo, pageSize int) *Page {
	return &Page{
		CurPageNo:    curPageNo,
		PageSize:     pageSize,
		TotalRecords: 0,
		TotalPages:   0,
		Rows:         make([]interface{}, 0),
	}
}

// CalcTotalPages 计算总页数
func CalcTotalPages(pageSize, totalRecords int) (totalPages int, err error) {
	if pageSize == 0 {
		return 0, errors.New("divide by zero")
	}
	if totalRecords%pageSize == 0 {
		totalPages = totalRecords / pageSize
	} else {
		totalPages = totalRecords/pageSize + 1
	}
	return totalPages, nil
}

// CalcStartRecordNo 计算分页起始记录号
func CalcStartRecordNo(curPageNo, pageSize, totalRecords int) (startRecordNo int, err error) {
	totalPages, err := CalcTotalPages(pageSize, totalRecords)
	if err != nil {
		return 0, err
	}

	if totalPages > 0 && curPageNo > totalPages {
		curPageNo = totalPages
	}
	startRecordNo = (curPageNo - 1) * pageSize
	return startRecordNo, nil
}

/*type DateTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"
const ctLayout_nosec = "2006-01-02 15:04"
const ctLayout_date = "2006-01-02"

func (this *DateTime) UnmarshalJSON(b []byte) (err error) {

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	sv := string(b)
	if len(sv) == 10 {
		sv += " 00:00:00"
	} else if len(sv) == 16 {
		sv += ":00"
	}
	this.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	if err != nil {
		if this.Time, err = time.ParseInLocation(ctLayout_nosec, string(b), loc); err != nil {
			this.Time, err = time.ParseInLocation(ctLayout_date, string(b), loc)
		}
	}

	return
}

func (this *DateTime) MarshalJSON() ([]byte, error) {

	rs := []byte(this.Time.Format(ctLayout))

	return rs, nil
}*/

/*var nilTime = (time.Time{}).UnixNano()

func (this *DateTime) IsSet() bool {
	return this.UnixNano() != nilTime
}*/

// xorm引擎
var x *xorm.Engine

func init() {
	// x, err := xorm.NewEngine("mysql", "root:abc#123@/test")//错误方式

	var err error
	// x, err = xorm.NewEngine("mysql", "root:abc#123@/cms4go?charset=utf8")

	dbname := beego.AppConfig.String("dbname")
	dbusername := beego.AppConfig.String("dbusername")
	dbuserpwd := beego.AppConfig.String("dbuserpwd")
	dsname := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbusername, dbuserpwd, dbname)

	x, err = xorm.NewEngine("mysql", dsname)

	if err != nil {
		beego.Error(err)
		return
	}

	x.ShowSQL = true
	x.ShowDebug = false
	x.ShowInfo = false
	x.ShowErr = true
	x.ShowWarn = true

	/*f, ferr := os.OpenFile("sql.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if ferr != nil {
		beego.Error(ferr)
		return
	}
	defer f.Close()
	x.Logger = xorm.NewSimpleLogger(f)*/

}
