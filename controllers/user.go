package controllers

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"corpweb/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	// beego.Controller
	baseController
}

// ToLogout 用户登出
func (this *UserController) ToLogout() {
	if user, ok := this.GetSession("UserInfo").(models.User); ok {
		beego.Debug("Session已删除")
		this.DelSession("UserInfo")
		this.CruSession.Flush()
		beego.GlobalSessions.SessionDestroy(this.Ctx.ResponseWriter, this.Ctx.Request)

		// flash := beego.NewFlash()
		// flash.Success("true")
		// flash.Store(&this.Controller)

		log := &models.DBLog{
			Sponsor:  user.UserId,
			Terminal: this.Ctx.Input.IP(),
			Action:   models.ACTION_LOGOUT,
			Result:   models.RESULT_SUCC,
		}
		models.AddDBLog(log)
	}

	this.Redirect("/login", 302)
}

// ToLogin 转到用户登录页面
func (this *UserController) ToLogin() {
	/*flash := beego.ReadFromRequest(&this.Controller)
	if _, ok := flash.Data["success"]; ok {
		this.Data["logout_result"] = "success"
		this.Data["logout_msg"] = "成功登出"
	}*/
	this.TplNames = "login.html"
}

// 用户登录校验
func (this *UserController) Login() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 2),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	// 数据校验
	name := strings.TrimSpace(this.GetString("Name"))
	if len(name) <= 0 {
		resp.Fields["Name"] = this.Tr("tips_account_name_cant_empty")
	}

	pwd := strings.TrimSpace(this.GetString("Pwd"))
	if len(pwd) <= 0 {
		resp.Fields["Pwd"] = this.Tr("tips_pwd_cant_empty")
	}

	if len(resp.Fields) > 0 {
		return
	}

	//登录校验
	user, has, err := models.CheckLogin(name, pwd)
	if err != nil {
		beego.Error(fmt.Sprintf("models.CheckLogin(%s, ***) err: %s", name, err))
		resp.Msg = this.Tr("tips_server_err")
		return
	}
	if !has || user == nil {
		// resp.Msg = this.Tr("tips_invalid_account_or_pwd")
		resp.Fields["Pwd"] = this.Tr("tips_invalid_account_or_pwd")
		return
	}

	// beego.Debug("登录成功....")
	this.SetSession("UserInfo", *user)

	log := &models.DBLog{
		Sponsor:  user.LoginName,
		Terminal: this.Ctx.Input.IP(),
		Action:   models.ACTION_LOGIN,
		Result:   models.RESULT_SUCC,
	}
	models.AddDBLog(log)
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// ChangePwd 修改当前用户登录密码
func (this *UserController) ChangePwd() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 2),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	user, ok := this.GetSession("UserInfo").(models.User)
	if !ok {
		return
	}

	loginName := user.LoginName
	oldPwd := strings.TrimSpace(this.GetString("OldPwd"))
	newPwd := strings.TrimSpace(this.GetString("NewPwd"))

	if oldPwdLen := utf8.RuneCountInString(oldPwd); oldPwdLen == 0 || oldPwdLen > 50 {
		resp.Fields["OldPwd"] = this.Tr("tips_content_cant_empty_and_too_long", 50)
	}

	if newPwdLen := utf8.RuneCountInString(newPwd); newPwdLen == 0 || newPwdLen > 50 {
		resp.Fields["NewPwd"] = this.Tr("tips_content_cant_empty_and_too_long", 50)
	}

	if len(resp.Fields) > 0 {
		return
	}

	// 校验旧密码
	login, _, err := models.ValidateLogin(loginName, oldPwd)
	if err != nil {
		if err == models.ErrLoginInvalid {
			resp.Fields["OldPwd"] = this.Tr("tips_invalid_account_or_pwd")
		} else {
			beego.Error(fmt.Sprintf("models.ValidateLogin(%s, ***) err:%s", loginName, err))
			resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		}
		return
	}

	// 修改密码
	affected, err := models.ChangePwd(loginName, newPwd, login.Salt)
	if err != nil {
		beego.Error(fmt.Sprintf("models.ChangePwd(%s, ***, ***) err:%s", loginName, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if affected > 0 {
		resp.Result = RESULT_RESP_SUCC
		resp.Msg = this.Tr("tips_action_success")
	}
}
