package controllers

import (
	"corpweb/models"
	"fmt"

	"github.com/astaxie/beego"
)

type SettingsController struct {
	// beego.Controller
	baseController
}

// ToPersonalSetting 转到个人设置页
func (this *SettingsController) ToPersonalSetting() {
	this.Data["menu_lv_1"] = "settings"
	this.Data["menu_lv_2"] = "settings_personal"

	this.TplNames = "admin/settings_personal.html"
}

// ToSysSetting 转到系统设置页
func (this *SettingsController) ToSysSetting() {
	this.Data["menu_lv_1"] = "settings"
	this.Data["menu_lv_2"] = "settings_sys"

	this.TplNames = "admin/settings_sys.html"
}

func (this *SettingsController) ToChangePwd() {
	this.TplNames = "admin/ajax_html/settings_personal_form_change_pwd.html"
}

func (this *SettingsController) ToEmail() {
	user := this.GetSession("UserInfo").(models.User)

	sett, has, err := models.GetMailSettingsByUserId(user.UserId)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetMailSettingsByUserId(%s) err: %s", user.UserId, err))
	} else if has {
		this.Data["sett"] = sett
	}

	this.TplNames = "admin/ajax_html/settings_personal_form_email.html"
}

// SaveMailBoxInfo 保存邮件服务信息
func (this *SettingsController) SaveMailBoxInfo() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 4),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	sett := models.MailSettings{}
	if err := this.ParseForm(&sett); err != nil {
		beego.Error(err)
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	// 数据校验
	if len(sett.Account) == 0 || len(sett.Account) > 128 {
		resp.Fields["Account"] = this.Tr("tips_content_cant_empty_and_too_long", 128)
	} else if !emailPattern.MatchString(sett.Account) {
		resp.Fields["Account"] = this.Tr("tips_invalid_email")
	}

	if len(sett.Pwd) == 0 || len(sett.Pwd) > 50 {
		resp.Fields["Pwd"] = this.Tr("tips_content_cant_empty_and_too_long", 50)
	}

	if len(sett.Outgoing) == 0 || len(sett.Outgoing) > 100 {
		resp.Fields["Outgoing"] = this.Tr("tips_content_cant_empty_and_too_long", 100)
	}

	if sett.OutgoingPort > 65535 || sett.OutgoingPort == 0 {
		resp.Fields["OutgoingPort"] = this.Tr("tips_invalid_port")
	}

	if len(resp.Fields) > 0 {
		return
	}

	user := this.GetSession("UserInfo").(models.User)
	(&sett).UserId = user.UserId
	affected, err := models.ModMailSettings(&sett)
	if err != nil {
		beego.Error(fmt.Sprintf("models.ModMailSettings(%#v) err:%s", &sett, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	if affected <= 0 {
		_, err = models.AddMailSettings(&sett)
		if err != nil {
			beego.Error(fmt.Sprintf("models.AddMailSettings(%#v) err:%s", sett, err))
			resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
			return
		}
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}
