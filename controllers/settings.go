package controllers

import (
	"corpweb/models"
	"corpweb/utils"
	"encoding/base64"
	"fmt"
	"time"

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

	if sett.Pwd == "******" {
		// 用户并非要修改密码
		(&sett).Pwd = ""
	} else {
		// 使用AES加密存储用户邮箱密码（此种方式不够安全）
		enc, err := utils.AesEncrypt([]byte(sett.Pwd), []byte(MailPwdKey))
		if err != nil {
			beego.Error(fmt.Sprintf("AesEncrypt err:%s", &sett, err))
			resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
			return
		}
		(&sett).Pwd = base64.StdEncoding.EncodeToString(enc)
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

// CheckAuth 检测邮箱帐号密码
func (this *SettingsController) CheckAuth() {
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

	if sett.Pwd == "******" {
		// 提取邮件发送者邮件服务器配置信息
		user := this.GetSession("UserInfo").(models.User)
		setting, has, err := models.GetMailSettingsByUserId(user.UserId)
		if err != nil {
			beego.Error(fmt.Sprintf("models.GetMailSettingsByUserId(%s) err: %s", user.UserId, err))
			resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
			return
		}

		if has {
			// 使用AES解密用户邮箱密码
			bEnc, err := base64.StdEncoding.DecodeString(setting.Pwd)
			if err != nil {
				beego.Error(fmt.Sprintf("base64.StdEncoding.DecodeString err: %s", err))
				resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
				return
			}
			bRaw, err := utils.AesDecrypt(bEnc, []byte(MailPwdKey))
			if err != nil {
				beego.Error(fmt.Sprintf("utils.AesDecrypt err: %s", err))
				resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
				return
			}
			sett.Pwd = string(bRaw)
		}

	}

	// 检测邮箱帐号密码（5秒超时）
	err := utils.CheckAuth(sett.Outgoing, sett.OutgoingPort, sett.Account, sett.Pwd, time.Second*5)
	if err != nil {
		if err == utils.ErrTimeout {
			resp.Msg = this.Tr("tips_mail_test_timeout", "5s")
		} else {
			resp.Msg = this.Tr("tips_mail_test_invalid")
		}
		return
	}

	resp.Msg = this.Tr("tips_mail_test_valid")
	resp.Result = RESULT_RESP_SUCC
}
