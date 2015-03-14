package controllers

import (
	"fmt"
	"html"
	// "regexp"
	"strings"
	"unicode/utf8"

	"corpweb/models"
	"corpweb/utils"

	"github.com/astaxie/beego"
)

type PortalController struct {
	// beego.Controller
	baseController
}

// ToHome 跳转至首页
func (this *PortalController) ToHome() {
	beego.Debug("PortalController.Home()...")

	// carousels, err := models.GetCarousels(nil, 3, 0)
	carousels, err := models.GetCarousels(nil, nil)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetCarouselList(nil,nil) err:%s\n", err))
		this.Abort("500")
	}

	fProds, err := models.GetFlagshipProducts(nil, 3, 0, false)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetFlagshipProducts(%#v,%d,%d) err:%s\n", nil, 3, 0, err))
		this.Abort("500")
	}

	this.Data["carousels"] = carousels
	this.Data["fProds"] = fProds
	this.Data["activeMenu"] = "home"
	this.TplNames = "home.html"
}

// ToProducts 跳转至产品列表
func (this *PortalController) ToProducts() {
	this.Data["activeMenu"] = "products"
	this.TplNames = "products.html"
}

// GetProductItems ajax方式获取产品数据
func (this *PortalController) GetProducts() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	curPageNo, _ := this.GetInt("curPageNo")
	pageSize, _ := this.GetInt("pageSize")

	if curPageNo <= 0 {
		curPageNo = 1
	}

	if pageSize <= 0 {
		pageSize = 6
	}
	cond := &models.Product{IsPublic: models.ACCESSABLE_PUBLIC}
	totalRecords, err := models.CountProducts(cond)
	if err != nil {
		beego.Error(fmt.Printf("models.CountProducts(%#v) err: %s", cond, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if totalRecords <= 0 {
		resp.Result = RESULT_RESP_SUCC
		return
	}

	startRecordNo, _ := this.CalcStartRecordNo(int64(curPageNo), int64(pageSize), totalRecords)
	prodCond := &models.Product{IsPublic: models.ACCESSABLE_PUBLIC}
	prods, err := models.GetProducts(prodCond, pageSize, int(startRecordNo), false)
	if err != nil {
		beego.Error(fmt.Printf("models.GetProducts(%#v,%d,%d) err: %s", prodCond, pageSize, startRecordNo, false, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	prodLen := len(prods)
	if prodLen == 0 {
		resp.Result = RESULT_RESP_SUCC
		return
	}

	// 2 items per row
	rowLen := 0
	if prodLen%2 == 0 {
		rowLen = prodLen / 2
	} else {
		rowLen = prodLen/2 + 1
	}

	rows := make([][]*models.Product, 0, rowLen)
	for i := 0; i < prodLen; i = i + 2 {
		if i%2 == 0 {
			row := make([]*models.Product, 0, 2)
			row = append(row, prods[i])
			if i+1 < prodLen {
				row = append(row, prods[i+1])
			}
			rows = append(rows, row)
		}
	}

	extMap := make(map[string]interface{}, 2)
	extMap["rows"] = rows
	totalPages, _ := this.CalcTotalPages(int64(pageSize), totalRecords)
	extMap["totalPages"] = totalPages

	resp.ExtMap = extMap
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// ToProductItem 跳转至产品明细
func (this *PortalController) ToProductItem() {
	beego.Debug("PortalController.ProductItem()...")

	prodId := this.Ctx.Input.Param(":prodId")
	cond := &models.Product{
		Id:       prodId,
		IsPublic: models.ACCESSABLE_PUBLIC,
	}
	prods, err := models.GetProducts(cond, 1, 0, false)

	if err != nil {
		beego.Error(fmt.Sprintf("models.GetProducts(%#v,1,0,false) err:\n", cond, err))
		// 重定向至500
		this.Abort("500")
	}
	if len(prods) == 0 {
		// 重定向至404
		this.Abort("404")
	}
	if prods[0].DescUseMarkdown == models.USE_MD_YES {
		prods[0].Desc = utils.Markdown2html(prods[0].Desc)
	}
	this.Data["prod"] = prods[0]
	this.Data["activeMenu"] = "products"
	this.TplNames = "product_item.html"
}

// ToBlog 跳转至博客列表
func (this *PortalController) ToBlog() {
	beego.Debug("PortalController.Blog()...")
	this.Data["activeMenu"] = "blog"
	this.TplNames = "blog.html"
}

// AjaxGetBlogPage 获取博客分页信息
func (this *PortalController) AjaxGetBlogPage() {
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

	cond := &models.Blog{
		IsPublic: models.ACCESSABLE_PUBLIC,
	}
	page, err := models.GetBlogPage(cond, curPageNo, pageSize, false)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	for i := range page.Rows {
		if blog, ok := page.Rows[i].(*models.Blog); ok {
			if blog.BodyUseMd == models.USE_MD_YES {
				blog.Body = utils.Markdown2html(blog.Body)
			}
			// 寻找博客内容中首个段落<p></p>，将其赋值给Intro字段
			bIdx := strings.Index(blog.Body, "<p>")
			eIdx := strings.Index(blog.Body, "</p>")
			if bIdx > -1 && eIdx > -1 && bIdx < eIdx {
				blog.Intro = string(blog.Body[bIdx+3 : eIdx])
			}
			if utf8.RuneCountInString(blog.Intro) > 150 {
				blog.Intro = string(blog.Intro[0:149])
			}
			blog.Body = ""
		}
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = page
}

// ToBlogPost 跳转至博客明细
func (this *PortalController) ToBlogPost() {
	blogId := this.Ctx.Input.Param(":blogId")
	blog, has, err := models.GetBlogById(blogId, true)
	if err != nil {
		beego.Error(fmt.Sprintf("models.GetBlogById(%s,true) err:%s", blogId, err))
		this.Abort("500")
	}
	if !has || blog.IsPublic == models.ACCESSABLE_PRIVATE {
		this.Abort("404")
	}

	if blog.BodyUseMd == models.USE_MD_YES {
		blog.Body = utils.Markdown2html(blog.Body)
	}
	this.Data["blog"] = blog
	this.Data["activeMenu"] = "blog"
	this.TplNames = "blog_post.html"
}

// ToAbout 跳转至关于我们
func (this *PortalController) ToAbout() {
	beego.Debug("PortalController.About()...")
	this.Data["activeMenu"] = "about"
	this.TplNames = "about.html"
}

// ToContact 跳转至联系我们
func (this *PortalController) ToContact() {
	beego.Debug("PortalController.ContactUs()...")

	this.Data["activeMenu"] = "contact"
	this.TplNames = "contact_us.html"
}

// AddContactMsg 跳转至添加联系人消息
func (this *PortalController) AddContactMsg() {
	beego.Debug("PortalController.AddContactMsg()...")

	result := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		// Msg:    "system error",
		Fields: make(map[string]string),
	}

	defer func() {
		this.Data["json"] = &result
		this.ServeJson(true)
	}()

	msg := models.ContactMessage{}
	if err := this.ParseForm(&msg); err != nil {
		beego.Error(err)
		result.Msg = this.Tr("tips_server_err")
		return
	}

	// xss过滤
	msg.Name = html.EscapeString(strings.TrimSpace(msg.Name))
	msg.Email = html.EscapeString(strings.TrimSpace(msg.Email))
	msg.Phone = html.EscapeString(strings.TrimSpace(msg.Phone))
	msg.Company = html.EscapeString(strings.TrimSpace(msg.Company))
	msg.Text = html.EscapeString(strings.TrimSpace(msg.Text))

	// 表单校验
	if msg.Name == "" || utf8.RuneCountInString(msg.Name) > 50 {
		result.Fields["Name"] = this.Tr("tips_content_cant_empty_and_too_long", 50)
	}
	if !emailPattern.MatchString(msg.Email) {
		result.Fields["Email"] = this.Tr("tips_invalid_email")
	}
	// !mobilePattern.MatchString(msg.Phone) || !telPattern.MatchString(msg.Phone)
	if msg.Phone == "" {
		result.Fields["Phone"] = this.Tr("tips_invalid_phone")
	}
	if utf8.RuneCountInString(msg.Company) > 150 {
		result.Fields["Companey"] = this.Tr("tips_content_too_long", 150)
	}
	if utf8.RuneCountInString(msg.Text) > 1024 {
		result.Fields["Text"] = this.Tr("tips_content_too_long", 1024)
	}

	if len(result.Fields) > 0 {
		return
	}

	id, err := models.AddContactMsg(&msg)
	if err != nil {
		beego.Error(err)
		result.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	log := &models.DBLog{
		Sponsor:  models.SPONSOR_VISITOR,
		Terminal: this.Ctx.Input.IP(),
		Action:   models.ACTION_SEND_CONTACT_MSG,
		Result:   models.RESULT_SUCC,
		Msg:      fmt.Sprintf("contact message id: %s", id),
	}
	models.AddDBLog(log)

	result.Result = RESULT_RESP_SUCC
	result.Msg = this.Tr("tips_thanks_contacting_us")
}
