package controllers

import (
	"fmt"
	"unicode/utf8"

	"corpweb/models"

	"github.com/astaxie/beego"
)

type BlogController struct {
	// beego.Controller
	baseController
}

// ToBlog 跳转至博客列表
func (this *BlogController) ToBlog() {
	this.Data["menu_lv_1"] = "blog"
	this.Data["menu_lv_2"] = "blog_list"
	this.TplNames = "admin/blog_list.html"
}

// ToAddBlog 跳转至写博客页面
func (this *BlogController) ToAddBlog() {
	this.Data["menu_lv_1"] = "blog"
	this.Data["menu_lv_2"] = "blog_write"
	this.TplNames = "admin/blog_form.html"
}

// ToEditBlog 跳转至博客编辑页
func (this *BlogController) ToEditBlog() {
	this.Data["menu_lv_1"] = "blog"
	this.Data["menu_lv_2"] = "blog_write"

	blogId := this.Ctx.Input.Param(":blogId")

	// beego.Info(fmt.Sprintf("blogId:%s", blogId))

	blog, has, err := models.GetBlogById(blogId, true)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetBlogById(%s) err:%s", blogId, err))
		this.Abort("500")
	}
	if !has {
		this.Abort("404")
	}
	this.Data["blog"] = blog
	this.TplNames = "admin/blog_form.html"
}

// AjaxAddBlog 添加博客
func (this *BlogController) AjaxAddBlog() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 3),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	blog := models.Blog{}
	if err := this.ParseForm(&blog); err != nil {
		beego.Error(err)
		return
	}

	// 表单数据校验
	if titleLen := utf8.RuneCountInString(blog.Title); titleLen == 0 || titleLen > 100 {
		resp.Fields["Title"] = this.Tr("tips_content_cant_empty_and_too_long", 100)
	}

	if bodyLen := utf8.RuneCountInString(blog.Body); bodyLen == 0 || bodyLen > 10000 {
		resp.Fields["Body"] = this.Tr("tips_content_cant_empty_and_too_long", 10000)
	}

	if tagsLen := utf8.RuneCountInString(blog.Tags); tagsLen > 250 {
		resp.Fields["Tags"] = this.Tr("tips_content_too_long", 250)
	}

	if len(resp.Fields) > 0 {
		return
	}

	if blog.BodyUseMd != models.USE_MD_YES {
		blog.BodyUseMd = models.USE_MD_NO
	}
	if blog.IsPublic != models.ACCESSABLE_PUBLIC {
		blog.IsPublic = models.ACCESSABLE_PRIVATE
	}

	err := models.AddBlog(&blog)
	if err != nil {
		beego.Error(fmt.Sprintf("models.AddBlog() err:%s", err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// AjaxModBlog 修改博客
func (this *BlogController) AjaxModBlog() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 3),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	blog := models.Blog{}
	if err := this.ParseForm(&blog); err != nil {
		beego.Error(err)
		return
	}

	// 表单数据校验
	if len(blog.Id) == 0 {
		resp.Fields["Id"] = this.Tr("tips_invalid_uuid")
	}

	if titleLen := utf8.RuneCountInString(blog.Title); titleLen == 0 || titleLen > 100 {
		resp.Fields["Title"] = this.Tr("tips_content_cant_empty_and_too_long", 100)
	}

	if bodyLen := utf8.RuneCountInString(blog.Body); bodyLen == 0 || bodyLen > 10000 {
		resp.Fields["Body"] = this.Tr("tips_content_cant_empty_and_too_long", 10000)
	}

	if tagsLen := utf8.RuneCountInString(blog.Tags); tagsLen > 250 {
		resp.Fields["Tags"] = this.Tr("tips_content_too_long", 250)
	}

	if len(resp.Fields) > 0 {
		return
	}

	if blog.BodyUseMd != models.USE_MD_YES {
		(&blog).BodyUseMd = models.USE_MD_NO
	}
	if blog.IsPublic != models.ACCESSABLE_PUBLIC {
		(&blog).IsPublic = models.ACCESSABLE_PRIVATE
	}

	_, err := models.ModBlog(&blog)
	if err != nil {
		beego.Error(fmt.Sprintf("models.ModBlog(%#v) err:%s", blog, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// AjaxRmBlog 删除博客
func (this *BlogController) AjaxRmBlog() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	blogId := this.Ctx.Input.Param(":blogId")
	affected, err := models.RmBlog(&models.Blog{Id: blogId})
	if err != nil {
		beego.Error(fmt.Sprintf("models.RmBlog(%s) err: %s", blogId, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	if affected > 0 {
		resp.Result = RESULT_RESP_SUCC
		resp.Msg = this.Tr("tips_action_success")
	}
}

// AjaxPushPin 置顶
func (this *BlogController) AjaxPushPin() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	blogId := this.Ctx.Input.Param(":blogId")
	affected, err := models.PushPinBlog(blogId)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = affected

}

// AjaxGetBlogList 获取博客分页列表
func (this *BlogController) AjaxGetBlogList() {
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

	page, err := models.GetBlogPage(nil, curPageNo, pageSize, true)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = page
}
