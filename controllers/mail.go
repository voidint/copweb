package controllers

import (
	"bytes"
	"corpweb/models"
	"corpweb/utils"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
)

type MailController struct {
	// beego.Controller
	baseController
}

const (
	// 邮件ContentType
	MailContentType = "text/html;charset=UTF-8"
	// 邮件地址间的分隔符
	MailSeparator = ";"
)

// SendMail 发送邮件
func (this *MailController) SendMail() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_mail_send_fail"),
		Fields: make(map[string]string, 3),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	mail := models.Mail{}
	if err := this.ParseForm(&mail); err != nil {
		beego.Error(err)
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	// 数据校验
	if toLen := len(mail.To); toLen == 0 {
		resp.Fields["To"] = this.Tr("tips_content_cant_empty_and_too_long", 1024)
	} else {
		// 逐个校验Email地址
		mails := strings.Split(mail.To, MailSeparator)
		var buf bytes.Buffer
		for i := range mails {
			addr := strings.TrimSpace(mails[i])
			if len(addr) == 0 {
				continue
			}
			if !emailPattern.MatchString(addr) {
				resp.Fields["To"] = this.Tr("tips_invalid_email")
				beego.Info(fmt.Sprintf("Invalid email addr:%s(%s)", addr, mail.To))
				break
			}
			buf.WriteString(addr)
			buf.WriteString(MailSeparator)
		}
		(&mail).To = strings.TrimSuffix(buf.String(), MailSeparator)
	}

	if subjectLen := len(mail.Subject); subjectLen == 0 || subjectLen > 512 {
		resp.Fields["Subject"] = this.Tr("tips_content_cant_empty_and_too_long", 512)
	}
	if bodyLen := len(mail.Body); bodyLen == 0 {
		resp.Fields["Body"] = this.Tr("tips_mail_body_cant_empty")
	}
	if len(resp.Fields) > 0 {
		return
	}

	// 提取邮件发送者邮件服务器配置信息
	user := this.GetSession("UserInfo").(models.User)
	sett, has, err := models.GetMailSettingsByUserId(user.UserId)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetMailSettingsByUserId(%s) err: %s", user.UserId, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if !has {
		resp.Msg = this.Tr("tips_mail_not_set")
		return
	}

	host := sett.Outgoing
	smtpPort := sett.OutgoingPort
	sender := sett.Account
	pwd := sett.Pwd

	addr := fmt.Sprintf("%s:%d", host, smtpPort)
	auth := smtp.PlainAuth("", sender, pwd, host)
	msgFmt := "From: %s<%s>\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s\r\n\r\n%s"
	msg := fmt.Sprintf(msgFmt, sender, sender, mail.To, mail.Subject, MailContentType, utils.Markdown2html(mail.Body))

	err = smtp.SendMail(addr, auth, sender, strings.Split(mail.To, MailSeparator), []byte(msg))
	if err != nil {
		beego.Error(fmt.Sprintf("send email err: %s\n%s", err, msg))
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_mail_send_success")
}
