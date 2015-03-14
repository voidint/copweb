package models

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
)

const (
	MSG_STATE_RAW int = iota + 1
)

// 联系人消息
type ContactMessage struct {
	Id       string    `xorm:"msg_id pk"`
	Name     string    `xorm:"msg_name notnull"`
	Email    string    `xorm:"msg_email notnull"`
	Phone    string    `xorm:"msg_phone notnull"`
	Company  string    `xorm:"msg_company"`
	Text     string    `xorm:"msg_text"`
	State    int       `xorm:"msg_state notnull"`
	Created  time.Time `xorm:"msg_created notnull created"`
	Modified time.Time `xorm:"msg_modified notnull updated"`
}

func (this *ContactMessage) TableName() string {
	return "t_contact_msg"
}

func AddContactMsg(msg *ContactMessage) (string, error) {
	msg.Id = uuid.NewV4().String()
	msg.State = MSG_STATE_RAW

	_, err := x.Insert(msg)
	if err != nil {
		return "", err
	}
	return msg.Id, nil
}

func GetContactMsgPage(cond *ContactMessage, curPageNo, pageSize int) (page *Page, err error) {
	if cond == nil {
		cond = &ContactMessage{}
	}

	totalRecords, err := CountContactMsgs(cond)
	if err != nil {
		beego.Error(fmt.Sprintf("CountContactMsgs(%#v) err:%s", cond, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(cond), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	rows, err := GetContactMsgs(cond, pager.BuildLimiter())
	if err != nil {
		beego.Error(fmt.Sprintf("GetContactMsgs(%#v,%#v) err:%s", cond, pager.BuildLimiter(), err))
		return nil, err
	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}

func CountContactMsgs(cond *ContactMessage) (int64, error) {
	if cond == nil {
		cond = &ContactMessage{}
	}
	return x.Count(cond)
}

func GetContactMsgs(cond *ContactMessage, lim *Limiter) (list []*ContactMessage, err error) {
	if cond == nil {
		cond = &ContactMessage{}
	}

	if lim != nil {
		list = make([]*ContactMessage, 0, lim.Limit)
		err = x.Desc("msg_created").Limit(lim.Limit, lim.Offset).Find(&list, cond)
	} else {
		list = make([]*ContactMessage, 0, 10)
		err = x.Desc("msg_created").Find(&list, cond)
	}

	return list, err
}

func ModContactMsg(cond *ContactMessage) (affected int64, err error) {
	return x.Update(cond, &ContactMessage{Id: cond.Id})
}

// SearchContactMsgPage 查询联系人消息分页记录。此函数与GetContactMsgPage函数的区别是，此函数中条件查询使用的是LIKE查询，并且支持条件之间的and和or逻辑运算。
// cond 联系人消息查询条件
// beginDate 联系人消息查询条件-消息创建时间下限
// endDate 联系人消息查询条件-消息创建时间上限
// op 查询条件之间的逻辑运算符（and、or）
// curPageNo 分页当前页号
// pageSize 分页每页显示记录数
func SearchContactMsgPage(cond *ContactMessage, beginDate time.Time, endDate time.Time, op LogicalOp, curPageNo, pageSize int) (page *Page, err error) {
	if string(op) != LOGICAL_OP_OR {
		op = LogicalOp(LOGICAL_OP_AND)
	}

	if cond == nil && beginDate.IsZero() && endDate.IsZero() {
		return GetContactMsgPage(cond, curPageNo, pageSize)
	}
	// 根据条件拼接生成WHERE条件语句
	whereSQL, args := getSearchContactMsgWhereSQL(cond, beginDate, endDate, op)

	// 计算符合条件的记录数量
	totalRecords, err := x.Where(whereSQL, args...).Count(&ContactMessage{})
	if err != nil {
		beego.Error(fmt.Sprintf("whereSQL:%s err:%s", whereSQL, err))
		return nil, err
	}

	pager := NewPager(reflect.TypeOf(&ContactMessage{}), curPageNo, pageSize, int(totalRecords))

	if totalRecords <= 0 {
		return page, nil
	}

	// 查询符合条件的分页记录
	lim := pager.BuildLimiter()
	rows := make([]*ContactMessage, 0, lim.Limit)
	err = x.Where(whereSQL, args...).Desc("msg_created").Limit(lim.Limit, lim.Offset).Find(&rows)

	if err != nil {
		beego.Error(fmt.Sprintf("whereSQL:%s err:%s", whereSQL, err))
		return nil, err
	}
	for i := range rows {
		pager.AddRow(rows[i])
	}

	return pager.BuildPage(), nil
}

// 根据条件拼接生成WHERE条件语句
func getSearchContactMsgWhereSQL(cond *ContactMessage, beginDate time.Time, endDate time.Time, op LogicalOp) (whereSQL string, args []interface{}) {
	sOp := " " + string(op) + " "
	args = make([]interface{}, 0, 8)

	var sqlBuf bytes.Buffer
	if cond != nil && cond.Name != "" {
		sqlBuf.WriteString("msg_name LIKE ?")
		sqlBuf.WriteString(sOp)
		args = append(args, "%"+cond.Name+"%")
	}

	if cond != nil && cond.Email != "" {
		sqlBuf.WriteString("msg_email LIKE ?")
		sqlBuf.WriteString(sOp)
		args = append(args, "%"+cond.Email+"%")
	}

	if cond != nil && cond.Phone != "" {
		sqlBuf.WriteString("msg_phone LIKE ?")
		sqlBuf.WriteString(sOp)
		args = append(args, "%"+cond.Phone+"%")
	}

	if cond != nil && cond.Company != "" {
		sqlBuf.WriteString("msg_company LIKE ?")
		sqlBuf.WriteString(sOp)
		args = append(args, "%"+cond.Company+"%")
	}

	if cond != nil && cond.Text != "" {
		sqlBuf.WriteString("msg_text LIKE ?")
		sqlBuf.WriteString(sOp)
		args = append(args, "%"+cond.Text+"%")
	}

	if cond != nil && cond.State > 0 {
		sqlBuf.WriteString("msg_state=?")
		sqlBuf.WriteString(sOp)
		args = append(args, cond.State)
	}

	if !beginDate.IsZero() {
		sqlBuf.WriteString("UNIX_TIMESTAMP(msg_created) >=?")
		sqlBuf.WriteString(sOp)
		args = append(args, beginDate.Unix())
	}

	if !endDate.IsZero() {
		sqlBuf.WriteString("UNIX_TIMESTAMP(msg_created) <=?")
		sqlBuf.WriteString(sOp)
		args = append(args, endDate.Unix())
	}

	return strings.TrimSuffix(strings.TrimSpace(sqlBuf.String()), string(op)), args
}
