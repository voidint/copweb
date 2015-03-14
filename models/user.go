package models

import (
	"errors"
	"time"

	"corpweb/utils"
)

var (
	ErrLoginInvalid = errors.New("login info invalid")
)

type Login struct {
	UserId string `xorm:"login_user_id index"`
	Name   string `xorm:"login_name pk"` // 登录账号（邮箱）
	Pwd    string `xorm:"login_pwd notnull"`
	Salt   string `xorm:"login_salt notnull"`
}

func (this *Login) TableName() string {
	return "t_login"
}

// 校验登录信息
func ValidateLogin(name, rawPwd string) (login *Login, has bool, err error) {
	login = &Login{}
	has, err = x.Id(name).Get(login)
	if err != nil {
		return nil, false, err
	}

	if !has {
		return nil, false, nil
	}

	if md5str := utils.Md5String([]byte(name + rawPwd + login.Salt)); md5str != login.Pwd {
		return nil, true, ErrLoginInvalid
	}
	return login, true, nil
}

func CheckLogin(loginName, rawPwd string) (user *User, has bool, err error) {
	login := &Login{}
	has, err = x.Id(loginName).Get(login)
	if err != nil {
		return nil, false, err
	}
	if !has {
		return nil, false, nil
	}

	if md5str := utils.Md5String([]byte(loginName + rawPwd + login.Salt)); md5str != login.Pwd {
		return nil, true, nil
	}
	user, has, err = GetUserById(login.UserId)
	if user != nil {
		user.LoginName = loginName
	}
	return user, has, err
}

// 修改用户登录密码
func ChangePwd(name, rawNewPwd, salt string) (affected int64, err error) {
	pwd := utils.Md5String([]byte(name + rawNewPwd + salt))
	return x.Update(&Login{Pwd: pwd}, &Login{Name: name})
}

type User struct {
	UserId    string    `xorm:"user_id pk"`
	FullName  string    `xorm:"user_full_name"`
	Birthday  time.Time `xorm:"user_birthday"`
	Gender    string    `xorm:"user_gender"`
	Nickname  string    `xorm:"user_nickname"`
	Created   time.Time `xorm:"user_created created"`
	Modified  time.Time `xorm:"user_modified updated"`
	LoginName string    `xorm:"-"`
}

func (this *User) TableName() string {
	return "t_user"
}

// 根据用户信息ID查询用户信息
func GetUserById(userId string) (user *User, has bool, err error) {
	user = new(User)
	has, err = x.Id(userId).Get(user)
	return user, has, err
}
