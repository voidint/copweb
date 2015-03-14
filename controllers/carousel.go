package controllers

import (
	"fmt"
	"unicode/utf8"

	"corpweb/models"

	"github.com/astaxie/beego"
)

type CarouselController struct {
	// beego.Controller
	baseController
}

// ToCarousel 跳转到首页滚动巨幕管理页
func (this *CarouselController) ToCarousel() {
	this.Data["menu_lv_1"] = "home"
	this.Data["menu_lv_2"] = "carousel"
	this.TplNames = "admin/home_carousel.html"
}

// ToAddCarousel 跳转到首页carousel表单页
func (this *CarouselController) ToAddCarousel() {
	this.Data["menu_lv_1"] = "home"
	this.Data["menu_lv_2"] = "carousel"

	imgwidth := 1000
	imgheight := 650

	this.Data["imgwidth"] = imgwidth
	this.Data["imgheight"] = imgheight
	this.TplNames = "admin/home_carousel_form.html"
}

// ToModCarousel 跳转到carousel编辑页
func (this *CarouselController) ToModCarousel() {
	this.Data["menu_lv_1"] = "home"
	this.Data["menu_lv_2"] = "carousel"

	carouselId := this.GetString("carouselId")
	carousel, has, err := models.GetCarouselById(carouselId)
	if err != nil {
		beego.Error(err)
		this.Abort("500")
	}
	if !has {
		this.Abort("404")
	}

	imgwidth := 1000
	imgheight := 650

	this.Data["imgwidth"] = imgwidth
	this.Data["imgheight"] = imgheight
	this.Data["carousel"] = carousel
	this.TplNames = "admin/home_carousel_form.html"
}

// GetCarousels ajax获取carousel信息
func (this *CarouselController) GetCarousels() {
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

	page, err := models.GetCarouselPage(nil, curPageNo, pageSize)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = page

}

// RmCarousel ajax删除carousel
func (this *CarouselController) RmCarousel() {
	result := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &result
		this.ServeJson(true)
	}()

	carouselId := this.GetString("carouselId")
	affected, err := models.RmCarousel(carouselId)
	if err != nil {
		beego.Error(fmt.Sprintf("models.RmCarousel(%s) err: %s\n", carouselId, err))
		result.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	if affected > 0 {
		result.Result = RESULT_RESP_SUCC
		result.Msg = this.Tr("tips_action_success")
	}
}

// AddHomeCarousel ajax添加carousel
func (this *CarouselController) AddCarousel() {
	result := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 2),
	}
	defer func() {
		this.Data["json"] = &result
		this.ServeJson(true)
	}()

	car := &models.Carousel{}
	if err := this.ParseForm(car); err != nil {
		beego.Error(err)
		result.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if pathLen := len(car.ImgPath); pathLen == 0 || pathLen > 512 {
		result.Fields["ImgPath"] = this.Tr("tips_invalid_img_path")
	}
	if captionLen := utf8.RuneCountInString(car.Caption); captionLen == 0 || captionLen > 200 {
		result.Fields["Caption"] = this.Tr("tips_content_cant_empty_and_too_long", 200)
	}
	if len(result.Fields) > 0 {
		return
	}

	affected, err := models.AddCarousel(car)
	if err != nil {
		beego.Error(fmt.Sprintf("models.AddCarousel(%#v) err:%s\n", car, err))
		result.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if affected > 0 {
		result.Result = RESULT_RESP_SUCC
		result.Msg = this.Tr("tips_action_success")
	}
}

// ModCarousel ajax修改carousel
func (this *CarouselController) ModCarousel() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
		Fields: make(map[string]string, 2),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	car := &models.Carousel{}
	if err := this.ParseForm(car); err != nil {
		beego.Error(err)
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	if len(car.Id) != 36 {
		resp.Fields["Id"] = this.Tr("tips_invalid_uuid")
	}

	if pathLen := len(car.ImgPath); pathLen == 0 || pathLen > 512 {
		resp.Fields["ImgPath"] = this.Tr("tips_invalid_img_path")
	}
	if captionLen := utf8.RuneCountInString(car.Caption); captionLen == 0 || captionLen > 200 {
		resp.Fields["Caption"] = this.Tr("tips_content_cant_empty_and_too_long", 200)
	}
	if len(resp.Fields) > 0 {
		return
	}

	_, err := models.ModCarousel(car)
	if err != nil {
		beego.Error(fmt.Sprintf("models.ModCarousel(%#v) err:%s\n", car, err))
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}
	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
}

// PushPin 置顶
func (this *CarouselController) PushPin() {
	resp := AjaxFormResp{
		Result: RESULT_RESP_FAIL,
		Msg:    this.Tr("tips_action_fail"),
	}
	defer func() {
		this.Data["json"] = &resp
		this.ServeJson(true)
	}()

	carouselId := this.GetString("carouselId")
	affected, err := models.PushPinCarousel(carouselId)
	if err != nil {
		resp.Msg = this.Tr("tips_sys_err_and_contact_tech")
		return
	}

	resp.Result = RESULT_RESP_SUCC
	resp.Msg = this.Tr("tips_action_success")
	resp.ExtObj = affected
}
