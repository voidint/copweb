package controllers

import (
	"fmt"
	"strings"
	"time"

	"corpweb/models"

	"github.com/astaxie/beego"
)

type MessageController struct {
	// beego.Controller
	baseController
}

// ToContactMsg 跳转至联系人消息页面
func (this *MessageController) ToContactMsg() {
	this.Data["menu_lv_1"] = "msg"
	this.Data["menu_lv_2"] = "msg_contact"

	this.TplNames = "admin/contact_msg_list.html"
}

// AjaxGetContactMsgPage 获取联系人消息分页列表
func (this *MessageController) AjaxGetContactMsgPage() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	curPageNo, _ := this.GetInt("curPageNo")
	pageSize, _ := this.GetInt("pageSize")

	cond := models.ContactMessage{}
	if err := this.ParseForm(&cond); err != nil {
		beego.Error(err)
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	page, err := models.GetContactMsgPage(&cond, curPageNo, pageSize)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = page
}

// SearchContactMsgPage 根据条件查询联系人消息记录。支持条件之间的and和or逻辑运算。
func (this *MessageController) SearchContactMsgPage() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	curPageNo, _ := this.GetInt("curPageNo")
	pageSize, _ := this.GetInt("pageSize")
	op := strings.ToUpper(this.GetString("op"))

	if op != models.LOGICAL_OP_OR {
		op = models.LOGICAL_OP_AND
	}

	// 转换失败则使用time.Time的零值
	beginDate, _ := time.Parse(models.LAYOUT_DATE_TIME, this.GetString("BeginDate"))
	endDate, _ := time.Parse(models.LAYOUT_DATE_TIME, this.GetString("EndDate"))

	cond := models.ContactMessage{}
	if err := this.ParseForm(&cond); err != nil {
		beego.Info(fmt.Sprintf("this.ParseForm() err:%s", err))
		return
	}

	page, err := models.SearchContactMsgPage(&cond, beginDate, endDate, models.LogicalOp(op), curPageNo, pageSize)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = page
}

// AjaxMarkMsg 标记联系人消息（修改消息状态）
func (this *MessageController) AjaxMarkMsg() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	id := this.GetString("Id")
	state, err := this.GetInt("State")

	if err != nil {
		beego.Error(fmt.Sprintf("this.GetInt(\"State\") err:%s", err))
		return
	}
	if state <= 0 {
		return
	}

	cond := &models.ContactMessage{
		Id:    id,
		State: state,
	}
	affected, err := models.ModContactMsg(cond)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if affected > 0 {
		resp.Result = RESULT_RESP_SUCC
		resp.Msg = this.Tr("tips_action_success")
	}
}
