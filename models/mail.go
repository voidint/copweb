package models

import (
	"time"

	"github.com/satori/go.uuid"
)

// 邮件
type Mail struct {
	Id      string
	To      string
	Subject string
	Body    string
}

// 邮箱服务设置
type MailSettings struct {
	Id           string    `xorm:"mailset_id pk"`
	Account      string    `xorm:"mailset_account notnull"`
	Pwd          string    `xorm:"mailset_pwd notnull"`
	Outgoing     string    `xorm:"mailset_outgoing notnull"`
	OutgoingPort uint      `xorm:"mailset_outgoing_port notnull"`
	Created      time.Time `xorm:"mailset_created created"`
	Modified     time.Time `xorm:"mailset_modified updated"`
	// Owner        User      `xorm:"mailset_owner"`
	UserId string `xorm:"mailset_owner notnull"`
}

func (this *MailSettings) TableName() string {
	return "t_settings_mail"
}

// AddMailSettings 新增邮件服务配置
func AddMailSettings(sett *MailSettings) (id string, err error) {
	sett.Id = uuid.NewV4().String()

	_, err = x.Insert(sett)
	if err != nil {
		return "", err
	}
	return sett.Id, nil
}

// ModMailSettings 修改邮件服务配置
func ModMailSettings(cond *MailSettings) (affected int64, err error) {
	return x.Update(cond, &MailSettings{UserId: cond.UserId})
}

// 根据用户ID获取用户邮件服务配置信息
func GetMailSettingsByUserId(userId string) (sett *MailSettings, has bool, err error) {
	if len(userId) == 0 {
		return nil, false, nil
	}
	sett = &MailSettings{UserId: userId}
	has, err = x.Get(sett)
	return sett, has, err
}
