package controllers

import (
	"fmt"
	"unicode/utf8"

	"corpweb/models"

	"github.com/astaxie/beego"
)

var (
	// 封面图片宽度
	coverImgWidth = 740
	// 封面图片高度
	coverImgHeight = 400
	// 封面图片数量上限
	coverFileNumLimit = 1
	//封面图片文件大小上限（单位MB）
	coverMaxSize = 5

	// 详情图片高度
	detailImgWidth = 750
	// 详情图片高度
	detailImgHeight = 500
	// 详情图片数量上限
	detailFileNumLimit = 5
	// 详情图片文件大小上限（单位MB）
	detailMaxSize = 5
)

type ProductController struct {
	// beego.Controller
	baseController
}

// ToProducts 转到产品列表页面
func (this *ProductController) ToProducts() {
	this.Data["menu_lv_1"] = "prod"
	this.Data["menu_lv_2"] = "prod_list"

	this.TplNames = "admin/product_list.html"
}

// ToProductAdd 跳转至产品添加页面
func (this *ProductController) ToProductAdd() {
	this.Data["menu_lv_1"] = "prod"
	this.Data["menu_lv_2"] = "prod_add"

	this.Data["coverImgWidth"] = coverImgWidth
	this.Data["coverImgHeight"] = coverImgHeight
	this.Data["coverFileNumLimit"] = coverFileNumLimit
	this.Data["coverMaxSize"] = coverMaxSize
	this.Data["detailImgWidth"] = detailImgWidth
	this.Data["detailImgHeight"] = detailImgHeight
	this.Data["detailFileNumLimit"] = detailFileNumLimit
	this.Data["detailMaxSize"] = detailMaxSize

	this.TplNames = "admin/product_form.html"
}

// ToProductMod 跳转至产品编辑页面
func (this *ProductController) ToProductMod() {
	this.Data["menu_lv_1"] = "prod"
	this.Data["menu_lv_2"] = "prod_list"

	// prodId := this.GetString("prodId")
	prodId := this.Ctx.Input.Param(":prodId")

	cond := &models.Product{Id: prodId}
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

	this.Data["prod"] = prods[0]
	this.Data["coverImgWidth"] = coverImgWidth
	this.Data["coverImgHeight"] = coverImgHeight
	this.Data["coverFileNumLimit"] = coverFileNumLimit
	this.Data["coverMaxSize"] = coverMaxSize
	this.Data["detailImgWidth"] = detailImgWidth
	this.Data["detailImgHeight"] = detailImgHeight
	this.Data["detailFileNumLimit"] = detailFileNumLimit
	this.Data["detailMaxSize"] = detailMaxSize

	this.TplNames = "admin/product_form.html"
}

//type DtGridPager struct {
//	IsSuccess    bool        `json:"isSuccess"`
//	PageSize     uint        `json:"pageSize"`
//	StartRecord  uint        `json:"startRecord"`
//	NowPage      uint        `json:"nowPage"`
//	RecordCount  uint        `json:"recordCount"`
//	PageCount    uint        `json:"pageCount"`
//	ExhibitDatas interface{} `json:"exhibitDatas"`
//}

// AjaxGetProductList ajax方式获取产品分页信息
func (this *ProductController) AjaxGetProductList() {
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

	cond := models.Product{}
	if err := this.ParseForm(&cond); err != nil {
		beego.Error(fmt.Printf("this.ParseForm() err: %s", err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	page, err := models.GetProductPage(&cond, curPageNo, pageSize, false)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if page == nil {
		page = models.EmptyPage(curPageNo, pageSize)
	}

	resp.ExtObj = page
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// AddProduct 添加产品信息
func (this *ProductController) AddProduct() {
	ajaxResp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_sys_err_and_contact_tech"),
		Fields: make(map[string]string),
	}
	defer func() {
		this.Data["json"] = &ajaxResp
		this.ServeJson(true)
	}()

	prod := models.Product{}
	if err := this.ParseForm(&prod); err != nil {
		beego.Error(err)
		ajaxResp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	// 表单数据校验
	if titleLen := utf8.RuneCountInString(prod.Title); titleLen == 0 || titleLen > 100 {
		ajaxResp.Fields["Title"] = this.Tr("tips_content_cant_empty_and_too_long", 100)
	}

	if introLen := utf8.RuneCountInString(prod.Intro); introLen == 0 || introLen > 200 {
		ajaxResp.Fields["Intro"] = this.Tr("tips_content_cant_empty_and_too_long", 200)
	}

	if descLen := utf8.RuneCountInString(prod.Desc); descLen == 0 || descLen > 5000 {
		ajaxResp.Fields["Desc"] = this.Tr("tips_content_cant_empty_and_too_long", 5000)
	}

	coverImgPath := this.GetString("CoverImg.Path")
	if length := len(coverImgPath); length == 0 {
		ajaxResp.Fields["CoverImg.Path"] = this.Tr("tips_need_img")
	} else if length > 512 {
		ajaxResp.Fields["CoverImg.Path"] = this.Tr("tips_invalid_img_path")
	}

	detailImgPaths := this.GetStrings("DetailImgs.Path")
	if len(detailImgPaths) == 0 {
		ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_need_img")
	}
	for _, path := range detailImgPaths {
		if length := len(path); length == 0 {
			continue
		} else if length > 512 {
			ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_invalid_img_path")
			break
		}
	}
	if len(detailImgPaths) > detailFileNumLimit {
		ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_img_over_file_num_limit", detailFileNumLimit)
	}

	if len(ajaxResp.Fields) > 0 {
		return
	}

	// xss过滤
	// prod.HtmlEscape()

	prod.CoverImg = &models.ProductImage{Path: coverImgPath}

	prod.DetailImgs = make([]*models.ProductImage, 0, len(detailImgPaths))
	for _, path := range detailImgPaths {
		img := &models.ProductImage{Path: path}
		prod.DetailImgs = append(prod.DetailImgs, img)
	}

	if err := models.AddProduct(&prod); err != nil {
		beego.Error(err)
		return
	}

	ajaxResp.Result = RESULT_RESP_SUCC
	ajaxResp.Msg = this.Tr("tips_action_success")

}

// ModProduct 修改产品信息
func (this *ProductController) ModProduct() {
	ajaxResp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_sys_err_and_contact_tech"),
		Fields: make(map[string]string),
	}
	defer func() {
		this.Data["json"] = &ajaxResp
		this.ServeJson(true)
	}()

	prod := models.Product{}
	if err := this.ParseForm(&prod); err != nil {
		beego.Error(err)
		ajaxResp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	// 表单数据校验
	if uuidLen := len(prod.Id); uuidLen != 36 {
		ajaxResp.Fields["Id"] = this.Tr("tips_invalid_uuid")
	}

	if titleLen := utf8.RuneCountInString(prod.Title); titleLen == 0 || titleLen > 100 {
		ajaxResp.Fields["Title"] = this.Tr("tips_content_cant_empty_and_too_long", 100)
	}

	if introLen := utf8.RuneCountInString(prod.Intro); introLen == 0 || introLen > 200 {
		ajaxResp.Fields["Intro"] = this.Tr("tips_content_cant_empty_and_too_long", 200)
	}

	if descLen := utf8.RuneCountInString(prod.Desc); descLen == 0 || descLen > 5000 {
		ajaxResp.Fields["Desc"] = this.Tr("tips_content_cant_empty_and_too_long", 5000)
	}

	coverImgPath := this.GetString("CoverImg.Path")
	if length := len(coverImgPath); length == 0 {
		ajaxResp.Fields["CoverImg.Path"] = this.Tr("tips_need_img")
	} else if length > 512 {
		ajaxResp.Fields["CoverImg.Path"] = this.Tr("tips_invalid_img_path")
	}

	detailImgPaths := this.GetStrings("DetailImgs.Path")
	if len(detailImgPaths) == 0 {
		ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_need_img")
	}
	for _, path := range detailImgPaths {
		if length := len(path); length == 0 {
			continue
		} else if length > 512 {
			ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_invalid_img_path")
			break
		}
	}
	if len(detailImgPaths) > detailFileNumLimit {
		ajaxResp.Fields["DetailImgs.Path"] = this.Tr("tips_img_over_file_num_limit", detailFileNumLimit)
	}

	if len(ajaxResp.Fields) > 0 {
		return
	}

	// xss过滤
	// prod.HtmlEscape()

	prod.CoverImg = &models.ProductImage{Path: coverImgPath}

	prod.DetailImgs = make([]*models.ProductImage, 0, len(detailImgPaths))
	for _, path := range detailImgPaths {
		img := &models.ProductImage{Path: path}
		prod.DetailImgs = append(prod.DetailImgs, img)
	}

	affected, err := models.ModProduct(&prod)
	if err != nil {
		beego.Error(err)
		return
	}
	if affected <= 0 {
		ajaxResp.Msg = this.Tr("tips_action_fail")
		return
	}

	ajaxResp.Result = RESULT_RESP_SUCC
	ajaxResp.Msg = this.Tr("tips_action_success")
}

// RmProduct 删除产品
func (this *ProductController) RmProducts() {
	prodId := this.GetString("prodId")
	ajaxResp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_sys_err_and_contact_tech"),
	}
	defer func() {
		this.Data["json"] = &ajaxResp
		this.ServeJson(true)
	}()

	affected, err := models.RmProducts(prodId)
	if err != nil {
		return
	}

	extMap := make(map[string]interface{}, 1)
	extMap["affected"] = affected
	ajaxResp.ExtMap = extMap
	ajaxResp.Result = RESULT_RESP_SUCC
	ajaxResp.Msg = this.Tr("tips_action_success")
}

// PushPin 产品排序置顶
func (this *ProductController) PushPin() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	prodId := this.GetString("prodId")
	affected, err := models.PushPinProduct(prodId)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = affected
}
