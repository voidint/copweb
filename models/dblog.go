package models

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	// 操作名称
	ACTION_LOGIN            string = "login"
	ACTION_LOGOUT           string = "logout"
	ACTION_SEND_EMAIL       string = "send_email"
	ACTION_SEND_CONTACT_MSG string = "send_contact_msg"

	// 操作结果
	RESULT_SUCC string = "success"
	RESULT_FAIL string = "fail"

	// 操作发起者
	SPONSOR_VISITOR = "visitor"
)

type DBLog struct {
	Id       string    `xorm:"log_id pk"`
	Sponsor  string    `xorm:"log_sponsor notnull"`
	Terminal string    `xorm:"log_terminal notnull"`
	Action   string    `xorm:"log_action notnull"`
	Result   string    `xorm:"log_result notnull"`
	Msg      string    `xorm:"log_msg"`
	Created  time.Time `xorm:"log_created notnull created"`
}

func (this *DBLog) TableName() string {
	return "t_log"
}

func AddDBLog(log *DBLog) {
	log.Id = uuid.NewV4().String()
	// log.Created = time.Now()
	x.Insert(log)
}
